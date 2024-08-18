package core

import (
	"database/sql"
)

type MonetaryAccount struct {
	db *sql.DB
}
