package platform

type MoneyExchangeService interface {
	// service ID
	ID() string

	// currency conversion method
	Convert(amount MonetaryAmount, to Currency) (MonetaryAmount, error)
}
