package platform

import "time"

type RecurrentTransactionManager interface {
	BillingServices() []RecurrentBillingService

	IntervalRegister(
		monetaryService MonetaryService,
		billingService RecurrentBillingService,
		user User,
		interval time.Duration,
	) (id int, err error)

	FixedPerYearRegister(
		monetaryService MonetaryService,
		billingService RecurrentBillingService,
		user User,
		date time.Time,
	) (id int, err error)

	FixedPerMonthRegister(
		monetaryService MonetaryService,
		billingService RecurrentBillingService,
		user User,
		date time.Time,
	) (id int, err error)

	Unregister(id int) error
}
