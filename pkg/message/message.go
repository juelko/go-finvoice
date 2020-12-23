package message

import (
	"time"
)

const TimeFormat = time.RFC3339Nano

type Meta struct {
	ID        ID
	Type      MessageType
	Sender    Party
	Receiver  Party
	Status    Status
	Error     Error
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Version   uint
}

func (m *Meta) SetReceiver(receiver Party) {
	m.Receiver = receiver
}

func (m *Meta) RejectWith(err Error) {
	m.Receiver = Party{}
	m.Status = StatusRejected
	m.Error = err
}

func NewMeta(id ID, t MessageType, sender Party) *Meta {
	return &Meta{
		ID:        id,
		Type:      t,
		Sender:    sender,
		Status:    StatusNew,
		CreatedAt: time.Now().Format(TimeFormat),
		Version:   1,
	}
}

type Error struct {
	Code     string
	Text     string
	Severity string
	Location string
}

type ID string

type MessageType string
type Party struct {
	Channel        ChannelID
	DeliveryMethod DeliveryMethod
	Type           PartyType
}

type ChannelID string

type DeliveryMethod string

type PartyType string

const (
	DeliveryMethodREST  = DeliveryMethod("REST")
	DeliveryMethodBatch = DeliveryMethod("BATCH")
)

type Status string

const (
	StatusNew       = Status("New")
	StatusDelivered = Status("Delivered")
	StatusPending   = Status("Pending")
	StatusPaid      = Status("Paid")
	StatusRejected  = Status("Rejected")
	StatusDeleted   = Status("Deleted")
)
