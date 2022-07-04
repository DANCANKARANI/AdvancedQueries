package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MyRouter(){
	mux:=mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/",HelloFunc).Methods("GET")
	mux.HandleFunc("/users",FindAllUsers).Methods("GET")
	mux.HandleFunc("/api",GetApiUsers).Methods("GET")
	mux.HandleFunc("/{age}/{name}/{email}",AddUser).Methods("POST")
	mux.HandleFunc("/{age}/{name}",AddApiUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080",mux))
}
func main(){
	
	 fmt.Println("Getting started...")
	 MyRouter()
	 InitialMigration()
}
