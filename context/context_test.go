package context

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

type StubStore struct {
	response string
}

var RequestCancelledError = errors.New("Request cancelled")

func (s *StubStore) Fetch(ctx context.Context) (string, error) {

	select {
	case <-ctx.Done():
		return "", RequestCancelledError
	default:
		return s.response, nil
	}

}

func TestServer(t *testing.T) {

	t.Run("Should return nil body when request gets cancelled", func(t *testing.T) {
		data := "Hello, Gophers"
		svr := Server(&StubStore{data})

		ctx, cancel := context.WithCancel(context.Background())
		request := httptest.NewRequest(http.MethodGet, "/", nil).WithContext(ctx)
		response := &SpyResponseWriter{}

		cancel()
		svr.ServeHTTP(response, request)

		if response.written {
			t.Errorf("Incorrect response body: expected nothing to be written, but there was a write")
		}
	})

	t.Run("Should return expected body when request does not get cancelled", func(t *testing.T) {
		data := "Hello, Gophers"
		svr := Server(&StubStore{data})

		ctx, _ := context.WithCancel(context.Background())
		request := httptest.NewRequest(http.MethodGet, "/", nil).WithContext(ctx)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		got := response.Body.String()
		if got != data {
			t.Errorf("Expected response body to be %s, got %s", data, got)
		}
	})

}
