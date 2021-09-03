package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/worrawutp/good-burger/initializers"

	. "github.com/worrawutp/good-burger/structs"
)

func ListMenusHandler(w http.ResponseWriter, r *http.Request) {
	var menus []Menu
	rows, err := initializers.Conn.Query(context.Background(), "select * from menus")

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

func MenusHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 0, 64)

	var menu = Menu{}
	err := initializers.Conn.QueryRow(context.Background(), "select * from menus where id=$1", id).Scan(&menu.Id, &menu.Name, &menu.Description, &menu.Price)
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
