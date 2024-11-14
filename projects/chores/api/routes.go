package api

import (
	"chores_app/configs"
	"chores_app/data"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Dependencies struct {
	Repository data.Repository
	Config     *configs.APIConfigs
}

type Routes struct {
	*chi.Mux
	Dependencies
}

func NewRoutes(deps Dependencies) *Routes {
	routes := new(Routes)
	routes.Dependencies = deps

	chir := chi.NewRouter()

	// middlewares
	chir.Use(middleware.Logger)
	chir.Use(middleware.Recoverer)
	chir.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"Host",
			"X-Real-IP",
			"X-Forwarded-For",
			"X-Forwarded-Host",
			"X-Forwarded-Server",
			"X-Forwarded-Port",
			"X-Forwarded-Proto",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// register the routes here

	chir.Route("/chores", func(r chi.Router) {
		r.Get("/", routes.GetAllChores)
		r.Post("/", routes.CreateChore)
		r.Get("/{choreID}", routes.GetChoreByID)
		r.Patch("/{choreID}", routes.UpdateChoreDetails)
		r.Patch("/{choreID}/reschedule", routes.UpdateChoreSchedule)
		r.Patch("/{choreID}/deactivate", routes.DeactivateChore)
	})

	chir.Route("/members", func(r chi.Router) {
		r.Get("/", routes.GetAllMembers)
		r.Post("/", routes.CreateMember)
		r.Get("/{choreID}", routes.GetMemberByID)
		r.Patch("/{choreID}", routes.UpdateMemberInfo)
		r.Patch("/{choreID}/deactivate", routes.DeactivateMember)
	})

	routes.Mux = chir

	return routes
}
