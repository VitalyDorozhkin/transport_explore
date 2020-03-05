//Package service http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpclient

import (
	"context"

	"github.com/valyala/fasthttp"

	"github.com/VitalyDorozhkin/transport_explore/pkg/models"
)

const (
	method = "http://"
)

var (

	GetUser = option{}
	PostOrder = option{}
	GetUserCount = option{}
	GetOrders = option{}
)

type option struct{}

// Option ...
type Option interface {
	Prepare(ctx context.Context, r *fasthttp.Request)
}

// Service implements Service interface
type Service interface {

	GetUser(ctx context.Context, request *models.GetUserRequest) (res models.Response, err error)
	PostOrder(ctx context.Context, request *models.PostOrderRequest) (res models.Response, err error)
	GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (res models.Response, err error)
	GetOrders(ctx context.Context) (res models.Response, err error)
}

type client struct {
	cli *fasthttp.HostClient

	transportGetUser GetUserClientTransport
	transportPostOrder PostOrderClientTransport
	transportGetUserCount GetUserCountClientTransport
	transportGetOrders GetOrdersClientTransport
	options map[interface{}]Option
}

// GetUser ...
func (s *client) GetUser(ctx context.Context, request *models.GetUserRequest) (res models.Response, err error) {
	req, ress := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(ress)
	}()
	if opt, ok := s.options[GetUser]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetUser.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, ress)
	if err != nil {
		return
	}
	return s.transportGetUser.DecodeResponse(ctx, ress)
}

// PostOrder ...
func (s *client) PostOrder(ctx context.Context, request *models.PostOrderRequest) (res models.Response, err error) {
	req, ress := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(ress)
	}()
	if opt, ok := s.options[PostOrder]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportPostOrder.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, ress)
	if err != nil {
		return
	}
	return s.transportPostOrder.DecodeResponse(ctx, ress)
}

// GetUserCount ...
func (s *client) GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (res models.Response, err error) {
	req, ress := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(ress)
	}()
	if opt, ok := s.options[GetUserCount]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetUserCount.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, ress)
	if err != nil {
		return
	}
	return s.transportGetUserCount.DecodeResponse(ctx, ress)
}

// GetOrders ...
func (s *client) GetOrders(ctx context.Context) (res models.Response, err error) {
	req, ress := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(ress)
	}()
	if opt, ok := s.options[GetOrders]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetOrders.EncodeRequest(ctx, req); err != nil {
		return
	}
	err = s.cli.Do(req, ress)
	if err != nil {
		return
	}
	return s.transportGetOrders.DecodeResponse(ctx, ress)
}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,
	
	transportGetUser GetUserClientTransport,
	transportPostOrder PostOrderClientTransport,
	transportGetUserCount GetUserCountClientTransport,
	transportGetOrders GetOrdersClientTransport,
	options map[interface{}]Option,
) Service {
	return &client{
		cli: cli,
		
		transportGetUser: transportGetUser,
		transportPostOrder: transportPostOrder,
		transportGetUserCount: transportGetUserCount,
		transportGetOrders: transportGetOrders,
		options: options,
	}
}

// NewPreparedClient create and set up http client
func NewPreparedClient(
	serverURL string, 
	serverHost string, 
	maxConns int, 
	options map[interface{}]Option,
	errorProcessor errorProcessor,
	errorCreator errorCreator,
	
	uriPathGetUser string,
	uriPathPostOrder string,
	uriPathGetUserCount string,
	uriPathGetOrders string,
	
	httpMethodGetUser string,
	httpMethodPostOrder string,
	httpMethodGetUserCount string,
	httpMethodGetOrders string,
) Service {
	serverURL = method + serverURL
	transportGetUser := NewGetUserClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetUser,
		httpMethodGetUser,
	)
	
	transportPostOrder := NewPostOrderClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathPostOrder,
		httpMethodPostOrder,
	)
	
	transportGetUserCount := NewGetUserCountClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetUserCount,
		httpMethodGetUserCount,
	)
	
	transportGetOrders := NewGetOrdersClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetOrders,
		httpMethodGetOrders,
	)
	
	return NewClient(
		&fasthttp.HostClient{
			Addr:     serverHost,
			MaxConns: maxConns,
		},
		
		transportGetUser,
		transportPostOrder,
		transportGetUserCount,
		transportGetOrders,
		options,
	)
}
