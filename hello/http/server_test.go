package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
	league []Player
}

func (s StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s StubPlayerStore) RecordWin(name string) {
	s.scores[name]++
}

func (s StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestHttp(t *testing.T) {
	store := StubPlayerStore{map[string]int{
		"rayliao": 20,
		"gg":      10,
	}, nil}
	server := NewPlayerServer(store)
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
	store := StubPlayerStore{make(map[string]int), nil}

	server := NewPlayerServer(&store)

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
	server := NewPlayerServer(store)
	player := "Sam"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Sam", 3},
		}

		assertLeague(t, got, want)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 33},
			{"Chris", 2},
			{"Liao", 33},
		}
		store := StubPlayerStore{nil, wantedLeague}
		server := NewPlayerServer(&store)
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := getLeagueFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)

		if response.Header().Get("content-type") != "application/json" {
			t.Errorf("response did not have content-type of application/json, got %v", response.Header())
		}
	})
}

type FileSystemStore struct {
	database io.Reader
}

func (f *FileSystemStore) GetLeague() []Player {
	league, _ := NewLeague(f.database)

	return league
}

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 40},
		]`)

		store := FileSystemStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 40},
		}

		assertLeague(t, got, want)
	})
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func getLeagueFromResponse(t *testing.T, body io.Reader) []Player {
	t.Helper()
	league, _ := NewLeague(body)
	return league
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
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
