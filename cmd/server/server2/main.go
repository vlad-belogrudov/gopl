package main

import (
	"log"
	"net/http"

	"github.com/vlad-belogrudov/gopl/pkg/lissajous"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
