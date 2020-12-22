package invoice

import (
	"github.com/juelko/invoice/pkg/finvoice"
	"github.com/juelko/invoice/pkg/message"
)

type Repository interface {
	New(inv Invoice) error
	Get(id message.ID) (Invoice, error)
	List(iid finvoice.IntermediatorID, aid finvoice.AddressID) []Invoice
	Update(inv Invoice) error
	Delete(id message.ID) error
}
