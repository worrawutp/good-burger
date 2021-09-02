package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ListBurgerHandler(w http.ResponseWriter, r *http.Request) {

}

func BurgerHandler(w http.ResponseWriter, r *http.Request) {

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/about.html")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/404.html")
}

func main() {
	fmt.Println("vim-go")

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Go")
	})
	router.HandleFunc("/about", AboutHandler)
	router.HandleFunc("/burgers", ListBurgerHandler)
	router.HandleFunc("/burgers/{id: [0-9]+}", BurgerHandler)
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	http.Handle("/", router)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/css"))))
	http.ListenAndServe(":8000", nil)
}
