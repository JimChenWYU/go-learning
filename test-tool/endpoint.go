package main

import (
	"go-learning/test-tool/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.Routes()

	log.Println("Listening : Started : Listening on :4000")
	_ = http.ListenAndServe(":4000", nil)
}
