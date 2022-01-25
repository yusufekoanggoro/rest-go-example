package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
	"github.com/yusufekoanggoro/rest-go-example/controllers"
	"github.com/yusufekoanggoro/rest-go-example/database"
	"github.com/yusufekoanggoro/rest-go-example/entity"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initalizeHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initDB() {
	config := database.Config{
		ServerName: "localhost",
		User:       "root",
		Password:   "",
		DB:         "rest_go_example",
	}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Person{})
}

func initalizeHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletPersonByID).Methods("DELETE")
}
