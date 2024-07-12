package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
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

func TestAscii_Art(t *testing.T){
	tests := []struct {
		name       string
		urlPath    string
		method     string
		body       string
		wantStatus int
	}{
		{
			name:       "Valid GET request",
			urlPath:    "/",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Invalid URL path",
			urlPath:    "/invalid",
			method:     http.MethodGet,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Invalid method",
			urlPath:    "/",
			method:     http.MethodPut,
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "Bad request with invalid form data",
			urlPath:    "/",
			method:     http.MethodPost,
			body:       "text=something&banner=badbanner",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.method == http.MethodPost {
				form := urlPath.Values{}
				form.Add("text", "something")
				form.Add("banner", "badbanner")
				req = httptest.NewRequest(tt.method, tt.urlPath, strings.NewReader(tt.body))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			} else {
				req = httptest.NewRequest(tt.method, tt.urlPath, nil)
			}

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(Home)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}
		})
	}
}
