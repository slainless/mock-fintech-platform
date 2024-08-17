package platform

type MoneyExchangeManager interface {
	ExchangeServices() map[string]MoneyExchangeService

	// specific service currency conversion
	ConvertWith(service MoneyExchangeService, amount MonetaryAmount, to Currency) (MonetaryAmount, error)
	// manager managed currency conversion
	Convert(amount MonetaryAmount, to Currency) (MonetaryAmount, error)
}
