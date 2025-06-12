package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/hudarashid/goCrud/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	// Function in Go is first class citizen, i.e can pass function as variable 'HealthCheck'
	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandlerGetWorkoutByID)
	r.Post("/workouts", app.WorkoutHandler.HandlerCreateWorkout)
	r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkoutByID)
	return r
}
