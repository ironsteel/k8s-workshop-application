package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ironsteel/k8s-workshop-application/internal/diagnostics"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Print("hello k8s!")

	router := mux.NewRouter()
	router.HandleFunc("/", hello)
	router.HandleFunc("/healthz", hello)

	go func() {
		if err := http.ListenAndServe(":8085", diagnostics.New()); err != nil {
			log.Fatal(err)
		}
	}()

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
