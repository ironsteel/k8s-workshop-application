package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ironsteel/k8s-workshop-application/internal/diagnostics"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
	log.Println("Hello handler was called")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Print("Server starting...")

	port := os.Getenv("PORT")
	diagPort := os.Getenv("DIAG_PORT")

	if len(port) == 0 {
		log.Fatal("The application PORT must be set")
	}
	if len(diagPort) == 0 {
		log.Fatal("The diagnostics DIAG_PORT must be set")
	}

	router := mux.NewRouter()
	router.HandleFunc("/", hello)
	router.HandleFunc("/healthz", hello)

	go func() {
		log.Print("Diagnostics server is praparing...")
		server := &http.Server{
			Addr:    ":" + diagPort,
			Handler: diagnostics.New(),
		}
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Print("Application server is praparing...")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}

}
