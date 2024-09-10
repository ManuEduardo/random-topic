package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ManuEduardo/random-topic/src/handlers"
	"github.com/ManuEduardo/random-topic/src/infraestructure"
	"github.com/ManuEduardo/random-topic/src/repository"
	"github.com/ManuEduardo/random-topic/src/services"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could not be loaded")
	}

	portServer := os.Getenv("PORT")
	dbName := os.Getenv("BD_NAME")
	dbUser := os.Getenv("BD_USER")
	dbPass := os.Getenv("BD_PASSWORD")

	urlDb := fmt.Sprintf("postgres://%v:%v@localhost:5432/%v", dbUser, dbPass, dbName)

	err = infraestructure.InitDB(urlDb)
	if err != nil {
		return
	}
	defer infraestructure.CloseDB()

	repository := repository.New(infraestructure.GetDBPool())
	services := services.New(repository)
	handlers := handlers.New(services)

	router := http.NewServeMux()

	router.HandleFunc("GET /topic/create", handlers.HandleTopicCreate)
	router.HandleFunc("GET /user/{id}", handlers.HandleGetUser)
	router.HandleFunc("POST /user", handlers.HandlePostUser)
	router.HandleFunc("GET /", handlers.HandleBase)

	log.Printf("Listening on %v\n", fmt.Sprintf("localhost:%v", portServer))
	err = http.ListenAndServe(fmt.Sprintf(":%v", portServer), router)
	log.Fatalln(err.Error())
}

// func initStorage(db *sql.DB) {
// 	err := db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to Database!")
// }
