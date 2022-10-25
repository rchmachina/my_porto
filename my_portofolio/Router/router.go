package Router

import (
	. "my_porto/Controlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Router()  {
	routing := mux.NewRouter()

	routing.HandleFunc("/", Home).Methods("GET")
	routing.HandleFunc("/test", Test).Methods("GET")
	http.ListenAndServe(":9000", routing)

	routing.Handle("/static/", 
	http.StripPrefix("/static/", 
		http.FileServer(http.Dir("styling"))))

}