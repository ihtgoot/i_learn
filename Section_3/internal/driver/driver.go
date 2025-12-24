package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// hold database connection
type DB struct {
	SQl *sql.DB
}

var dbconn = &DB{}

const maxOpenDBconn = 10              // maximum connection
const maxIdleDBconn = 10              // max ideal time for a connection
const maxDBLifeTime = 5 * time.Minute // max time for a connection

// establish database connection
func ConnectSQL(dsn string) (*DB, error) {
	db, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxOpenDBconn)
	db.SetConnMaxLifetime(maxIdleDBconn)
	db.SetConnMaxIdleTime(maxOpenDBconn)

	dbconn.SQl = db
	err = testDB(db)
	if err != nil {
		return nil, err
	}
	return dbconn, nil
}

// test databaseconnection
func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// creat new dataase
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
