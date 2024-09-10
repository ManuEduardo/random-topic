package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ManuEduardo/random-topic/src/domain"
	"github.com/ManuEduardo/random-topic/src/services"
)

type Handler struct {
	_services services.IServices
}

func New(services services.IServices) *Handler {
	return &Handler{
		_services: services,
	}
}

func (handler *Handler) HandleBase(w http.ResponseWriter, r *http.Request) {
	log.Println("Using Base Handler")
}

func (handler *Handler) HandleTopicCreate(w http.ResponseWriter, r *http.Request) {
	log.Println("Using Register Handler")
}

func (handler *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Using Get User Handler")
	id_user := r.PathValue("id")
	response, err := handler._services.GetUserById(id_user)
	if err != nil {
		log.Panicln("Error getting user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err != json.NewEncoder(w).Encode(response) {
		log.Panicln("Error parsing user")
	}
}

func (handler *Handler) HandlePostUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Using Register User Handler")
	var user domain.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Panicln("Invalid JSON")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = handler._services.PostUser(user)
	if err != nil {
		log.Panicln("Error creating user")
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err != json.NewEncoder(w).Encode(user) {
		log.Panicln("Error parsing user")
	}
}

func (handler *Handler) GetRandomCard(w http.ResponseWriter, r *http.Request) {
	log.Println("Using Get Random Card Handler")
	id_user := r.PathValue("id")
	response, err := handler._services.GetRandomCard(id_user)
	if err != nil {
		log.Panicln("Error getting Random Card")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err != json.NewEncoder(w).Encode(response) {
		log.Panicln("Error parsing random card response")
	}
}

func (handler *Handler) HandlePostCard(w http.ResponseWriter, r *http.Request) {
	log.Println("Using Register Card Handler")
	var card domain.Card
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&card)
	if err != nil {
		log.Panicln("Invalid JSON")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = handler._services.PostCard(card)
	if err != nil {
		log.Panicln("Error creating card")
		http.Error(w, "Error creating card", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err != json.NewEncoder(w).Encode(card) {
		log.Panicln("Error parsing card")
	}
}
