package platform

import "time"

type TransactionStatus uint32
type TransactionType uint32

type TransactionHistory interface {
	ID() string
	ServiceID() string
	UserID() string

	BalanceMutation() *MonetaryAmount

	Timestamp() *time.Time
	Address() string

	Status() TransactionStatus
	Type() TransactionType

	Note() string
}
