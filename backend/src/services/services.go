package services

import (
	"log"
	"time"

	"github.com/ManuEduardo/random-topic/src/domain"
	"github.com/ManuEduardo/random-topic/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type IServices interface {
	GetUserById(id string) (domain.User, error)
	PostUser(user domain.User) error
	PostCard(user domain.Card) error
	GetRandomCard(id string) (domain.Card, error)
}

type Services struct {
	_repo repository.IRepository
}

func New(repo repository.IRepository) *Services {
	return &Services{
		_repo: repo,
	}
}

func (serv *Services) GetUserById(id string) (domain.User, error) {
	user, err := serv._repo.GetUserById(id)

	if err != nil {
		log.Println("Error getting user")
		return domain.User{}, err
	}

	reponseUser := domain.User{
		ID:        int64((user["id"]).(int)),
		Name:      user["name"].(string),
		BirthDate: user["birth_date"].(time.Time).Format("2006-01-02"),
		Gender:    user["gender"].(string),
	}

	return reponseUser, nil
}

func (serv *Services) PostUser(user domain.User) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Panicln(err.Error())
		return err
	}
	user.Password = string(passwordHash)
	err = serv._repo.PostUser(user)

	if err != nil {
		log.Println("Error creating user and inserting default cards:", err)
		return err
	}

	return nil
}

func (serv *Services) GetRandomCard(id string) (domain.Card, error) {
	cardData, err := serv._repo.GetRandomCard(id)

	if err != nil {
		log.Println("Error getting card")
		return domain.Card{}, err
	}

	responseCard := domain.Card{
		ID:        int64(cardData["id"].(int)),        // Mapeo del ID de la tarjeta
		Title:     cardData["title"].(string),         // Mapeo del título de la tarjeta
		Content:   cardData["content"].(string),       // Mapeo del contenido de la tarjeta
		IsDefault: cardData["is_default"].(bool),      // Mapeo de si es tarjeta predeterminada
		UserID:    int64(cardData["user_id"].(int64)), // Mapeo del ID de usuario

		// Mapeo de la relación con el tipo de tarjeta (CardType)
		Type: domain.CardType{
			ID:          cardData["type"].(map[string]interface{})["id"].(int64),
			Name:        cardData["type"].(map[string]interface{})["name"].(string),
			Description: cardData["type"].(map[string]interface{})["description"].(string),
		},
	}

	return responseCard, nil
}

func (serv *Services) PostCard(card domain.Card) error {
	err := serv._repo.PostCard(card)

	if err != nil {
		log.Println("Error creating card")
		return err
	}

	return nil
}
