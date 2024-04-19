package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(helloHandler))
	defer ts.Close()

	tests := []struct {
		name   string
		method string
		want   string
		code   int
	}{
		{name: "GET request", method: http.MethodGet, want: "Hello, World!", code: http.StatusOK},
		{name: "POST request", method: http.MethodPost, want: "Hello, World!", code: http.StatusOK},
		{name: "PUT request", method: http.MethodPut, want: "Hello, World!", code: http.StatusOK},
		{name: "DELETE request", method: http.MethodDelete, want: "Hello, World!", code: http.StatusOK},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/", ts.URL)

			req, err := http.NewRequest(tc.method, url, nil)
			if err != nil {
				t.Fatal(err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			got, err := io.ReadAll(res.Body) // Use ioutil for test simplicity
			if err != nil {
				t.Fatal(err)
			}

			// Check response status code
			if res.StatusCode != tc.code {
				t.Errorf("Unexpected status code: got %d, want %d", res.StatusCode, tc.code)
			}

			// Check response body (if expected)
			if tc.want != "" && !bytes.Equal(got, []byte(tc.want)) {
				t.Errorf("Unexpected response body: got %s, want %s", string(got), tc.want)
			}
		})
	}
}
