package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Example hold the results of our queries
type Example struct {
	Name    string
	Created *time.Time
}

// Setup configures and returns our database
// connection poold
func Setup() (*sql.DB, error) {
	os.Setenv("MYSQLUSERNAME", "test")
	os.Setenv("MYSQLPASSWORD", "testdbpass")
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@/test?parseTime=true",
			os.Getenv("MYSQLUSERNAME"),
			os.Getenv("MYSQLPASSWORD"),
		),
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
