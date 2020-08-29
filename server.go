package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/urfave/negroni"
	"net/http"
)

var upgrader = websocket.Upgrader{}


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
