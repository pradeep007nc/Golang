package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/", apiHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "welcome nigger")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	data := "some data"
	fmt.Println(w, data)
}
