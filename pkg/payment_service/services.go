package payment_service

import "github.com/slainless/mock-fintech-platform/pkg/platform"

func InitiatePaymentServices() map[string]platform.PaymentService {
	return map[string]platform.PaymentService{
		"bank_of_the_xyz": &MockPaymentService{},
		"infinite_loan":   &MockPaymentService{},
		"fishtech":        &MockPaymentService{},
	}
}
