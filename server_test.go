package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	// "encoding/json"
	// "github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:3000", nil)
	if err != nil {
		t.Fatalf("Could not create request : %v", err)
	}
	rec := httptest.NewRecorder()
	homeHandler(rec, req)
	res := rec.Result()
  assert.Equal(t, http.StatusOK, res.StatusCode, "status not ok got status as : %s ",  res.StatusCode)

}


func TestNewGameHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(newGameHandler))
	defer server.Close()

	t.Run("Should return valid json of leaderboard", func(t *testing.T) {

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

		err = ws.WriteMessage(websocket.TextMessage, []byte("0"))
		if err != nil {
			t.Fatalf("could not send message ws connection %v", err)
    }
		_, recvData, err := ws.ReadMessage()
		if err != nil {
			t.Fatalf("could read message from server due to %v", err)
		}
		fmt.Println(string(recvData))
		// ws.Close()
	})

	t.Run("Should generate grid for difficulty level easy", func(t *testing.T) {

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

		if err != nil {
			t.Fatalf("could not connect to websocket due to %v", err)
		}

		err = ws.WriteMessage(websocket.TextMessage, []byte("0"))
		if err != nil {
			t.Fatalf("could not send message ws connection %v", err)
		}

		_, _, err = ws.ReadMessage()

		var flag int = 0
		_, recvData, err := ws.ReadMessage()

		if err != nil {
			t.Fatalf("could read message from server due to %v", err)
		}
		for _, value := range recvData {
			if string(value) == "0" {
				flag = 1
			}
		}
		if flag == 0 {
			t.Errorf("expected complete grid but got zero value in grid")
		}
	// ws.Close()
	})

}

func TestCheckWin(t *testing.T) {

  server := httptest.NewServer(http.HandlerFunc(newGameHandler))
	defer server.Close()

  t.Run("Should return true if user won the game", func(t *testing.T) {
		Game := Sudoku{}
    gameLevel := "1"
    Game.initializeGame(9, 3, gameLevel)
    Game.createPuzzle(gameLevel)
    Game.answerGrid = replicateOriginalGrid(Game.sudokuGrid)
    str := Game.generateStream()
		fmt.Println(str)
    winStatus := Game.checkAnswer()
		assert.Equal(t, false, winStatus)
	})

	t.Run("Should return true if user won the game", func(t *testing.T) {
		Game := Sudoku{}
    gameLevel := "1"
    Game.initializeGame(9, 3, gameLevel)
    Game.createPuzzle(gameLevel)
    Game.answerGrid = replicateOriginalGrid(Game.sudokuGrid)
    str := Game.generateStream()
		fmt.Println(str)
    winStatus := Game.checkAnswer()
		assert.Equal(t, false, winStatus)
	})

}
