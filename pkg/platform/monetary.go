package platform

type Currency string

type MonetaryAmount struct {
	currency Currency
	value    int64
}
