package funcs

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHome(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		method       string
		expectedCode int
		expectedFile string
		expectedBody string
	}{
		{
			name:         "valid request",
			url:          "/",
			method:       http.MethodGet,
			expectedCode: http.StatusOK,
			expectedFile: "testdata/index.html.golden",
		},
		{
			name:         "invalid URL",
			url:          "/invalid",
			method:       http.MethodGet,
			expectedCode: http.StatusNotFound,
			expectedBody: not_found,
		},
		{
			name:         "invalid method",
			url:          "/",
			method:       http.MethodPost,
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: not_allowed,
		},
		{
			name:         "template parse error",
			url:          "/",
			method:       http.MethodGet,
			expectedCode: http.StatusInternalServerError,
			expectedBody: internal_error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Home)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedCode)
			}

			if tt.expectedCode == http.StatusOK && tt.expectedFile != "" {
				expected, err := ioutil.ReadFile(tt.expectedFile)
				if err != nil {
					t.Fatal(err)
				}
				if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(string(expected)) {
					t.Errorf("handler returned unexpected body: got %v want %v",
						rr.Body.String(), string(expected))
				}
			} else if tt.expectedBody != "" {
				if strings.TrimSpace(rr.Body.String()) != tt.expectedBody {
					t.Errorf("handler returned unexpected body: got %v want %v",
						rr.Body.String(), tt.expectedBody)
				}
			}
		})
	}
}
