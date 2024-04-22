package pkg

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Application struct {
	Pql   *pgxpool.Pool
}

func App() (Application, error) {
	app := &Application{}
	conn, err := NewPgxConn()
	if err != nil {
		return Application{}, err
	}
	app.Pql = conn

	return *app, nil
}

func (app *Application) CloseDBConnection() {
	app.CloseDBConnection()
}
