package dbinterface

import _ "github.com/go-sql-driver/mysql"

// Create makes a table called example
// and populates it
func Create(db DB) error {
	// create the database
	if _, err := db.Exec(`CREATE TABLE example 
	(name VARCHAR(20), created DATETIME)`); err != nil {
		return err
	}
	if _, err := db.Exec(`INSERT INTO example (name, created) values
	("Aaron", NOW())`); err != nil {
		return err
	}
	return nil
}
