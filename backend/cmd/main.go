// @title La Liga Tracker API
// @version 1.0
// @description API para la gestión de partidos de fútbol.
// @host localhost:8080
// @BasePath /api
// @schemes http
// @contact.name Diego Linares
// @contact.url https://github.com/DiegoLinares11/Lab-6-Backend-Only
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/DiegoLinares11/Lab-6-Backend-Only/docs" // Importa la documentación generada por Swaggo
	"github.com/DiegoLinares11/Lab-6-Backend-Only/internal/handlers"
	"github.com/DiegoLinares11/Lab-6-Backend-Only/internal/storage"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger" // Para servir la documentación Swagger
)

func main() {
	// Configurar almacenamiento
	postgresStorage := storage.NewPostgresStorage()

	// Configurar manejadores
	matchHandler := handlers.NewMatchHandler(postgresStorage)

	// Configurar router
	r := mux.NewRouter()

	// Configurar rutas
	r.HandleFunc("/api/matches", matchHandler.GetAllMatches).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/matches/{id}", matchHandler.GetMatchByID).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/matches", matchHandler.CreateMatch).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/matches/{id}", matchHandler.UpdateMatch).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/matches/{id}", matchHandler.DeleteMatch).Methods("DELETE", "OPTIONS")

	// Rutas PATCH adicionales
	r.HandleFunc("/api/matches/{id}/goals", matchHandler.RegisterGoal).Methods("PATCH", "OPTIONS")
	r.HandleFunc("/api/matches/{id}/yellowcards", matchHandler.RegisterYellowCard).Methods("PATCH", "OPTIONS")
	r.HandleFunc("/api/matches/{id}/redcards", matchHandler.RegisterRedCard).Methods("PATCH", "OPTIONS")
	r.HandleFunc("/api/matches/{id}/extratime", matchHandler.SetExtraTime).Methods("PATCH", "OPTIONS")

	// Ruta para Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Configurar CORS con opciones más permisivas
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permitir cualquier origen
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Envolver el router con el middleware CORS
	handler := c.Handler(r)

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciado en el puerto %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
