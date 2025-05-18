package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type User struct {
	ID int
	Name string
	Email string
	Age int
}

func Read(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	defer rows.Close()
	data := make([]User, 0)

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
		if err != nil {
			fmt.Println("Server failed to handle", err)
			return
		}
		data = append(data, user)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(data)
}

func Create(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	user := User{}
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	_, err = db.Exec("INSERT INTO users (name, email, age) VALUES ($1, $2, $3)", user.Name, user.Email, user.Age)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	writer.WriteHeader(http.StatusCreated)
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