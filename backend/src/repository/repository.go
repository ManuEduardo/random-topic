package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ManuEduardo/random-topic/src/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type IRepository interface {
	GetUserById(id string) (map[string]interface{}, error)
	PostUser(user domain.User) error
}

type PostgresRepository struct {
	_dbPool *pgxpool.Pool
}

func New(dbPool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		_dbPool: dbPool,
	}
}

func (repo *PostgresRepository) GetUserById(id string) (map[string]interface{}, error) {
	row := repo._dbPool.QueryRow(context.Background(), "SELECT id, name, birth_date, gender FROM users WHERE id = $1", id)
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
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Panicln(err.Error())
	}
	sql := "INSERT INTO users (name, password, birth_date, gender ) VALUES ($1, $2, $3, $4)"
	_, err = repo._dbPool.Exec(context.Background(), sql, user.Name, passwordHash, user.BirthDate, user.Gender)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
