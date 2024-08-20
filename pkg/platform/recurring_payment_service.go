package platform

import "context"

// type RecurringPaymentType uint8
type RecurringPaymentChargingMethod uint8

// const (
// 	RecurringPaymentYearly RecurringPaymentType = iota
// 	RecurringPaymentMonthly
// 	RecurringPaymentFixedInterval
// )

const (
	RecurringPaymentChargeUpfront RecurringPaymentChargingMethod = iota
	RecurringPaymentChargeNonUpfront
)

type RecurringPaymentService interface {
	// should return history when subscribing to upfront recurring payments
	Subscribe(ctx context.Context, account *PaymentAccount, billingID, callbackData string) (*RecurringPayment, *TransactionHistory, error)
	// should return history when subscribing to non-upfront recurring payments
	Unsubscribe(ctx context.Context, payment *RecurringPayment) (*TransactionHistory, error)

	Bill(ctx context.Context, payment *RecurringPayment) (*TransactionHistory, error)
}
