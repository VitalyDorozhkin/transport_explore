//Package service http server
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver

import (
	"context"
	"net/http"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"github.com/VitalyDorozhkin/transport_explore/pkg/models"
)

type errorProcessor2 interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type service interface {

	GetUser(ctx context.Context, request *models.GetUserRequest) (res models.Response, err error)
	PostOrder(ctx context.Context, request *models.PostOrderRequest) (res models.Response, err error)
	GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (res models.Response, err error)
	GetOrders(ctx context.Context) (res models.Response, err error)
}


type getUserServer struct {
	transport      GetUserTransport
	service        service
	errorProcessor errorProcessor2
}

// ServeHTTP implements http.Handler.
func (s *getUserServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	res, err := s.service.GetUser(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &res); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetUserServer the server creator
func NewGetUserServer(transport GetUserTransport, service service, errorProcessor errorProcessor2) fasthttp.RequestHandler {
	ls := getUserServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
type postOrderServer struct {
	transport      PostOrderTransport
	service        service
	errorProcessor errorProcessor2
}

// ServeHTTP implements http.Handler.
func (s *postOrderServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	res, err := s.service.PostOrder(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &res); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewPostOrderServer the server creator
func NewPostOrderServer(transport PostOrderTransport, service service, errorProcessor errorProcessor2) fasthttp.RequestHandler {
	ls := postOrderServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
type getUserCountServer struct {
	transport      GetUserCountTransport
	service        service
	errorProcessor errorProcessor2
}

// ServeHTTP implements http.Handler.
func (s *getUserCountServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(*ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	res, err := s.service.GetUserCount(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &res); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetUserCountServer the server creator
func NewGetUserCountServer(transport GetUserCountTransport, service service, errorProcessor errorProcessor2) fasthttp.RequestHandler {
	ls := getUserCountServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getOrdersServer struct {
	transport      GetOrdersTransport
	service        service
	errorProcessor errorProcessor2
}

// ServeHTTP implements http.Handler.
func (s *getOrdersServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	_, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	res, err := s.service.GetOrders(ctx)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &res); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetOrdersServer the server creator
func NewGetOrdersServer(transport GetOrdersTransport, service service, errorProcessor errorProcessor2) fasthttp.RequestHandler {
	ls := getOrdersServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}


func NewPreparedServer(svc service) *fasthttprouter.Router {
	errorProcessor := NewErrorProcessor(http.StatusInternalServerError, "internal error")
	getUserTransport := NewGetUserTransport(NewError)
	postOrderServer := NewPostOrderTransport(NewError)
	getCountTransport := NewGetUserCountTransport(NewError)
	getOrderTransport := NewGetOrdersTransport(NewError)

	return MakeFastHTTPRouter(
		[]*HandlerSettings{
			{
				Path:    URIPathGetUser,
				Method:  http.MethodGet,
				Handler: NewGetUserServer(getUserTransport, svc, errorProcessor),
			},
			{
				Path:    URIPathPostOrder,
				Method:  http.MethodPost,
				Handler: NewPostOrderServer(postOrderServer, svc, errorProcessor),
			},
			{
				Path:    URIPathGetUserCount,
				Method:  http.MethodGet,
				Handler: NewGetUserCountServer(getCountTransport, svc, errorProcessor),
			},
			{
				Path:    URIPathGetOrders,
				Method:  http.MethodGet,
				Handler: NewGetOrdersServer(getOrderTransport, svc, errorProcessor),
			},
		},
	)
}
