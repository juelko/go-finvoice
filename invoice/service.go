package invoice

import (
	"context"
	"errors"

	"github.com/juelko/invoice/pkg/finvoice"
	"github.com/juelko/invoice/pkg/message"
)

//Exported errors
var (
	ErrSenderMismatch = errors.New("MessageSenderDetails does not match parameters")
)

type Service interface {
	Send(ctx context.Context, params SendParameters) (message.ID, error)
	Get(ctx context.Context, params GetParameters) (Invoice, error)
	List(ctx context.Context, params ListParameters) ([]message.Meta, error)
}

type SendParameters struct {
	Sender        message.Party
	Intermadiator finvoice.IntermediatorID
	Address       finvoice.AddressID
	Finvoice      *finvoice.Finvoice
}

type GetParameters struct {
	Intermadiator finvoice.IntermediatorID
	Address       finvoice.AddressID
	ID            message.ID
}

type ListParameters struct {
	Intermadiator finvoice.IntermediatorID
	Address       finvoice.AddressID
}

func NewInvoiceService(repo Repository, receivers ReceiverService) Service {
	return &invoiceSvc{
		repo:      repo,
		receivers: receivers,
	}
}

type invoiceSvc struct {
	repo      Repository
	receivers ReceiverService
}

func (svc *invoiceSvc) Send(ctx context.Context, params SendParameters) (message.ID, error) {

	err := validateParams(params)
	if err != nil {
		return message.ID(""), err
	}

	meta, err := createMeta(svc.receivers, params)
	if err != nil {
		return message.ID(""), err
	}

	inv := Invoice{Meta: meta, Data: params.Finvoice}

	err = svc.repo.New(inv)

	return inv.Meta.ID, nil
}

func (svc *invoiceSvc) Get(ctx context.Context, params GetParameters) (Invoice, error) {
	return Invoice{}, errors.New("Not implemented")
}

func (svc *invoiceSvc) List(ctx context.Context, params ListParameters) ([]message.Meta, error) {
	return nil, errors.New("Not implemented")
}

func validateParams(params SendParameters) error {
	return matchSenderIntermadiator(params)
}

func matchSenderIntermadiator(params SendParameters) error {

	if params.Intermadiator != params.Finvoice.FromIntermediator() {
		return ErrSenderMismatch
	}

	return matchSenderAddress(params)
}

func matchSenderAddress(params SendParameters) error {

	if params.Address != params.Finvoice.From() {
		return ErrSenderMismatch
	}

	return nil
}

func createMeta(receivers ReceiverService, params SendParameters) (*message.Meta, error) {
	meta := newMeta(params.Sender, params.Finvoice)

	receiver, err := receivers.GetReceiverParty(params.Finvoice.ToIntermediator(), params.Finvoice.To())
	if err != nil {
		return handleReceiverError(meta, err)

	}
	meta.SetReceiver(receiver)

	return meta, nil
}

func handleReceiverError(meta *message.Meta, err error) (*message.Meta, error) {
	if err == ErrIntermadiatorNotFound {
		return handleIntermadiatorNotFound(meta, err)
	}
	if err == ErrReceiverNotFound {
		return handleReceiverNotfound(meta, err)
	}

	return nil, err
}

func handleIntermadiatorNotFound(meta *message.Meta, err error) (*message.Meta, error) {
	meta.RejectWith(message.Error{Code: "4", Text: err.Error()})
	return meta, nil
}

func handleReceiverNotfound(meta *message.Meta, err error) (*message.Meta, error) {
	meta.RejectWith(message.Error{Code: "4", Text: err.Error()})
	return meta, nil
}

func newMeta(sender message.Party, inv *finvoice.Finvoice) *message.Meta {
	return message.NewMeta(
		inv.MessageID(),
		inv.MessageType(),
		sender,
	)
}
