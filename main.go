package main

import (
	"log"
	"net/http"
	"os/exec"
	"time"

	d "./backend/dependencies"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", ws)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("frontend"))))
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	cmnd := exec.Command(d.GetPath(), "--app=http://127.0.0.1:3000", "--new-window")
	cmnd.Start()
	srv.ListenAndServe()
}

var upgrader = websocket.Upgrader{}

func ws(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
