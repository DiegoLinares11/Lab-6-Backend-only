package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/DiegoLinares11/Lab-6-Backend-Only/internal/models"
	"github.com/DiegoLinares11/Lab-6-Backend-Only/internal/storage"
	"github.com/gorilla/mux"
)

type MatchHandler struct {
	storage *storage.PostgresStorage
}

func NewMatchHandler(ps *storage.PostgresStorage) *MatchHandler {
	return &MatchHandler{storage: ps}
}

// GetAllMatches obtiene todos los partidos registrados.
//
// @Summary Obtiene la lista de partidos
// @Description Retorna todos los partidos almacenados en la base de datos
// @Tags Matches
// @Produce json
// @Success 200 {array} models.Match
// @Failure 500 {object} map[string]string
// @Router /matches [get]
func (mh *MatchHandler) GetAllMatches(w http.ResponseWriter, r *http.Request) {
	// Establecer encabezados CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	matches, err := mh.storage.GetAllMatches()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}

// GetMatchByID obtiene un partido por su ID.
//
// @Summary Obtiene un partido
// @Description Retorna la información de un partido específico por su ID
// @Tags Matches
// @Produce json
// @Param id path int true "ID del partido"
// @Success 200 {object} models.Match
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /matches/{id} [get]
func (mh *MatchHandler) GetMatchByID(w http.ResponseWriter, r *http.Request) {

	// Establecer encabezados CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	match, err := mh.storage.GetMatchByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}

func (mh *MatchHandler) CreateMatch(w http.ResponseWriter, r *http.Request) {
	// Establecer encabezados CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	var match models.Match
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&match); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Convertir string a time.Time y normalizar a UTC
	matchDate, err := time.Parse("2006-01-02", match.MatchDate)
	if err != nil {
		http.Error(w, "Invalid date format, expected YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	match.MatchDate = matchDate.UTC().Format("2006-01-02") // Convertir a UTC y formatear como string

	if err := mh.storage.CreateMatch(&match); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(match)
}

func (mh *MatchHandler) UpdateMatch(w http.ResponseWriter, r *http.Request) {
	// Establecer encabezados CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	var match models.Match
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&match); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	match.ID = id

	// Convertir string a time.Time y normalizar a UTC
	matchDate, err := time.Parse("2006-01-02", match.MatchDate)
	if err != nil {
		http.Error(w, "Invalid date format, expected YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	match.MatchDate = matchDate.UTC().Format("2006-01-02") // Convertir a UTC y formatear como string

	if err := mh.storage.UpdateMatch(&match); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}

func (mh *MatchHandler) DeleteMatch(w http.ResponseWriter, r *http.Request) {
	// Establecer encabezados CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	if err := mh.storage.DeleteMatch(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (mh *MatchHandler) RegisterGoal(w http.ResponseWriter, r *http.Request) {
	// Establecer encabezados CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	if err := mh.storage.RegisterGoal(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (mh *MatchHandler) RegisterYellowCard(w http.ResponseWriter, r *http.Request) {
	// Establecer encabezados CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	if err := mh.storage.RegisterYellowCard(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (mh *MatchHandler) RegisterRedCard(w http.ResponseWriter, r *http.Request) {
	// Establecer encabezados CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	if err := mh.storage.RegisterRedCard(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (mh *MatchHandler) SetExtraTime(w http.ResponseWriter, r *http.Request) {
	// Establecer encabezados CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	if err := mh.storage.SetExtraTime(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Método para manejar solicitudes OPTIONS
func (mh *MatchHandler) OptionsHandler(w http.ResponseWriter, r *http.Request) {
	// Establecer encabezados CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
}
