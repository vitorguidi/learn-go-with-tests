package http_server

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("Should return 404 when there is no data in store", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "players/vitor", nil)
		response := httptest.NewRecorder()

		store := &InMemoryStore{data: make(map[string]string)}
		handler := &PlayerServer{store: store}

		handler.ServeHTTP(response, request)

		desiredStatusCode := http.StatusNotFound
		gotStatusCode := response.Code

		if gotStatusCode != desiredStatusCode {
			t.Errorf("got status code %d, want %d", gotStatusCode, desiredStatusCode)
		}

	})

	t.Run("should correctly return vitor's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "players/vitor", nil)
		response := httptest.NewRecorder()

		store := &InMemoryStore{data: map[string]string{"vitor": "20"}}
		handler := &PlayerServer{store: store}

		handler.ServeHTTP(response, request)

		gotResult := response.Body.String()
		gotStatusCode := response.Code
		desiredResult := "20"
		desiredStatusCode := http.StatusOK

		if gotStatusCode != desiredStatusCode {
			t.Errorf("got status code %d, want %d", gotStatusCode, desiredStatusCode)
		}
		if gotResult != desiredResult {
			t.Errorf("got %s, want %s", gotResult, desiredResult)
		}
	})
}

func TestPut(t *testing.T) {
	t.Run("should correctly persist score on a put request", func(t *testing.T) {
		desiredVal := "20"
		buf := bytes.Buffer{}
		fmt.Fprint(&buf, desiredVal)
		request, _ := http.NewRequest(http.MethodPut, "players/vitor", &buf)
		response := httptest.NewRecorder()

		store := &SpyStore{lastPersistedVal: "", wasPersisted: false}
		handler := &PlayerServer{store: store}

		handler.ServeHTTP(response, request)

		gotStatusCode := response.Code
		desiredStatusCode := http.StatusOK

		if gotStatusCode != desiredStatusCode {
			t.Errorf("got status code %d, want %d", gotStatusCode, desiredStatusCode)
		}
		if !store.wasPersisted || store.lastPersistedVal != desiredVal {
			t.Errorf("failed to persist to store in put request: expected %s to be persisted, it was not", desiredVal)
		}
	})

	t.Run("should return internal error when persistence fails", func(t *testing.T) {
		desiredVal := "20"
		buf := bytes.Buffer{}
		fmt.Fprint(&buf, desiredVal)
		request, _ := http.NewRequest(http.MethodPut, "players/vitor", &buf)
		response := httptest.NewRecorder()

		store := &StubStore{returnErr: PersistFailedError, returnVal: ""}
		handler := &PlayerServer{store: store}

		handler.ServeHTTP(response, request)

		gotStatusCode := response.Code
		desiredStatusCode := http.StatusInternalServerError

		if gotStatusCode != desiredStatusCode {
			t.Errorf("got status code %d, want %d", gotStatusCode, desiredStatusCode)
		}
	})
}
