package main

import (
	"gogpt/api"
	"log"
	"net/http"
	"os"
)

func main() {

	mux := api.NewMux()

	addr := ":3000"
	log.Println("starting server on addr:", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
