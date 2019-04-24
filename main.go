package main

import (
	"PhoneBook/model"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pleycpl/godotenv"
	"log"
	"net/http"
)

func main() {
	InitDatabase()
	StartServer()
}

func StartServer() {
	env := godotenv.Godotenv("system_variable_env")
	PORT := env["SERVER_PORT"]

	router := mux.NewRouter()
	router.HandleFunc("/", GetAll).Methods("GET")
	fmt.Println("Start Server On Port... " + PORT)
	http.ListenAndServe(":"+PORT, router)
}
func GetAll(writer http.ResponseWriter, request *http.Request) {
	var att = "<h1>" + request.URL.Host + "</h1>" + "\n"
	vars := mux.Vars(request)
	att += vars["token"]
	att += vars["username"]
	fmt.Fprintf(writer, "<h1>"+"Salam Azizam Fisrt Web Server"+"</h1>", att)
}
func InitDatabase() {

	env := godotenv.Godotenv("system_variable_env")
	//fmt.Println(env)
	DB_TYPE := env["DATABASE_TYPE"]
	DB_HOST := env["DB_HOST"]
	DB_PORT := env["DB_PORT"]
	DB_USERNAME := env["DB_USERNAME"]
	DB_DATABASE := env["DB_DATABASE"]
	DB_PASSWORD := env["DB_PASSWORD"]

	db, err := gorm.Open(DB_TYPE, "host="+DB_HOST+" port="+DB_PORT+
		" user="+DB_USERNAME+" dbname="+DB_DATABASE+" password="+DB_PASSWORD)
	if err != nil {
		log.Fatal("Unable to Connect Database.")
		defer db.Close()
	}
	var user = model.User{}
	var video = model.Video{}

	if db.HasTable(model.User{}) == false {
		db.Table("users").CreateTable(&user)
		user = model.User{Username: "mahdihp", Pasword: "110", Email: "mahdihp@yahoo.com", Regfirebase: "wow"}
		log.Println("Insert Record User Rows Affected :", db.Create(&user).RowsAffected)
	}

	if db.HasTable(model.Video{}) == false {
		db.Table("videos").CreateTable(&video)
		video = model.Video{Title: "Learn C#", Description: "Learn C# language with step by step", Url: "www.google.com"}
		log.Println("Insert Record Video Rows Affected :", db.Create(&video).RowsAffected)
	}

	defer db.Close()
}
