package funcs

import (
	"net/http"
	"net/http/httptest"
	"testing"
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

		})
	}
}
