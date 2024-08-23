package payment_service

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/internal/util"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/model"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type MockPaymentService struct{}

// Balance implements platform.PaymentService.
func (s *MockPaymentService) Balance(ctx context.Context, account *platform.PaymentAccount) (*platform.MonetaryAmount, error) {
	util.MockSleep(2 * time.Second)
	return &platform.MonetaryAmount{
		Value:    rand.Int63(),
		Currency: "USD",
	}, nil
}

// GetMatchingHistory implements platform.PaymentService.
func (*MockPaymentService) GetMatchingHistory(ctx context.Context, account *platform.PaymentAccount, history *platform.TransactionHistory) (*platform.TransactionHistory, error) {
	util.MockSleep(2 * time.Second)
	// let us always receive the matching history...
	if util.LeaveItToRNG() && false {
		return nil, errors.New("Oops! Failed to get matching history")
	} else {
		note := "Received from " + account.UserUUID.String()
		match := *history
		match.TransactionNote = &note
		match.AccountUUID = *history.DestUUID
		match.DestUUID = nil
		match.Mutation = history.Mutation * -1
		match.Address = nil
		return &match, nil
	}
}

// Send implements platform.PaymentService.
func (*MockPaymentService) Send(ctx context.Context, user *platform.User, source *platform.PaymentAccount, des *platform.PaymentAccount, amount int64) (*platform.TransactionHistory, error) {
	util.MockSleep(3 * time.Second)
	if util.LeaveItToRNG() {
		return nil, errors.New("Failed to send money")
	} else {
		return &platform.TransactionHistory{
			TransactionHistories: model.TransactionHistories{
				UUID:            uuid.New(),
				AccountUUID:     source.UUID,
				DestUUID:        &des.UUID,
				Mutation:        amount * -1,
				Currency:        "USD",
				TransactionDate: time.Now(),
				IssuerUUID:      &user.UUID,
			},
			ServiceID: source.ServiceID,
			UserUUID:  source.UserUUID,
		}, nil
	}
}

// Validate implements platform.PaymentService.
func (*MockPaymentService) Validate(ctx context.Context, user *platform.User, accountForeignID string, callbackData string) error {
	util.MockSleep(2 * time.Second)
	if util.LeaveItToRNG() {
		return errors.New("Failed to validate user")
	} else {
		return nil
	}
}

// Withdraw implements platform.PaymentService.
func (*MockPaymentService) Withdraw(ctx context.Context, user *platform.User, account *platform.PaymentAccount, amount int64, callbackData string) (*platform.TransactionHistory, error) {
	util.MockSleep(2 * time.Second)
	if util.LeaveItToRNG() {
		return nil, errors.New("Failed to withdraw money")
	} else {
		return &platform.TransactionHistory{
			TransactionHistories: model.TransactionHistories{
				UUID:            uuid.New(),
				AccountUUID:     account.UUID,
				Mutation:        amount * -1,
				Currency:        "USD",
				TransactionDate: time.Now(),
				IssuerUUID:      &user.UUID,
			},
			ServiceID: account.ServiceID,
			UserUUID:  account.UserUUID,
		}, nil
	}
}

func NewMockPaymentService() platform.PaymentService {
	return &MockPaymentService{}
}
