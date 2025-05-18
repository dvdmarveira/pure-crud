package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func Read(writer http.ResponseWriter, response *http.Request) {

}

func Create(writer http.ResponseWriter, response *http.Request) {

}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:2513@postgres/crudgo?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to database.")

}

func main() {
	http.HandleFunc("/users", Read)
	http.HandleFunc("/users/create", Create)
	http.ListenAndServe(":8585", nil)
}