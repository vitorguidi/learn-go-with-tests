package http_server

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var NotFoundError = errors.New("Could not find data for player")
var PersistFailedError = errors.New("Backend store was not able to persist")

type Store interface {
	Get(player string) (string, error)
	Put(player, score string) error
}

type SpyStore struct {
	wasPersisted     bool
	lastPersistedVal string
}

func (s *SpyStore) Get(player string) (string, error) {
	if s.wasPersisted {
		return s.lastPersistedVal, nil
	}
	return "", NotFoundError
}

func (s *SpyStore) Put(player, score string) error {
	s.wasPersisted = true
	s.lastPersistedVal = score
	return nil
}

type StubStore struct {
	returnVal string
	returnErr error
}

func (s *StubStore) Get(player string) (string, error) {
	return s.returnVal, s.returnErr
}

func (s *StubStore) Put(player, score string) error {
	return s.returnErr
}

type InMemoryStore struct {
	data map[string]string
}

func (s *InMemoryStore) Get(player string) (string, error) {
	score, found := s.data[player]
	if !found {
		return "", NotFoundError
	}
	return score, nil
}

func (s *InMemoryStore) Put(player, score string) error {
	s.data[player] = score
	return nil
}

type PlayerServer struct {
	store Store
}

func (s *PlayerServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "players/")
	switch request.Method {
	case http.MethodGet:
		s.handleGet(writer, player)
	case http.MethodPut:
		scoreBytes, _ := io.ReadAll(request.Body)
		score := string(scoreBytes)
		s.handlePut(writer, player, score)
	default:
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (s *PlayerServer) handleGet(writer http.ResponseWriter, player string) {
	score, err := s.store.Get(player)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprintf(writer, score)
}

func (s *PlayerServer) handlePut(writer http.ResponseWriter, player, score string) {
	err := s.store.Put(player, score)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	store := &InMemoryStore{data: make(map[string]string)}
	handler := &PlayerServer{store: store}
	log.Fatal(http.ListenAndServe(":5000", handler))
}
