package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	
	mux.Handle("/static/", http.StripPrefix("/static", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(writer http.ResponseWriter, request *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads(); if err == nil {
		tempaltes.ExecuteTemplate(writer, "layout", threads)
	}
}

func err(writer http.ResponseWriter, request *http.Request) {
}

func login(writer http.ResponseWriter, request *http.Request) {
}

func signup(writer http.ResponseWriter, request *http.Request) {
}

func signupAccount(writer http.ResponseWriter, request *http.Request) {
}

func authenticate(writer http.ResponseWriter, request *http.Request) {
}

func newThread(writer http.ResponseWriter, request *http.Request) {
}

func createThread(writer http.ResponseWriter, request *http.Request) {
}

func postThread(writer http.ResponseWriter, request *http.Request) {
}

func readThread(writer http.ResponseWriter, request *http.Request) {
}
