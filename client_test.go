package gobacklog

import (
	"net/http"
	"net/url"
	"testing"
)

// func TestNewClientAdjustedBaseURL(t *testing.T) {
// 	tab := []struct {
// 		BaseURL     string
// 		AdjustedURL string
// 	}{
// 		{
// 			BaseURL:     "http://example.com/",
// 			AdjustedURL: "http://example.com",
// 		},
// 		{
// 			BaseURL:     "http://example.com",
// 			AdjustedURL: "http://example.com",
// 		},
// 		{
// 			BaseURL:     "",
// 			AdjustedURL: "",
// 		},
// 	}
// 	for _, v := range tab {
// 		url, _ := url.Parse(v.BaseURL)
// 		c := NewClient(url, "")
// 		if c.BaseURL.String() != v.AdjustedURL {
// 			t.Errorf(`NewClient(%q, ""): BaseURL = %q; want %q`, v.BaseURL, c.BaseURL, v.AdjustedURL)
// 		}
// 	}
// }

func values(kv map[string][]string) url.Values {
	pairs := url.Values{}
	for k, v := range kv {
		for _, s := range v {
			pairs.Add(k, s)
		}
	}
	return pairs
}

func TestResolvingURL(t *testing.T) {
	apiKey := "apikey"
	tab := []struct {
		Endpoint    string
		BaseURL     string
		Params      url.Values
		AdjustedURL string
	}{
		{
			Endpoint:    "/api/v2/space",
			BaseURL:     "http://example.com/",
			Params:      url.Values{},
			AdjustedURL: "http://example.com/api/v2/space?apiKey=apikey",
		},
		{
			Endpoint:    "/api/v2/space",
			BaseURL:     "http://example.com",
			Params:      url.Values{},
			AdjustedURL: "http://example.com/api/v2/space?apiKey=apikey",
		},
		{
			Endpoint:    "/api/v2/space",
			BaseURL:     "http://example.com",
			Params:      values(map[string][]string{"a": []string{"b"}}),
			AdjustedURL: "http://example.com/api/v2/space?a=b&apiKey=apikey",
		},
	}
	for _, v := range tab {
		url, _ := url.Parse(v.BaseURL)
		c := NewClient(url, apiKey)
		// result := c.buildURLWithValues(v.BaseURL, v.Endpoint, v.Params)
		result := c.composeURL(v.Endpoint, v.Params)
		if result != v.AdjustedURL {
			_, requestErr := http.NewRequest("GET",
				result,
				nil,
			)
			if requestErr != nil {
				t.Error(requestErr)
			}

			t.Errorf(`Result = %q; BaseURL = %q; want %q`, result, c.BaseURL, v.AdjustedURL)
		}
	}
}
