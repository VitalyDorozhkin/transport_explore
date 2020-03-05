//Package service http transport
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver

import (
	"context"
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"

	"github.com/VitalyDorozhkin/transport_explore/pkg/models"
)

type errorCreator func(status int, format string, v ...interface{}) error


// GetUserTransport transport interface
type GetUserTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.GetUserRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, res *models.Response) (err error)
}

type getUserTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getUserTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.GetUserRequest, err error) {

	request.Id, err = strconv.Atoi(string(r.URI().QueryArgs().Peek("id")))
	if err != nil {
		return request, t.errorCreator(
			http.StatusBadRequest,
			"Bad request, check the fields.",
			"failed to get incomeId from query: %v",
			err,
		)
	}

	return
}

// EncodeResponse method for encoding response on server side
func (t *getUserTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, res *models.Response) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(res, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewGetUserTransport the transport creator for http requests
func NewGetUserTransport(errorCreator errorCreator) GetUserTransport {
	return &getUserTransport{
		errorCreator: errorCreator,
	}
}







// PostOrderTransport transport interface
type PostOrderTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.PostOrderRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, res *models.Response) (err error)
}

type postOrderTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *postOrderTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.PostOrderRequest, err error) {
	if err = request.UnmarshalJSON(r.Body()); err != nil {
		return models.PostOrderRequest{}, t.errorCreator(http.StatusBadRequest, "failed to decode JSON request: %v", err)
	}

	return
}

// EncodeResponse method for encoding response on server side
func (t *postOrderTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, res *models.Response) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(res, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewPostOrderTransport the transport creator for http requests
func NewPostOrderTransport(errorCreator errorCreator) PostOrderTransport {
	return &postOrderTransport{
		errorCreator: errorCreator,
	}
}





// GetUserCountTransport transport interface
type GetUserCountTransport interface {
	DecodeRequest(ctx fasthttp.RequestCtx, r *fasthttp.Request) (request models.GetUserCountRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, res *models.Response) (err error)
}

type getUserCountTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getUserCountTransport) DecodeRequest(ctx fasthttp.RequestCtx, r *fasthttp.Request) (request models.GetUserCountRequest, err error) {
	id := ctx.UserValue("id").(string)
	request.Id, err = strconv.Atoi(id)
	if id == "" || err != nil {
		return models.GetUserCountRequest{}, t.errorCreator(http.StatusBadRequest, "failed to decode JSON request: %v", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getUserCountTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, res *models.Response) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(res, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewGetUserCountTransport the transport creator for http requests
func NewGetUserCountTransport(errorCreator errorCreator) GetUserCountTransport {
	return &getUserCountTransport{
		errorCreator: errorCreator,
	}
}
// GetOrdersTransport transport interface
type GetOrdersTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.GetUserRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.Response) (err error)
}

type getOrdersTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getOrdersTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.GetUserRequest, err error) {
	return
}

// EncodeResponse method for encoding response on server side
func (t *getOrdersTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.Response) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewGetOrdersTransport the transport creator for http requests
func NewGetOrdersTransport(errorCreator errorCreator) GetOrdersTransport {
	return &getOrdersTransport{
		errorCreator: errorCreator,
	}
}
