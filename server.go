package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var upgrader = websocket.Upgrader{}
var DBUSER = env.Getenv("DBUSER")
var DBPASS = env.Getenv("DBPASS")
var DBNAME = env.Getenv("DBNAME")

//Struct to hold Player data
type User struct {
	Name string
	Time string
}

//function to get database connection
func getConnection() (conn *sql.DB) {
	conn, err := sql.Open("mysql", DBUSER+":"+DBPASS+"@/"+DBNAME)
	if err != nil {
		panic(err)
	}
	return
}

//function to get top players to display on leaderboard
func getTopPlayers(conn *sql.DB, grid_size int, difficulty int) string {
	var users []User
	query := "SELECT Name, Time FROM Leaderboard WHERE Sudoku_size = ? AND Difficulty = ? ORDER BY Time LIMIT 5"
	result, err := conn.Query(query, grid_size, difficulty)

	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		var user User
		err = result.Scan(&user.Name, &user.Time)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	return string(jsonData)
}

//function to add player record in database
func addPlayer(conn *sql.DB, name string, duration time.Duration, difficulty int, grid_size int) {
	//converting duration into HH:MM:SS format
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) - hours*60
	seconds := int(duration.Seconds()) - hours*60*60 - minutes*60
	time := strconv.Itoa(hours) + ":" + strconv.Itoa(minutes) + ":" + strconv.Itoa(seconds)

	query := "INSERT INTO Leaderboard (Name, Time, Difficulty, Sudoku_size) VALUES (?, ?, ?, ?)"
	result, err := conn.Query(query, name, time, difficulty, grid_size)
	if err != nil {
		panic(err)
	}
	defer result.Close()
}

//Generate Stream for Sending Over Web Socket
func (s *Sudoku) generateStream() string {
	var puzzleDataStream string
	for i := 0; i < s.gridSize; i++ {
		for j := 0; j < s.gridSize; j++ {
			puzzleDataStream = puzzleDataStream + strconv.Itoa(s.sudokuGrid[i][j])
		}
	}
	return puzzleDataStream
}

//Compare Puzzle with the Answer Grid
func (s *Sudoku) checkAnswer() bool {
	return (s.answerGrid == s.replicatedGrid)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "game.html")
}

func newGameHandler(rw http.ResponseWriter, req *http.Request) {
	start := time.Now()

	type Score struct {
		name string
		time []int
	}
	c, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Print("Upgrade : ", err)
	}

	// To get difficuly level from UI
	_, level, err := c.ReadMessage()

	gameLevel := string(level)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Difficulty Level : ", gameLevel)
	Game := Sudoku{}
	Game.initializeGame(9, 3, gameLevel)

	conn := getConnection()
	c.WriteMessage(websocket.TextMessage, []byte(getTopPlayers(conn, Game.gridSize, Game.gameLevel)))

	Game.createPuzzle(gameLevel)
	fmt.Println("phase 2")
	Game.answerGrid = replicateOriginalGrid(Game.sudokuGrid)
	fmt.Println("phase 3")
	str := Game.generateStream()
	fmt.Println("phase 4")
	c.WriteMessage(websocket.TextMessage, []byte(str))
	fmt.Println("phase 5")

	for {
		// score := Score{}
		var userData map[string]int

		_, recvData, err := c.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		//Extracting data from UI
		_ = json.Unmarshal(recvData, &userData)
		value := userData["value"]
		row := userData["row"]
		col := userData["col"]

		Game.answerGrid[row][col] = value

		if Game.answerGrid[row][col] != Game.replicatedGrid[row][col] {
			c.WriteMessage(websocket.TextMessage, []byte("Violation"))
		} else {
			w := Game.checkAnswer()
			fmt.Println("status : ", w)
			if Game.checkAnswer() {
				c.WriteMessage(websocket.TextMessage, []byte("WIN"))
				userTiming := time.Since(start)

				//Getting Player Name
				_, nameData, _ := c.ReadMessage()
				name := string(nameData)
				addPlayer(conn, name, userTiming, Game.gameLevel, Game.gridSize)
				break
			}
		}
	}
}

func InitRouter() (router *mux.Router) {

	router = mux.NewRouter()
	// router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./assets/"))))
	router.HandleFunc("/", homeHandler).Methods(http.MethodGet)
	router.HandleFunc("/ws", newGameHandler).Methods(http.MethodGet)

	return
}

func serverStart() {

	router := InitRouter()
	server := negroni.Classic()
	server.UseHandler(router)
	server.Run(":3000")

}

func main() {
	serverStart()
}
