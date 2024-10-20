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
	"github.com/rs/cors"
)

func main() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could not be loaded")
	}

	// Configuración de la base de datos y servidor
	portServer := os.Getenv("PORT")
	dbName := os.Getenv("BD_NAME")
	dbUser := os.Getenv("BD_USER")
	dbPass := os.Getenv("BD_PASSWORD")

	urlDb := fmt.Sprintf("postgres://%v:%v@localhost:5432/%v", dbUser, dbPass, dbName)

	// Inicializar la base de datos
	dbInstance := infraestructure.New(urlDb)
	err = dbInstance.InitDB()
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
		return
	}
	defer dbInstance.CloseDB()

	// Inicializar repositorios, servicios y handlers
	repo := repository.New(dbInstance)
	svc := services.New(repo)
	handler := handlers.New(svc)

	// Configurar el enrutador
	router := http.NewServeMux()

	// Enrutar solicitudes SOAP
	router.HandleFunc("/soap", handler.SoapHandler)

	// Configurar CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://your-frontend-domain.com"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	// Iniciar el servidor
	serverAddress := fmt.Sprintf(":%v", portServer)
	log.Printf("Listening on %v\n", serverAddress)
	err = http.ListenAndServe(serverAddress, corsHandler.Handler(router))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
