package platform

type PaymentAccount interface {
	ID() string

	Currency() Currency

	// lets just assume balance can be negative
	// since the value came from payment service
	TrackedBalance() MonetaryAmount
}
