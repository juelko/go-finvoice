package invoice

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/juelko/invoice/pkg/finvoice"
	"github.com/juelko/invoice/pkg/message"
)

func MakeSendEndpoint(is Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SendRequest)
		id, err := is.Send(ctx, req.Parameters)
		return SendResponse{ID: id, Error: err}, nil
	}
}

type SendRequest struct {
	Parameters SendParameters
	Finvoice   *finvoice.Finvoice
}

type SendResponse struct {
	ID    message.ID `xml:"ID,omitempty"`
	Error error      `xml:"Error,omitempty"`
}

func MakeGetEndpoint(is Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRequest)
		inv, err := is.Get(ctx, req.Parameters)
		return GetResponse{Meta: inv.Meta, Data: inv.Data, Error: err}, nil
	}
}

type GetRequest struct {
	Parameters GetParameters
}

type GetResponse struct {
	Meta  message.Meta      `xml:"Meta,omitempty"`
	Data  finvoice.Finvoice `xml:"Data,omitempty"`
	Error error             `xml:"Error,omitempty"`
}

func MakeListEndpoint(is Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ListRequest)
		metas, err := is.List(ctx, req.Parameters)
		return ListResponse{Metas: metas, Error: err}, nil
	}
}

type ListRequest struct {
	Parameters ListParameters
}

type ListResponse struct {
	Metas []message.Meta `xml:"Meta,omitempty"`
	Error error          `xml:"Error,omitempty"`
}
