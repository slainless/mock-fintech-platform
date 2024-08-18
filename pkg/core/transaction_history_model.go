package core

import (
	"time"

	"github.com/slainless/mock-fintech-platform/pkg/internal/query"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

type TransactionHistory struct {
	model *query.History
}

func (h *TransactionHistory) AccountID() string {
	return h.model.AccountUUID
}

func (h *TransactionHistory) Address() *string {
	return h.model.Address
}

func (h *TransactionHistory) BalanceMutation() *platform.MonetaryAmount {
	return &platform.MonetaryAmount{
		Currency: platform.Currency(h.model.Currency),
		Value:    h.model.Mutation,
	}
}

func (h *TransactionHistory) DestAccountID() *string {
	return h.model.DestUUID
}

func (h *TransactionHistory) ID() string {
	return h.model.UUID
}

func (h *TransactionHistory) Note() *string {
	return h.model.TransactionNote
}

func (h *TransactionHistory) ServiceID() string {
	return h.model.ServiceID
}

func (h *TransactionHistory) Status() platform.TransactionStatus {
	return platform.TransactionStatus(h.model.Status)
}

func (h *TransactionHistory) Timestamp() *time.Time {
	return &h.model.TransactionDate
}

func (h *TransactionHistory) Type() platform.TransactionType {
	return platform.TransactionType(h.model.TransactionType)
}

func (h *TransactionHistory) UserID() string {
	return h.model.UserUUID
}

func HistoryFrom(model *query.History) platform.TransactionHistory {
	return &TransactionHistory{model: model}
}

func HistoriesFrom(models []query.History) []platform.TransactionHistory {
	histories := make([]platform.TransactionHistory, 0, len(models))
	for _, model := range models {
		histories = append(histories, HistoryFrom(&model))
	}

	return histories
}
