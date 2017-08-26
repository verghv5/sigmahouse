package sigmahouse

import (
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB	   *sql.DB
}

func (a *App) Initialize(user, password, dbname string) { }

func (a *App) Run(addr string) { }