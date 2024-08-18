package platform

type MonetaryAccount interface {
	ID() string

	Currency() Currency

	// lets just assume balance can be negative
	// since the value came from payment service
	TrackedBalance() MonetaryAmount
}
