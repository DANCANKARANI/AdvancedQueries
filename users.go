package main

import (
	//"gorm.io/gorm"
	//"gorm.io/driver/sqlite"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
var db *gorm.DB
type User struct{
	gorm.Model
	Name string
	Age string
	Email string
}
type ApiUser struct{
	gorm.Model
	Name string
	Age string
}
func InitialMigration(){
	db,err:=gorm.Open("sqlite3","user_db")
	if err!=nil{
		panic("failed to connect to the database")
	}
	defer db.Close()
	db.AutoMigrate(&User{}).AutoMigrate(&ApiUser{})
	db.Create(&User{}).Create(&ApiUser{})

}
func AddUser(w http.ResponseWriter, r *http.Request){
	db,err:=gorm.Open("sqlite3","user_db")
	if err!=nil{
		panic("failed to connect to the database")
	}
	defer db.Close()
	vars:=mux.Vars(r)
	name:=vars["name"]
	age:=vars["age"]
	email:=vars["email"]

	db.AutoMigrate(&User{})
	db.Create(&User{Name: name,Age: age,Email: email})
	db.Create(&ApiUser{Name: name})

	fmt.Fprintf(w,"posting users")
}
func FindAllUsers(w http.ResponseWriter, r *http.Request){
	db,err:=gorm.Open("sqlite3","user_db")
	if err!=nil {
		panic("failed to connect to the database")
	}
	defer db.Close()
	var user []User
	//var apiUser ApiUser
	db.Find(&user)
	//db.Find(&apiUser)
	json.NewEncoder(w).Encode(user)
	//json.NewEncoder(w).Encode(apiUser)
	fmt.Fprintf(w,"getting all users")
}

func HelloFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"hello user")
}
func AddApiUser(w http.ResponseWriter, r *http.Request){
	db,err:=gorm.Open("sqlite3","user_db")
	if err!=nil{
		panic("faied to connect to the datase")
	}
	defer db.Close()
	vars:=mux.Vars(r)
	name:=vars["name"]
	age:=vars["age"]
	db.AutoMigrate(&ApiUser{})
	db.Create(&ApiUser{Name:name,Age: age})
	fmt.Fprintf(w,"dancan")
}
func GetApiUsers(w http.ResponseWriter, r *http.Request){
	db,err:=gorm.Open("sqlite3","user_db")
	if err!=nil {
		panic("failed to connect to the database")
	}
	defer db.Close()
	//var user []User
	var apiUser []ApiUser
	//db.Find(&user)
	db.Find(&apiUser)
	json.NewEncoder(w).Encode(apiUser)
	//json.NewEncoder(w).Encode(apiUser)
	fmt.Fprintf(w,"getting all users")
}

