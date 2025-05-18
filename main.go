package main

import "net/http"

func Read(writer http.ResponseWriter, response *http.Request) {

}

func Create(writer http.ResponseWriter, response *http.Request) {

}

func main() {
	http.HandleFunc("/users", Read)
	http.HandleFunc("/users/create", Create)
	http.ListenAndServe(":8585", nil)
}