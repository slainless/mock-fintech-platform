package platform

import "time"

type TransactionStatus uint32

type TransactionHistory interface {
	ID() string

	BalanceMutation() *MonetaryAmount
	Service() MonetaryService
	User() User

	Timestamp() *time.Time
	Address() string

	Status() TransactionStatus
	Note() string
}
