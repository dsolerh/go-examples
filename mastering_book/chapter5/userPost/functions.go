package userPost

import (
	"errors"
	"fmt"
	"strings"
)

// exists returns the User ID of the provided username
// if doesn't exist returns -1
func exists(username string) (userID int) {
	userID = -1
	username = strings.ToLower(username)

	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return userID
	}
	defer db.Close()

	statement := fmt.Sprintf(`select "id" from "users" where username = %s`, username)
	rows, err := db.Query(statement)
	if err != nil {
		fmt.Println(err)
		return userID
	}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			return userID
		}
		userID = id
	}
	defer rows.Close()
	return userID
}

// AddUser adds a new user to the database
// Returns new User ID
// or -1 if there was an error
func AddUser(d UserData) int {
	d.Username = strings.ToLower(d.Username)

	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	userID := exists(d.Username)
	if userID != -1 {
		fmt.Println("User already exists:", d.Username)
		return -1
	}

	insertStatement := `insert into "users" ("username") values ($1)`
	_, err = db.Exec(insertStatement, d.Username)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	userID = exists(d.Username)
	if userID == -1 {
		return userID
	}

	insertStatement = `insert into "userdata" ("userid", "name", "surname", "description") values ($1, $2, $3, $4)`

	_, err = db.Exec(insertStatement, userID, d.Name, d.Surname, d.Description)
	if err != nil {
		fmt.Println("Insert userdata:", err)
		return -1
	}

	return userID
}

// DeleteUser deletes an user from the database
func DeleteUser(id int) error {
	db, err := openConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Does the ID exists?
	statement := fmt.Sprintf(`select "username" from "users" where id = %d`, id)
	rows, err := db.Query(statement)
	if err != nil {
		return err
	}

	var username string
	for rows.Next() {
		err = rows.Scan(&username)
		if err != nil {
			return err
		}
	}
	defer db.Close()

	if exists(username) != id {
		return fmt.Errorf("user with ID %d does not exist", id)
	}

	// delete from userdata
	deleteStatement := `delete from "userdata" where userid=$1`
	_, err = db.Exec(deleteStatement, id)
	if err != nil {
		return err
	}

	// delete from users
	deleteStatement = `delete from "users" where id=$1`
	_, err = db.Exec(deleteStatement, id)
	if err != nil {
		return err
	}

	return nil
}

// ListUsers list all the users in the database
func ListUsers() ([]UserData, error) {
	db, err := openConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`select "id","username","name","surname","description" 
	from "users","userdata" where users.id = userdata.userid`)
	if err != nil {
		return nil, err
	}
	data := make([]UserData, 10)
	for rows.Next() {
		temp := UserData{}
		err = rows.Scan(&temp.ID, &temp.Username, &temp.Name, &temp.Surname, &temp.Description)
		if err != nil {
			return nil, err
		}
		data = append(data, temp)
	}
	defer rows.Close()

	return data, nil
}

// UpdateUser updates a user by it's ID
func UpdateUser(d UserData) error {
	db, err := openConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	userID := exists(d.Username)
	if userID == -1 {
		return errors.New("user does not exist")
	}

	d.ID = userID
	updateStatement := `update "userdata" set "name"=$1, "surname"=$2, "description"=$3 where "userid"=$4`
	_, err = db.Exec(updateStatement, d.Name, d.Surname, d.Description, d.ID)
	if err != nil {
		return err
	}

	return nil
}
