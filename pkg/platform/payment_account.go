package platform

type PaymentAccount interface {
	ID() string
	ServiceID() string

	Currency() Currency

	// lets just assume balance can be negative
	// since the value came from payment service
	Balance() MonetaryAmount
}
