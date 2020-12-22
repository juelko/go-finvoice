package message

import (
	"time"

	"github.com/google/uuid"
)

const TimeFormat = time.RFC3339Nano

func NewMeta(t Type, sender CustomerID, receiver CustomerID, dm DeliveryMethod) Meta {
	return Meta{
		ID:        NewID(),
		Type:      t,
		Sender:    sender,
		Receiver:  receiver,
		Delivery:  dm,
		Status:    StatusNew,
		CreatedAt: time.Now().Format(TimeFormat),
		Version:   1,
	}
}

type Meta struct {
	ID        ID
	Type      Type
	Sender    CustomerID
	Receiver  CustomerID
	Delivery  DeliveryMethod
	Status    Status
	Error     Error
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Version   uint
}

type Error struct {
	Code     string
	Text     string
	Severity string
	Location string
}

func NewID() ID {
	return ID(uuid.New().String())
}

type ID string

type CustomerID string

type Type string

type DeliveryMethod string

const (
	DeliveryMethodREST  = DeliveryMethod("REST")
	DeliveryMethodBatch = DeliveryMethod("BATCH")
)

type Status string

const (
	StatusNew       = Status("New")
	StatusDelivered = Status("Delivered")
	StatusPaid      = Status("Paid")
	StatusRejected  = Status("Rejected")
	StatusDeleted   = Status("Deleted")
)
