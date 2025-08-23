package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	db *sql.DB
}

func New() *MySQLDB {

	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{db}
}
