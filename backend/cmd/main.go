package main

import (
	"log"
	"net/http"
	"os"

	"github.com/DiegoLinares11/Lab-6-Backend-Only/internal/handlers"
	"github.com/DiegoLinares11/Lab-6-Backend-Only/internal/storage"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Configurar almacenamiento
	postgresStorage := storage.NewPostgresStorage()

	// Configurar manejadores
	matchHandler := handlers.NewMatchHandler(postgresStorage)

	// Configurar router
	r := mux.NewRouter()

	// Configurar rutas
	r.HandleFunc("/api/matches", matchHandler.GetAllMatches).Methods("GET")
	r.HandleFunc("/api/matches/{id}", matchHandler.GetMatchByID).Methods("GET")
	r.HandleFunc("/api/matches", matchHandler.CreateMatch).Methods("POST")
	r.HandleFunc("/api/matches/{id}", matchHandler.UpdateMatch).Methods("PUT")
	r.HandleFunc("/api/matches/{id}", matchHandler.DeleteMatch).Methods("DELETE")

	// Rutas PATCH adicionales
	r.HandleFunc("/api/matches/{id}/goals", matchHandler.RegisterGoal).Methods("PATCH")
	r.HandleFunc("/api/matches/{id}/yellowcards", matchHandler.RegisterYellowCard).Methods("PATCH")
	r.HandleFunc("/api/matches/{id}/redcards", matchHandler.RegisterRedCard).Methods("PATCH")
	r.HandleFunc("/api/matches/{id}/extratime", matchHandler.SetExtraTime).Methods("PATCH")

	// Configurar CORS
	handler := cors.Default().Handler(r)

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	log.Printf("Servidor iniciado en el puerto %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
