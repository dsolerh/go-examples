package main

import (
	"chores_app/api"
	"chores_app/configs"
	"chores_app/data"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := configs.New()
	repo := data.NewRepository(data.RepoDependencies{Config: config.RepositoryConfigs})
	routes := api.NewRoutes(api.Dependencies{Repository: repo, Config: config.APIConfigs})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Port),
		Handler: routes,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
