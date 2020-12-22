package invoice

import (
	"github.com/juelko/invoice/pkg/finvoice"
	"github.com/juelko/invoice/pkg/message"
)

type Invoice struct {
	Meta message.Meta
	Data finvoice.Finvoice
}
