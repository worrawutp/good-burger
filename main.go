package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

type Menu struct {
	Id          int
	Name        string
	Description string
	Price       int
}

var Conn *pgx.Conn

// Conn is initialize with zero value

func ListBurgersHandler(w http.ResponseWriter, r *http.Request) {
	var menus []Menu
	rows, err := Conn.Query(context.Background(), "select * from menus")

	for rows.Next() {
		var m = Menu{}
		rows.Scan(&m.Id, &m.Name, &m.Description, &m.Price)
		menus = append(menus, m)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", menus)
	menuTemplate, err := template.ParseFiles("views/menus/index.html")
	if err != nil {
		http.ServeFile(w, r, "public/500.html")
		return
	}
	menuTemplate.Execute(w, menus)
}

func BurgersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 0, 64)

	var menu = Menu{}
	err := Conn.QueryRow(context.Background(), "select * from menus where id=$1", id).Scan(&menu.Id, &menu.Name, &menu.Description, &menu.Price)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", menu)
	menuTemplate, err := template.ParseFiles("views/menus/show.html")
	if err != nil {
		http.ServeFile(w, r, "public/500.html")
		return
	}
	menuTemplate.Execute(w, menu)

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/about.html")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/404.html")
}

func main() {
	// Load environment variables
	env := os.Getenv("GOOD_BURGER_ENV")
	if "" == env {
		env = "development"
	}
	godotenv.Load(".env." + env)
	godotenv.Load()
	fmt.Println(os.Getenv("DATABASE_URL"))

	// Connect database
	var err error
	Conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Problem! cannot connect to database")
	}
	fmt.Println(Conn)

	// Routing
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Go")
	})
	router.HandleFunc("/about", AboutHandler)
	router.HandleFunc("/burgers", ListBurgersHandler)
	router.HandleFunc("/burgers/{id:[0-9]+}", BurgersHandler)
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	http.Handle("/", router)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/css"))))
	http.ListenAndServe(":8000", nil)
}
