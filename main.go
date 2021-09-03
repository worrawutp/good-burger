package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/worrawutp/good-burger/initializers"

	. "github.com/worrawutp/good-burger/handlers"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/about.html")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/404.html")
}

func main() {
	initializers.LoadDotEnv()
	initializers.InitDatabase()

	// Routing
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Go")
	})
	router.HandleFunc("/about", AboutHandler)
	router.HandleFunc("/menus", ListMenusHandler).Methods("GET")
	router.HandleFunc("/menus", CreateMenusHandler).Methods("POST")
	router.HandleFunc("/menus/{id:[0-9]+}", MenuHandler)
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	http.Handle("/", router)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/css"))))
	http.ListenAndServe(":8000", nil)
}
