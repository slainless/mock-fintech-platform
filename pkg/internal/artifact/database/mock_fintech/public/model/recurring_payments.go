//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type RecurringPayments struct {
	UUID           uuid.UUID `sql:"primary_key"`
	ServiceID      string
	AccountUUID    uuid.UUID
	SchedulerType  int16
	LastCharge     *time.Time
	ForeignID      string
	ChargingMethod int16
}
