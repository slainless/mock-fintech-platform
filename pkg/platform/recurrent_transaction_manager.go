package platform

import "time"

type RecurrentTransactionManager interface {
	BillingServices() map[string]RecurrentPaymentService

	IntervalRegister(
		monetaryService MonetaryService,
		billingService RecurrentPaymentService,
		user User,
		interval time.Duration,
	) (id int, err error)

	FixedPerYearRegister(
		monetaryService MonetaryService,
		billingService RecurrentPaymentService,
		user User,
		date time.Time,
	) (id int, err error)

	FixedPerMonthRegister(
		monetaryService MonetaryService,
		billingService RecurrentPaymentService,
		user User,
		date time.Time,
	) (id int, err error)

	Unregister(id int) error
}
