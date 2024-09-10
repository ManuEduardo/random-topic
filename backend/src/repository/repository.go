package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ManuEduardo/random-topic/src/domain"
	"github.com/ManuEduardo/random-topic/src/infraestructure"
)

type IRepository interface {
	GetUserById(id string) (map[string]interface{}, error)
	PostUser(user domain.User) error
	PostCard(user domain.Card) error
	GetRandomCard(id string) (map[string]interface{}, error)
}

type PostgresRepository struct {
	_database infraestructure.DbInstance
}

func New(db infraestructure.DbInstance) *PostgresRepository {
	return &PostgresRepository{
		_database: db,
	}
}

func (repo *PostgresRepository) GetUserById(id string) (map[string]interface{}, error) {
	row := repo._database.GetDBPool().QueryRow(context.Background(), "SELECT id, name, birth_date, gender FROM users WHERE id = $1", id)
	var userID int
	var name string
	var birth_date time.Time
	var gender string

	err := row.Scan(&userID, &name, &birth_date, &gender) //aqui pasa el error
	if err != nil {
		log.Printf("Error getting record %v", err.Error())
		return nil, fmt.Errorf("no record found: %v", err)
	}

	return map[string]interface{}{
		"id":         userID,
		"name":       name,
		"birth_date": birth_date,
		"gender":     gender,
	}, nil
}

func (repo *PostgresRepository) PostUser(user domain.User) error {
	sql := "INSERT INTO users (name, password, birth_date, gender ) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo._database.GetDBPool().QueryRow(context.Background(), sql, user.Name, user.Password, user.BirthDate, user.Gender).Scan(&user.ID)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	sqlCards := `
        INSERT INTO cards (title, content, is_default, user_id, type_id) VALUES 
        ('¿Cuál es tu libro favorito?', 'Comparte el libro que más te ha impactado y por qué.', true, $1, 1),
        ('Tema de conversación: Viajes', 'Hablemos sobre los lugares que has visitado y los que te gustaría conocer.', true, $1, 2),
        ('Dinámica: Dos verdades y una mentira', 'Cuenta dos cosas verdaderas y una falsa sobre ti, y deja que los demás adivinen.', true, $1, 3),
        ('¿Qué superpoder te gustaría tener?', 'Imagina que puedes elegir un superpoder. ¿Cuál sería y por qué?', true, $1, 1),
        ('Tema de conversación: Tecnología', 'Discutamos sobre los avances tecnológicos más recientes y su impacto en la sociedad.', true, $1, 2);
    `
	_, err = repo._database.GetDBPool().Exec(context.Background(), sqlCards, user.ID)
	if err != nil {
		log.Println("Error inserting default cards:", err)
		return err
	}

	return nil
}

func (repo *PostgresRepository) GetRandomCard(id string) (map[string]interface{}, error) {
	row := repo._database.GetDBPool().QueryRow(context.Background(), "SELECT cards.id, cards.title, cards.content, cards.is_default, cards.user_id, cards.type_id, card_types.name AS type_name, card_types.description AS type_description FROM cards JOIN card_types ON cards.type_id = card_types.id WHERE cards.user_id = $1 ORDER BY random() LIMIT 1", id)
	var cardID int
	var cardTitle string
	var cardContent string
	var isDefault bool
	var cardUserID int64
	var cardTypeID int64
	var typeName string
	var typeDescription string

	err := row.Scan(&cardID, &cardTitle, &cardContent, &isDefault, &cardUserID, &cardTypeID, &typeName, &typeDescription)
	if err != nil {
		log.Printf("Error getting record %v", err.Error())
		return nil, fmt.Errorf("no record found: %v", err)
	}

	result := map[string]interface{}{
		"id":         cardID,
		"title":      cardTitle,
		"content":    cardContent,
		"is_default": isDefault,
		"user_id":    cardUserID,
		"type": map[string]interface{}{
			"id":          cardTypeID,
			"name":        typeName,
			"description": typeDescription,
		},
	}

	return result, nil
}

func (repo *PostgresRepository) PostCard(card domain.Card) error {
	sql := "INSERT INTO cards (title, content, is_default, user_id, type_id) VALUES ($1, $2, 'false', $3, $4)"
	_, err := repo._database.GetDBPool().Exec(context.Background(), sql, card.Title, card.Content, card.UserID, card.TypeID)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
