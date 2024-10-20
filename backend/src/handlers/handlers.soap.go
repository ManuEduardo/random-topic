package handlers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ManuEduardo/random-topic/src/domain"
)

// soapHandler maneja todas las solicitudes SOAP y las enruta a la operación adecuada.
func (handler *Handler) SoapHandler(w http.ResponseWriter, r *http.Request) {
	var envelope domain.SOAPEnvelope

	// Leer el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Printf("Request body: %s", string(body))

	// Decodificar el cuerpo de la solicitud
	if err := xml.Unmarshal(body, &envelope); err != nil {
		log.Printf("Error decoding SOAP request: %v", err)
		http.Error(w, "Invalid SOAP request", http.StatusBadRequest)
		return
	}

	fmt.Printf("Decoded envelope: %+v\n", envelope)

	// Enrutar a la operación correspondiente según la solicitud recibida
	if envelope.Body.CreateUserRequest != nil {
		handler.HandleCreateUser(w, r, *envelope.Body.CreateUserRequest)
	} else if envelope.Body.GetUserRequest != nil {
		handler.HandleGetUser(w, r, *envelope.Body.GetUserRequest)
	} else if envelope.Body.CreateCardRequest != nil {
		handler.HandlePostCard(w, r, *envelope.Body.CreateCardRequest)
	} else if envelope.Body.GetRandomCardRequest != nil {
		handler.HandleGetRandomCard(w, r, *envelope.Body.GetRandomCardRequest)
	} else {
		http.Error(w, "Operación no soportada", http.StatusNotImplemented)
	}
}

// HandleCreateUser maneja la creación de un nuevo usuario.
func (handler *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request, req domain.CreateUserRequest) {
	user := req.User
	err := handler._services.PostUser(user)
	success := err == nil

	response := domain.SOAPEnvelope{
		SOAPNS: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: domain.SOAPBody{
			CreateUserResponse: &domain.CreateUserResponse{
				Success: success,
			},
		},
	}

	w.Header().Set("Content-Type", "text/xml")
	xml.NewEncoder(w).Encode(response)
}

// HandleGetUser maneja la obtención de un usuario por su ID.
func (handler *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request, req domain.GetUserRequest) {
	userID := req.UserID
	user, err := handler._services.GetUserById(fmt.Sprintf("%d", userID))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := domain.SOAPEnvelope{
		SOAPNS: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: domain.SOAPBody{
			GetUserResponse: &domain.GetUserResponse{
				User: user,
			},
		},
	}

	w.Header().Set("Content-Type", "text/xml")
	xml.NewEncoder(w).Encode(response)
}

// HandlePostCard maneja la creación de una nueva tarjeta.
func (handler *Handler) HandlePostCard(w http.ResponseWriter, r *http.Request, req domain.CreateCardRequest) {
	card := req.Card
	err := handler._services.PostCard(card)
	if err != nil {
		http.Error(w, "Error creating card", http.StatusInternalServerError)
		return
	}
	success := true

	response := domain.SOAPEnvelope{
		SOAPNS: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: domain.SOAPBody{
			CreateCardResponse: &domain.CreateCardResponse{
				Success: success,
			},
		},
	}

	w.Header().Set("Content-Type", "text/xml")
	xml.NewEncoder(w).Encode(response)
}

// HandleGetRandomCard maneja la obtención de una tarjeta aleatoria para un usuario específico.
func (handler *Handler) HandleGetRandomCard(w http.ResponseWriter, r *http.Request, req domain.GetRandomCardRequest) {
	userID := req.UserID
	card, err := handler._services.GetRandomCard(fmt.Sprintf("%d", userID))
	if err != nil {
		http.Error(w, "Random card not found", http.StatusNotFound)
		return
	}

	response := domain.SOAPEnvelope{
		SOAPNS: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: domain.SOAPBody{
			GetRandomCardResponse: &domain.GetRandomCardResponse{
				Card: card,
			},
		},
	}

	w.Header().Set("Content-Type", "text/xml")
	xml.NewEncoder(w).Encode(response)
}
