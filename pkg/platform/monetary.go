package platform

type Currency string

type MonetaryAmount struct {
	Currency Currency
	Value    int64
}
