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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Base handler running"))
}

func (handler *Handler) HandleTopicCreate(w http.ResponseWriter, r *http.Request) {
	log.Println("Using Register Handler")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not implemented"))
}

// func (handler *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Using Get User Handler")
// 	id_user := r.PathValue("id")
// 	response, err := handler._services.GetUserById(id_user)

// 	if err != nil {
// 		log.Println("Error getting user")
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}

//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK) // 200 OK
//		if err != json.NewEncoder(w).Encode(response) {
//			log.Println("Error parsing user")
//			http.Error(w, "Error encoding user data", http.StatusInternalServerError)
//		}
//	}
func (handler *Handler) HandlePostUserr(w http.ResponseWriter, r *http.Request) {
	log.Println("Using Register User Handler")
	var user domain.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("Invalid JSON")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = handler._services.PostUser(user)
	if err != nil {
		log.Println("Error creating user")
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	if err != json.NewEncoder(w).Encode(user) {
		log.Println("Error parsing user")
		http.Error(w, "Error encoding user data", http.StatusInternalServerError)
	}
}

// func (handler *Handler) GetRandomCard(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Using Get Random Card Handler")
// 	id_user := r.PathValue("id")
// 	response, err := handler._services.GetRandomCard(id_user)

// 	if err != nil {
// 		log.Println("Error getting Random Card")
// 		http.Error(w, "Random card not found", http.StatusNotFound)
// 		return
// 	}

//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK) // 200 OK
//		if err != json.NewEncoder(w).Encode(response) {
//			log.Println("Error parsing random card response")
//			http.Error(w, "Error encoding card data", http.StatusInternalServerError)
//		}
//	}

// func (handler *Handler) HandlePostCard(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Using Registe Card Handler")
// 	var card domain.Card
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&card)

// 	if err != nil {
// 		log.Println("Invalid JSON")
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}

// 	err = handler._services.PostCard(card)
// 	if err != nil {
// 		log.Println("Error creating card")
// 		http.Error(w, "Error creating card", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	if err != json.NewEncoder(w).Encode(card) {
// 		log.Println("Error parsing card")
// 		http.Error(w, "Error encoding card data", http.StatusInternalServerError)
// 	}
// }
