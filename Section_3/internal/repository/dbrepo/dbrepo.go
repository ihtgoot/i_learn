package dbrepo

import (
	"database/sql"

	"github.com/ihtgoot/i_learn/Section_3/internal/config"
	"github.com/ihtgoot/i_learn/Section_3/internal/repository"
)

type postgresDBrepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBrepo{
		App: a,
		DB:  conn,
	}
}
