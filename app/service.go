package app

import (
	"context"

	"github.com/juelko/invoice/pkg/finvoice"
	"github.com/juelko/invoice/pkg/message"
)

type InvoiceService interface {
	Deliver(ctx context.Context, rp RequestParams, finv finvoice.Finvoice) error
	Get(ctx context.Context, rp RequestParams, id message.ID) (Invoice, error)
	List(ctx context.Context, rp RequestParams) []Invoice
}

type RequestParams struct {
	Addr          finvoice.AddressID
	Intermadiator finvoice.IntermediatorID
}
