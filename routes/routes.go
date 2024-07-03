package routes

import (
	"log"
	"net/http"

	"github.com/bianca-martinsg/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalities", controllers.GetPersonalities).Methods("Get")
	r.HandleFunc("/personalities/{id}", controllers.GetPersonality).Methods("Get")
	log.Fatal(http.ListenAndServe(":8080", r))
}
