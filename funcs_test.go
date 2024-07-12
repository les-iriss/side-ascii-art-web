package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	funcs "ascii-art-web/funcs"
)

func TestHome(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		method       string
		expectedCode int
	}{
		{
			name:         "valid request",
			url:          "/",
			method:       http.MethodGet,
			expectedCode: http.StatusOK,
		},
		{
			name : "Invalid URL Path",
			url: "/invalid",
			method: http.MethodGet,
			expectedCode: http.StatusNotFound,
		},
		{
			name: "Invalid method",
			url: "/",
			method: http.MethodPost,
			expectedCode: http.StatusMethodNotAllowed,
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(funcs.Home)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedCode)
			}

		})
	}
}



func TestAscii_Art(t *testing.T) {
    tests := []struct {
        name       string
        url        string
        method     string
        body       string // used for valid request
        status     int
    }{
        {
            name:       "Valid POST request",
            url:        "/ascii-art",
            method:     http.MethodPost,
            body:       "text=something&banner=standard",
            status:     http.StatusFound,
        },
        {
            name:       "Invalid URL path",
            url:        "/ascii-art/invalid",
            method:     http.MethodPut,
            status:     http.StatusNotFound,
        },
        {
            name:       "Invalid method",
            url:        "/ascii-art",
            method:     http.MethodGet,
            status:     http.StatusMethodNotAllowed,
        },
        {
            name:       "Bad request",
            url:        "/ascii-art",
            method:     http.MethodPost,
            body:       "invalid_body_format",
            status:     http.StatusBadRequest,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var req *http.Request
	    // this is for the case of good request , i need to set up the body correctly
            if tt.status == http.StatusFound {
                req = httptest.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
                req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
            } else {
                req = httptest.NewRequest(tt.method, tt.url, nil)
            }

            rr := httptest.NewRecorder()

            handler := http.HandlerFunc(funcs.Ascii_Art)
            handler.ServeHTTP(rr, req)

            if status := rr.Code; status != tt.status {
                t.Errorf("handler returned wrong status code: got %v want %v", status, tt.status)
            }
        })
    }
}

