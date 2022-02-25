package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s StubPlayerStore) RecordWin(name string) {
	s.scores[name]++
}

func TestHttp(t *testing.T) {
	store := StubPlayerStore{map[string]int{
		"rayliao": 20,
		"gg":      10,
	}}
	server := &PlayerServer{store}
	t.Run("returns rayliao's score", func(t *testing.T) {
		request := newGetScoreRequest("rayliao")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns gg's score", func(t *testing.T) {
		request := newGetScoreRequest("gg")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("cc")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{make(map[string]int)}

	server := &PlayerServer{&store}

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "jonliao"

		request := newPostWinRequest(player)
		responsePost := httptest.NewRecorder()

		server.ServeHTTP(responsePost, request)
		server.ServeHTTP(responsePost, request)

		assertStatus(t, responsePost.Code, http.StatusAccepted)
		responseGet := httptest.NewRecorder()

		server.ServeHTTP(responseGet, newGetScoreRequest(player))

		if len(store.scores) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.scores), 1)
		}

		if responseGet.Body.String() != "2" {
			t.Errorf("did not store correct winner got %d want %d", store.scores[player], 2)
		}
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Sam"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
	}
}
