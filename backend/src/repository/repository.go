package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IRepository interface {
	GetUserById(id string) (map[string]interface{}, error)
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
