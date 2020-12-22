package invoice

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/juelko/invoice/pkg/finvoice"
	"github.com/juelko/invoice/pkg/message"
)

func MakeDeliveryEndpoint(is Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeliverRequest)
		id, err := is.Deliver(ctx, req.Finvoice)
		return DeliverResponse{ID: id, Error: err}, nil
	}
}

type DeliverRequest struct {
	Finvoice finvoice.Finvoice
}

type DeliverResponse struct {
	ID    message.ID `xml:"ID,omitempty"`
	Error error      `xml:"Error,omitempty"`
}

func MakeGetEndpoint(is Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRequest)
		inv, err := is.Get(ctx, req.IntermediatorID, req.AddressID, req.MessageID)
		return GetResponse{Meta: inv.Meta, Data: inv.Data, Error: err}, nil
	}
}

type GetRequest struct {
	IntermediatorID finvoice.IntermediatorID
	AddressID       finvoice.AddressID
	MessageID       message.ID
}

type GetResponse struct {
	Meta  message.Meta      `xml:"Meta,omitempty"`
	Data  finvoice.Finvoice `xml:"Data,omitempty"`
	Error error             `xml:"Error,omitempty"`
}

func MakeListEndpoint(is Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ListRequest)
		metas, err := is.List(ctx, req.IntermediatorID, req.AddressID)
		return ListResponse{Metas: metas, Error: err}, nil
	}
}

type ListRequest struct {
	IntermediatorID finvoice.IntermediatorID
	AddressID       finvoice.AddressID
}

type ListResponse struct {
	Metas []message.Meta `xml:"Meta,omitempty"`
	Error error          `xml:"Error,omitempty"`
}
