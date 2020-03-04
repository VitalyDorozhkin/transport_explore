//Package service http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpclient


import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/VitalyDorozhkin/transportexplore/pkg/models"
)

type errorCreator func(status int, format string, v ...interface{}) error

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
	Decode(r *fasthttp.Response) error
}


// GetUserClientTransport transport interface
type GetUserClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.GetUserRequest) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (res models.Response, err error)
}

type getUserClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *getUserClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.GetUserRequest) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)

	r.URI().QueryArgs().Set("id", strconv.Itoa(request.Id))
	return
}

// DecodeResponse method for decoding response on client side
func (t *getUserClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (res models.Response, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = res.UnmarshalJSON(r.Body())
	return
}

// NewGetUserClientTransport the transport creator for http requests
func NewGetUserClientTransport(
	errorProcessor errorProcessor,
	errorCreator   errorCreator,
	pathTemplate   string,
	method         string,
) GetUserClientTransport {
	return &getUserClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}
// PostOrderClientTransport transport interface
type PostOrderClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.PostOrderRequest) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (res models.Response, err error)
}

type postOrderClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *postOrderClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.PostOrderRequest) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)
	r.Header.Set("Content-Type", "application/json")
	r.SetBodyStreamWriter(func(w *bufio.Writer) {
		if err = json.NewEncoder(w).Encode(request); err != nil {
			return
		}
	})
	return
}

// DecodeResponse method for decoding response on client side
func (t *postOrderClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (res models.Response, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = res.UnmarshalJSON(r.Body())
	return
}

// NewPostOrderClientTransport the transport creator for http requests
func NewPostOrderClientTransport(
	errorProcessor errorProcessor,
	errorCreator   errorCreator,
	pathTemplate   string,
	method         string,
) PostOrderClientTransport {
	return &postOrderClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}
// GetUserCountClientTransport transport interface
type GetUserCountClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.GetUserCountRequest) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (res models.Response, err error)
}

type getUserCountClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *getUserCountClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.GetUserCountRequest) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, strconv.Itoa(request.Id)))
	r.Header.Set("Content-Type", "application/json")
	return
}

// DecodeResponse method for decoding response on client side
func (t *getUserCountClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (res models.Response, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = res.UnmarshalJSON(r.Body())
	return
}

// NewGetUserCountClientTransport the transport creator for http requests
func NewGetUserCountClientTransport(
	errorProcessor errorProcessor,
	errorCreator   errorCreator,
	pathTemplate   string,
	method         string,
) GetUserCountClientTransport {
	return &getUserCountClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}
// GetOrdersClientTransport transport interface
type GetOrdersClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request,  ) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (res models.Response, err error)
}

type getOrdersClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *getOrdersClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request,  ) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)
	r.Header.Set("Content-Type", "application/json")
	return
}

// DecodeResponse method for decoding response on client side
func (t *getOrdersClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (res models.Response, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = res.UnmarshalJSON(r.Body())
	return
}

// NewGetOrdersClientTransport the transport creator for http requests
func NewGetOrdersClientTransport(
	errorProcessor errorProcessor,
	errorCreator   errorCreator,
	pathTemplate   string,
	method         string,
) GetOrdersClientTransport {
	return &getOrdersClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}
