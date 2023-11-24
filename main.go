package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 Not Found",http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"Mehod is not supported",http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w,"Hello Rafi")
}

func taskHandler(){

}

func main(){
	fileServer := http.FileServer(http.Dir("./views"))
	http.Handle("/",fileServer)
	//http.HandleFunc("/task",taskHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Server start at port : 8080")
	if err := http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err)
	}
}