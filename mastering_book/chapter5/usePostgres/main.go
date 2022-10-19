package main

import (
	"fmt"
	"math/rand"
	"time"

	randomutil "github.com/dsolerh/examples/mastering.book/chapter2/random/utils"
	"github.com/dsolerh/examples/mastering.book/chapter5/userPost"
)

var (
	MIN = 0
	MAX = 26
)

func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64
	for i = 1; i <= length; i++ {
		myRand := randomutil.Random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp += newChar
	}
	return temp
}

func main() {
	userPost.Hostname = "localhost"
	userPost.Port = 5432
	userPost.Username = "root"
	userPost.Password = "pass"
	userPost.Database = "master"

	data, err := userPost.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range data {
		fmt.Println(v)
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)
	random_username := getString(5)

	t := userPost.UserData{
		Username:    random_username,
		Name:        "Daniel",
		Surname:     "Soler",
		Description: "This is me!",
	}
	id := userPost.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	err = userPost.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	// Trying to delete it again!
	err = userPost.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	id = userPost.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	t = userPost.UserData{
		Username:    random_username,
		Name:        "Daniel",
		Surname:     "Soler",
		Description: "This might not be me!",
	}

	err = userPost.UpdateUser(t)
	if err != nil {
		fmt.Println(err)
	}
}
