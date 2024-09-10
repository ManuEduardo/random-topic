package infraestructure

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func InitDB(dbURL string) error {

	var err error
	dbPool, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return err
	}
	return nil
}

func CloseDB() {
	dbPool.Close()
}

func GetDBPool() *pgxpool.Pool {
	return dbPool
}

// // InsertData inserts a new record into the database
// func InsertData(name string, age int) error {
// 	sql := "INSERT INTO users (name, age) VALUES ($1, $2)"
// 	_, err := dbPool.Exec(context.Background(), sql, name, age)
// 	if err != nil {
// 		return fmt.Errorf("failed to insert data: %v", err)
// 	}
// 	return nil
// }

// // GetAllData retrieves all records from the database
// func GetAllData() ([]map[string]interface{}, error) {
// 	rows, err := dbPool.Query(context.Background(), "SELECT id, name, age FROM users")
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to query data: %v", err)
// 	}
// 	defer rows.Close()

// 	var result []map[string]interface{}

// 	for rows.Next() {
// 		var id int
// 		var name string
// 		var age int

// 		err = rows.Scan(&id, &name, &age)
// 		if err != nil {
// 			return nil, err
// 		}

// 		result = append(result, map[string]interface{}{
// 			"id":   id,
// 			"name": name,
// 			"age":  age,
// 		})
// 	}
// 	return result, nil
// }

// // GetDataByID retrieves a single record by ID
// func GetDataByID(id int) (map[string]interface{}, error) {
// 	row := dbPool.QueryRow(context.Background(), "SELECT id, name, age FROM users WHERE id = $1", id)

// 	var userID int
// 	var name string
// 	var age int

// 	err := row.Scan(&userID, &name, &age)
// 	if err != nil {
// 		return nil, fmt.Errorf("no record found: %v", err)
// 	}

// 	return map[string]interface{}{
// 		"id":   userID,
// 		"name": name,
// 		"age":  age,
// 	}, nil
// }
