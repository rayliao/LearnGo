package main

// func TestRecordingWinsAndRetrievingThem(t *testing.T) {
// 	database, cleanDatabase := createTempFile(t, `[]`)
// 	defer cleanDatabase()
// 	store, err := NewFileSystemPlayerStore(database)

// 	store := NewInMemoryPlayerStore()
// 	server := NewPlayerServer(store)
// 	player := "Sam"

// 	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
// 	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
// 	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

// 	t.Run("get score", func(t *testing.T) {
// 		response := httptest.NewRecorder()
// 		server.ServeHTTP(response, newGetScoreRequest(player))
// 		assertStatus(t, response.Code, http.StatusOK)
// 		assertResponseBody(t, response.Body.String(), "3")
// 	})

// 	t.Run("get league", func(t *testing.T) {
// 		response := httptest.NewRecorder()
// 		server.ServeHTTP(response, newLeagueRequest())
// 		assertStatus(t, response.Code, http.StatusOK)

// 		got := getLeagueFromResponse(t, response.Body)
// 		want := []Player{
// 			{"Sam", 3},
// 		}

// 		assertLeague(t, got, want)
// 	})
// }
