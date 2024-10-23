package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB(dataSourceName string) {
	var err error
	fmt.Print("dataSourceName:")
	fmt.Print(dataSourceName)
	DB, err = sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\nConnected to DB:")
	// Create users table if it doesn't exist
	// schema := `CREATE TABLE IF NOT EXISTS users (
	//     id SERIAL PRIMARY KEY,
	//     username TEXT UNIQUE NOT NULL,
	//     password TEXT NOT NULL
	// );`

	// if _, err := DB.Exec(schema); err != nil {
	// 	log.Fatalln(err)
	//}
}
