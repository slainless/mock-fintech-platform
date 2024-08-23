package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-co-op/gocron/v2"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/internal/query"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

var (
	ErrRecurringPaymentServiceNotSupported = errors.New("recurring payment service not supported")
	ErrRecurringPaymentServiceInvalidType  = errors.New("recurring payment service invalid type")
	ErrRecurringPaymentNotFound            = errors.New("recurring payment not found or no permission")
)

type RecurringPaymentManager struct {
	services map[string]platform.RecurringPaymentService
	db       *sql.DB

	errorTracker platform.ErrorTracker

	historyManager *TransactionHistoryManager
	scheduler      gocron.Scheduler
}

func NewRecurringPaymentManager(
	db *sql.DB,
	services map[string]platform.RecurringPaymentService,
	historyManager *TransactionHistoryManager,
	tracker platform.ErrorTracker,
) *RecurringPaymentManager {
	return &RecurringPaymentManager{
		db: db,

		errorTracker: tracker,

		historyManager: historyManager,
	}
}

func (m *RecurringPaymentManager) InitScheduler() (err error) {
	m.scheduler, err = gocron.NewScheduler()
	if err != nil {
		return err
	}

	return nil
}

func (m *RecurringPaymentManager) Subscribe(ctx context.Context, account *platform.PaymentAccount, serviceID, billingID, callbackData string) (*platform.RecurringPayment, *platform.TransactionHistory, error) {
	service := m.services[serviceID]
	if service == nil {
		return nil, nil, ErrRecurringPaymentServiceNotSupported
	}

	payment, history, err := service.Subscribe(ctx, account, billingID, callbackData)
	if history == nil {
		m.errorTracker.Report(ctx, m.historyManager.Records(ctx, []platform.TransactionHistory{*history}))
	}
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, nil, err
	}

	if payment.ServiceID != serviceID {
		err := errors.Join(
			fmt.Errorf("Invalid payment service received! %+v", payment),
			ErrRecurringPaymentServiceInvalidType,
		)

		m.errorTracker.Report(ctx, err)
		return nil, nil, err
	}

	_, err = GetRecurringPaymentType(payment)
	if err != nil {
		err := errors.Join(
			fmt.Errorf("failed to parse recurring payment type for payment: %+v", payment),
			ErrRecurringPaymentServiceInvalidType,
		)

		m.errorTracker.Report(ctx, err)
		return nil, nil, err
	}

	err = query.InsertRecurringPayment(ctx, m.db, payment)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, nil, err
	}

	err = m.Schedule(ctx, service, payment)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, nil, err
	}

	return payment, history, nil
}

func (m *RecurringPaymentManager) Unsubscribe(ctx context.Context, payment *platform.RecurringPayment) (*platform.TransactionHistory, error) {
	service := m.services[payment.ServiceID]
	if service == nil {
		m.errorTracker.Report(ctx, errors.Join(
			fmt.Errorf("Trying to unsubscribe from unsupported service: %+v", payment),
			ErrRecurringPaymentServiceNotSupported,
		))
		return nil, ErrRecurringPaymentServiceNotSupported
	}

	history, err := service.Unsubscribe(ctx, payment)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	err = query.DeleteRecurringPayment(ctx, m.db, payment.UUID)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		// let the process pass...
		// we still need to stop the job since we have successfully unsubscribed from the service
		// return nil, err
	}

	err = m.scheduler.RemoveJob(payment.UUID)
	if err != nil {
		if err != gocron.ErrJobNotFound {
			m.errorTracker.Report(ctx, err)
			// let the process pass...
			// return nil, err
		}
	}

	return history, nil
}

func (m *RecurringPaymentManager) Bill(ctx context.Context, account *platform.RecurringPayment) (*platform.TransactionHistory, error) {
	service := m.services[account.ServiceID]
	if service == nil {
		return nil, ErrRecurringPaymentServiceNotSupported
	}

	history, err := service.Bill(ctx, account)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	m.historyManager.Records(ctx, []platform.TransactionHistory{*history})
	return history, nil
}

func (m *RecurringPaymentManager) Schedule(ctx context.Context, service platform.RecurringPaymentService, payment *platform.RecurringPayment) error {
	job, err := m.scheduler.NewJob(
		gocron.MonthlyJob(1, gocron.NewDaysOfTheMonth(1), gocron.NewAtTimes(gocron.NewAtTime(0, 0, 0))),
		gocron.NewTask(func(m *RecurringPaymentManager, id uuid.UUID) {
			payment, err := m.GetPayment(context.TODO(), id)
			if err != nil {
				if err == ErrRecurringPaymentNotFound {
					m.scheduler.RemoveJob(id)
					return
				}
				return
			}

			m.Bill(context.TODO(), payment)
		}, m, payment.UUID),
		gocron.WithIdentifier(payment.UUID),
	)

	if err != nil {
		return err
	}

	_ = job
	return nil
}

func (m *RecurringPaymentManager) GetPayment(ctx context.Context, uuid uuid.UUID) (*platform.RecurringPayment, error) {
	payment, err := query.GetRecurringPayment(ctx, m.db, uuid)
	if err != nil {
		if err == qrm.ErrNoRows {
			return nil, ErrRecurringPaymentNotFound
		}

		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return payment, nil
}

func (m *RecurringPaymentManager) GetPaymentWithAccess(ctx context.Context, user *platform.User, uuid uuid.UUID) (*platform.RecurringPayment, error) {
	payment, err := query.GetRecurringPaymentWithAccess(ctx, m.db, user.UUID, uuid, AccountPermissionSubscription)
	if err != nil {
		if err == qrm.ErrNoRows {
			return nil, ErrRecurringPaymentNotFound
		}

		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return payment, nil
}

func (m *RecurringPaymentManager) GetPayments(ctx context.Context, user *platform.User, account *platform.PaymentAccount) (payments []platform.RecurringPayment, err error) {
	if account != nil {
		payments, err = query.GetRecurringPayments(ctx, m.db, account.UUID)
	} else {
		payments, err = query.GetRecurringPaymentsOfUser(ctx, m.db, user.UUID)
	}

	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return payments, nil
}

func GetRecurringPaymentType(payment *platform.RecurringPayment) (
	// platform.RecurringPaymentType,
	platform.RecurringPaymentChargingMethod,
	error,
) {
	if payment == nil {
		return 0, ErrRecurringPaymentServiceInvalidType
	}

	// typ := platform.RecurringPaymentType(payment.SchedulerType)
	// switch typ {
	// case
	// 	platform.RecurringPaymentYearly,
	// 	platform.RecurringPaymentMonthly,
	// 	platform.RecurringPaymentFixedInterval:
	// default:
	// 	return 0, 0, ErrRecurringPaymentServiceInvalidType
	// }

	method := platform.RecurringPaymentChargingMethod(payment.ChargingMethod)
	switch method {
	case
		platform.RecurringPaymentChargeUpfront,
		platform.RecurringPaymentChargeNonUpfront:
	default:
		return 0, ErrRecurringPaymentServiceInvalidType
	}

	return method, nil
}
