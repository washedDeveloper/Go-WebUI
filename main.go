package main

import (
	"net/http"
	"os"
	"os/exec"

	d "./backend/dependencies"
)

func main() {
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/", fs)
	cmnd := exec.Command(d.GetPath(), "--app=http://localhost:3000", "--new-window")
	cmnd.Start()
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		os.Exit(3)
	}
}
