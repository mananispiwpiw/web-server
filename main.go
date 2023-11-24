package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello Rafi")
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful ")
	task := r.FormValue("task")
	fmt.Fprintf(w, "Task = %s\n", task)
}

func main() {
	fileServer := http.FileServer(http.Dir("./views"))
	http.Handle("/", fileServer)
	http.HandleFunc("/task", taskHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server start at port : 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
