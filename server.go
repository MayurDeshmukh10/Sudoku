package main

import (
	"time"
	"strconv"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/urfave/negroni"
	"net/http"
)

var upgrader = websocket.Upgrader{}
var DBUSER = env.Getenv("DBUSER")
var DBPASS = env.Getenv("DBPASS")
var DBNAME = env.Getenv("DBNAME")

//Struct to hold Player data
type User struct{
	Name string
	Time string
}

//function to get database connection
func getConnection() (conn *sql.DB){
	conn, err := sql.Open("mysql",DBUSER+":"+DBPASS+"@/"+DBNAME)
	if err != nil{
		panic(err)
	}
	return 
}

//function to get top players to display on leaderboard 
func getPlayers(conn *sql.DB, grid_size int, difficulty int) (users []User){
	query := "SELECT Name, Time FROM Leaderboard WHERE Sudoku_size = ? AND Difficulty = ?"
	result, err := conn.Query(query, grid_size, difficulty)
	
	if err != nil{
		panic(err)
	}
	
	defer result.Close()
	
	for result.Next(){
		var user User
		err = result.Scan(&user.Name,&user.Time)
		if err != nil{
			panic(err)
		}
		users = append(users,user)
	}
	return 
}

//function to add player record in database
func addPlayer(conn *sql.DB, name string, duration time.Duration, difficulty int, grid_size int){
	//converting duration into HH:MM:SS format
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) - hours*60
	seconds := int(duration.Seconds()) - hours*60*60 - minutes*60
	time := strconv.Itoa(hours)+":"+strconv.Itoa(minutes)+":"+strconv.Itoa(seconds)
	
	query := "INSERT INTO Leaderboard (Name, Time, Difficulty, Sudoku_size) VALUES (?, ?, ?, ?)"
	result, err := conn.Query(query, name, time, difficulty, grid_size) 
	if err != nil {
		panic(err)
	}
	defer result.Close()
}



func homeHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "game.html")
}

func newGameHandler(rw http.ResponseWriter, req *http.Request) {

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
