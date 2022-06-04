package userPost

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

/*
This block of global variables holds the connection details to the Postgres server
	Hostname: is theIP or the hostname of the server
	Port: is the TCP port the DB server listens to
	Username: is the username of the database user
	Password: is the password of the database user
	Database: is the name of the Database in PostgreSQL
*/
var (
	Hostname = ""
	Port     = 2345
	Username = ""
	Password = ""
	Database = ""
)

// openConnection() is for opening the Postgres connection
// in order to be used by the other functions of the package.
func openConnection() (*sql.DB, error) {
	// connection string
	conn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Hostname,
		Port,
		Username,
		Password,
		Database,
	)

	// open connection
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
