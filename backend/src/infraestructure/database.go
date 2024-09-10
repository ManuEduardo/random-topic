package infraestructure

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DbInstance interface {
	InitDB() error
	CloseDB()
	GetDBPool() *pgxpool.Pool
}

type dbPostgres struct {
	dbUrl  string
	dbPool *pgxpool.Pool
}

func New(dbURL string) *dbPostgres {
	log.Println("INSTANCIA DE LA BASE DE DATOS")
	return &dbPostgres{dbUrl: dbURL}
}

func (db *dbPostgres) InitDB() error {

	var err error
	db.dbPool, err = pgxpool.New(context.Background(), db.dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return err
	}
	return nil
}

func (db dbPostgres) CloseDB() {
	db.dbPool.Close()
}

func (db dbPostgres) GetDBPool() *pgxpool.Pool {
	return db.dbPool
}
