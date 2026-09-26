/*

The package works on 2 tables on a PostgreSQL data base server.

The names of the tables are:

	* Users
	* Userdata

The definitions of the tables in the PostgreSQL server are:

	CREATE TABLE Users (
		ID SERIAL,
		Username VARCHAR(100) PRIMARY KEY
	);

	CREATE TABLE Userdata (
		UserID Int NOT NULL,
		Name VARCHAR(100),
		Surname VARCHAR(100),
		Description VARCHAR(200)
	);

	This is rendered as code

This is not rendered as code
*/
package userPost

// The Userdata structure is for holding full user data
// from the Userdata table and the Username from the
// Users table
type UserData struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}
