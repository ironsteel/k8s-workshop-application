package diagnostics

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/healthz", healthz)
	router.HandleFunc("/info", info)
	return router

}
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
func info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
