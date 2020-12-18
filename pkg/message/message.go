package message

import (
	"time"

	"github.com/google/uuid"
)

const TimeFormat = time.RFC3339Nano

const (
	StatusNew       = Status("New")
	StatusDelivered = Status("Delivered")
	StatusPaid      = Status("Paid")
	StatusRejected  = Status("Rejected")
	StatusDeleted   = Status("Deleted")
)

func NewMeta(t Type, sid SenderID, rid ReceiverID, dm DeliveryMethod) Meta {
	ret := Meta{
		ID:        NewID(),
		Type:      t,
		Sender:    sid,
		Receiver:  rid,
		Delivery:  dm,
		Status:    StatusNew,
		CreatedAt: time.Now().Format(TimeFormat),
		Version:   1,
	}

	return ret
}

type Meta struct {
	ID        ID
	Type      Type
	Sender    SenderID
	Receiver  ReceiverID
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

type SenderID string

type ReceiverID string

type Type string

type DeliveryMethod string

type Status string
