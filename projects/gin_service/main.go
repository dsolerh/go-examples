package main

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func main() {
	router := gin.Default()
	router.GET("/:name", IndexHandler)
	router.GET("/xml/:firstName/:lastName", PersonXMLHandler)
	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func IndexHandler(ctx *gin.Context) {
	name := ctx.Params.ByName("name")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello " + name,
	})
}

func PersonXMLHandler(ctx *gin.Context) {
	firstName := ctx.Params.ByName("firstName")
	lastName := ctx.Params.ByName("lastName")
	ctx.XML(http.StatusOK, Person{FirstName: firstName, LastName: lastName})
}
