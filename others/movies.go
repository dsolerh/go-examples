package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func toJSON() []byte {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON Marshaling failed: %s", err)
	}
	return data
}

func toJSONIdent() {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON Marshaling failed: %s", err)
	}
	fmt.Printf("data: %s\n", data)
}

func fromJSON() {
	var titles []struct{ Title string }
	data := toJSON()
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("titles: %v\n", titles)
}
