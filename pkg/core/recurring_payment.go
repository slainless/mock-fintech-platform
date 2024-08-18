package core

import "database/sql"

type RecurringPaymentManager struct {
	db *sql.DB
}

func NewRecurringPaymentManager(db *sql.DB) *RecurringPaymentManager {
	return &RecurringPaymentManager{
		db: db,
	}
}
