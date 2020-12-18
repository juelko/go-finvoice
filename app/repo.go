package app

import (
	"github.com/juelko/invoice/pkg/message"
)

type InvoiceRepo interface {
	New(inv Invoice) error
	Get(id message.ID) (Invoice, error)
	List(rp RequestParams) []Invoice
	Update(inv Invoice) error
	Delete(id message.ID) error
}
