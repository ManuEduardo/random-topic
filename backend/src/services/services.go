package services

import (
	"log"
	"time"

	"github.com/ManuEduardo/random-topic/src/domain"
	"github.com/ManuEduardo/random-topic/src/repository"
)

type IServices interface {
	GetUserById(id string) (domain.User, error)
	PostUser(user domain.User) error
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
	err := serv._repo.PostUser(user)

	if err != nil {
		log.Println("Error creating user")
		return err
	}

	return nil
}
