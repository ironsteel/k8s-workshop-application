package main

import (
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Print("hello k8s!")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
