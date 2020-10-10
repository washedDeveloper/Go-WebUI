package main

import (
	"net/http"
	"os/exec"
	"time"

	d "./backend/dependencies"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("frontend"))))
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	cmnd := exec.Command(d.GetPath(), "--app=http://127.0.0.1:3000", "--new-window")
	cmnd.Start()
	srv.ListenAndServe()
}
