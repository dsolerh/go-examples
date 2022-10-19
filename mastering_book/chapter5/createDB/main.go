package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	pass     = "pass"
	database = "master"
)

func main() {

	// connection string
	conn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		pass,
		database,
	)

	// open PostgreSQL database
	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("Open():", err)
		return
	}
	defer db.Close()

	statements := []string{
		"DROP DATABASE IF EXISTS go;",
		"CREATE DATABASE go;",
		"DROP TABLE IF EXISTS Users;",
		"DROP TABLE IF EXISTS Userdata;",
		`CREATE TABLE Users (
			ID SERIAL,
			Username VARCHAR(100) PRIMARY KEY
		);`,
		`CREATE TABLE Userdata (
			UserID Int NOT NULL,
			Name VARCHAR(100),
			Surname VARCHAR(100),
			Description VARCHAR(200)
		);`,
	}
	for _, statement := range statements {
		_, err := db.Exec(statement)
		if err != nil {
			fmt.Println(err)
		}
	}

}
