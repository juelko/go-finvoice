package invoice

import (
	"context"
	"errors"

	"github.com/juelko/invoice/pkg/finvoice"
	"github.com/juelko/invoice/pkg/message"
)

type Service interface {
	Deliver(ctx context.Context, finv finvoice.Finvoice) (message.ID, error)
	Get(ctx context.Context, iid finvoice.IntermediatorID, aid finvoice.AddressID, id message.ID) (Invoice, error)
	List(ctx context.Context, iid finvoice.IntermediatorID, aid finvoice.AddressID) ([]message.Meta, error)
}

func NewInvoiceService(repo Repository) Service {
	return &invoiceSvc{repo: repo}
}

type invoiceSvc struct {
	repo Repository
}

func (svc *invoiceSvc) Deliver(ctx context.Context, finv finvoice.Finvoice) (message.ID, error) {
	return message.ID(""), errors.New("Not implemented")
}

func (svc *invoiceSvc) Get(ctx context.Context, iid finvoice.IntermediatorID, aid finvoice.AddressID, id message.ID) (Invoice, error) {
	return Invoice{}, errors.New("Not implemented")
}

func (svc *invoiceSvc) List(ctx context.Context, iid finvoice.IntermediatorID, aid finvoice.AddressID) ([]message.Meta, error) {
	return nil, errors.New("Not implemented")
}
