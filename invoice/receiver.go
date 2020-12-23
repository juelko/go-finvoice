package invoice

import (
	"errors"

	"github.com/juelko/invoice/pkg/finvoice"
	"github.com/juelko/invoice/pkg/message"
)

var (
	ErrIntermadiatorNotFound = errors.New("Receiveing Intermediator not found")
	ErrReceiverNotFound      = errors.New("Receiver Address not found")
)

type ReceiverService interface {
	GetReceiverParty(iid finvoice.IntermediatorID, aid finvoice.AddressID) (message.Party, error)
}
