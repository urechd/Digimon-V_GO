package routes

import (
	"fmt"
	"net/http"

	"github.com/urechd/digimon/internal/database"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/data", apiDataHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the homepage!")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	digiInfo := database.GetDigimonInfoById(3)
	fmt.Fprint(w, string(digiInfo.Name))
}
