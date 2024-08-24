package payment_service

import (
	"context"
	"errors"
	"fmt"
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
func (*MockPaymentService) Send(ctx context.Context, user *platform.User, source *platform.PaymentAccount, des *platform.PaymentAccount, amount int64, callbackData string) (*platform.TransactionHistory, error) {
	util.MockSleep(3 * time.Second)
	if util.LeaveItToRNG() {
		return nil, errors.New("Failed to send money")
	} else {
		var extraNote string
		if user.UUID != source.UserUUID {
			extraNote = fmt.Sprintf("This is issued by shared user")
		} else {
			extraNote = fmt.Sprintf("This is issued by owner")
		}

		note := fmt.Sprintf(
			"User [%s] sending from account [%s, service: %s] with amount [%d] to account [%s, service: %s]. (note: %s)",
			user.UUID, source.UUID, source.ServiceID, amount, des.UUID, des.ServiceID, extraNote,
		)

		return &platform.TransactionHistory{
			TransactionHistories: model.TransactionHistories{
				UUID:            uuid.New(),
				AccountUUID:     source.UUID,
				DestUUID:        &des.UUID,
				Mutation:        amount * -1,
				Currency:        "USD",
				TransactionDate: time.Now(),
				TransactionNote: &note,
				TransactionType: int16(platform.TransactionTypeSend),
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
		var extraNote string
		if user.UUID != account.UserUUID {
			extraNote = fmt.Sprintf("This is issued by shared user")
		} else {
			extraNote = fmt.Sprintf("This is issued by owner")
		}

		note := fmt.Sprintf(
			"User [%s] withdrawing from account [%s, service: %s] with amount [%d]. Callback data: %s. (note: %s)",
			user.UUID, account.UUID, account.ServiceID, amount, callbackData, extraNote,
		)
		return &platform.TransactionHistory{
			TransactionHistories: model.TransactionHistories{
				UUID:            uuid.New(),
				AccountUUID:     account.UUID,
				Mutation:        amount * -1,
				Currency:        "USD",
				TransactionDate: time.Now(),
				TransactionType: int16(platform.TransactionTypeWithdraw),
				IssuerUUID:      &user.UUID,
				TransactionNote: &note,
			},
			ServiceID: account.ServiceID,
			UserUUID:  account.UserUUID,
		}, nil
	}
}

func NewMockPaymentService() platform.PaymentService {
	return &MockPaymentService{}
}
