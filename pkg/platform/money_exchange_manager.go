package platform

type MoneyExchangeManager interface {
	ExchangeServices() map[string]MoneyExchangeService

	Convert(service MoneyExchangeService, amount MonetaryAmount, to Currency) (MonetaryAmount, error)
}
