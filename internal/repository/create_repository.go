package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sql.DB
}

func CreateRepository() *Repository {
	userName := os.Getenv("USERNAME_MEDODS")
	password := os.Getenv("PASSWORD_MEDODS")
	dbname := os.Getenv("DB_MEDODS")
	host := os.Getenv("HOST_MEDODS")
	port := os.Getenv("PORT_MEDODS")

	str := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", userName, password, dbname, host, port)
	fmt.Println(str)
	db, err := sql.Open("postgres", str)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &Repository{
		DB: db,
	}
}
