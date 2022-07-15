package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Docker"))
	})
	log.Print("go docker project is running...")
	http.ListenAndServe(":8888", nil)
}
