package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DataBaseConnector struct {
	db sql.DB
}
