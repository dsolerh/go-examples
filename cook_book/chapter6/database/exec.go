package database

import "database/sql"

// Exec replaces the Exec from the previous recipe
func Exec(db *sql.DB) error {
	// uncaught error on cleanup, but we always
	// want to cleanup
	defer db.Exec("DROP TABLE example")
	if err := Create(db); err != nil {
		return err
	}
	if err := Query(db, "Aaron"); err != nil {
		return err
	}
	return nil
}
