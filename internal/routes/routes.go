package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/urechd/digimon/internal/database"
	"github.com/urechd/digimon/internal/models"
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
	digiInfo := []models.Digimon{}
	dat := database.GetFileData("./resources/digimon.json")
	err := json.Unmarshal(dat, &digiInfo)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(dat))
}
