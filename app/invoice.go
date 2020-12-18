package app

import (
	"github.com/juelko/invoice/pkg/finvoice"
	"github.com/juelko/invoice/pkg/message"
)

type Invoice struct {
	message.Meta
	finvoice.Finvoice
}
