package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/hudarashid/goCrud/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)
		r.Get("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.HandlerGetWorkoutByID))
		r.Post("/workouts", app.Middleware.RequireUser(app.WorkoutHandler.HandlerCreateWorkout))
		r.Put("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.HandleUpdateWorkoutByID))
		r.Delete("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.DeleteWorkoutByID))

	})

	// Function in Go is first class citizen, i.e can pass function as variable 'HealthCheck'
	r.Get("/health", app.HealthCheck)
	r.Post("/users", app.UserHandler.HandleRegisterUser)
	r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)
	return r
}
