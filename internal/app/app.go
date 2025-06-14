package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hudarashid/goCrud/internal/api"
	"github.com/hudarashid/goCrud/internal/store"
	"github.com/hudarashid/goCrud/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

// Constructor - returning the pointer '*' type of Application
func NewApplication() (*Application, error) {
	pgDB, err := store.Open()

	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}
	// use logger as oppose to use the fmt
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// store will go here
	workoutStore := store.NewPostgresWorkoutStore(pgDB)

	// handlers
	workoutHandler := api.NewWorkoutHandler(workoutStore, logger)

	// '&' is to point to the memory of the Application
	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}

	return app, nil

}

// r is a pointer because it goint to contain request data from Client - so need to persist the data what Client send
// w is coming from our own http server -  so its going to be in the scope of this function and don't need any modification
func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available\n")
}
