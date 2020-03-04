package httpclient

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"

	"github.com/VitalyDorozhkin/transportexplore/pkg/models"
	"github.com/VitalyDorozhkin/transportexplore/pkg/service"
	"github.com/VitalyDorozhkin/transportexplore/pkg/service/httpserver"
)

const (
	serverAddr               = "localhost:8080"
	serverHost               = "localhost:8080"
	maxConns                 = 512
	maxRequestBodySize       = 15 * 1024 * 1024
	serverTimeout            = 1 * time.Millisecond
	serverLaunchingWaitSleep = 1 * time.Second

	// service name method
	getUser      = "GetUser"
	postOrder    = "PostOrder"
	getUserCount = "GetUserCount"
	getOrders    = "GetOrders"
	ID           = 12
)

func Test_client_GetUser(t *testing.T) {
	t.Run("TestGetUserSuccess", func(t *testing.T) {
		request := makeGetUserRequest()
		response := makeResponse()
		serviceMock := new(service.MockService)
		serviceMock.On(getUser, context.Background(), request).Return(response, nil)
		server, client := makeServerClient(serverAddr, serviceMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		resp, err := client.GetUser(context.Background(), request)
		assert.Equal(t, resp, response)
		assert.NoError(t, err, "unexpected error:", err)

	})
}

func Test_client_PostOrder(t *testing.T) {
	t.Run("TestPostOrderSuccess", func(t *testing.T) {
		request := makePostOrderRequest()
		response := makeResponse()
		serviceMock := new(service.MockService)
		serviceMock.On(postOrder, context.Background(), request).Return(response, nil)
		server, client := makeServerClient(serverAddr, serviceMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)
		resp, err := client.PostOrder(context.Background(), request)
		assert.Equal(t, resp, response)
		assert.NoError(t, err, "unexpected error:", err)

	})
}

func Test_client_GetUserCount(t *testing.T) {
	t.Run("TestGetUserCountSuccess", func(t *testing.T) {
		request := makeGetUserCountRequest()
		response := makeResponse()
		serviceMock := new(service.MockService)
		serviceMock.On(getUserCount, context.Background(), request).Return(response, nil)
		server, client := makeServerClient(serverAddr, serviceMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		resp, err := client.GetUserCount(context.Background(), request)
		assert.Equal(t, resp, response)
		assert.NoError(t, err, "unexpected error:", err)

	})
}

func Test_client_GetOrders(t *testing.T) {
	t.Run("TestGetOrdersSuccess", func(t *testing.T) {
		response := models.Response{}
		serviceMock := new(service.MockService)
		serviceMock.On(getOrders, context.Background()).Return(response, nil)
		server, client := makeServerClient(serverAddr, serviceMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		resp, err := client.GetOrders(context.Background())
		assert.Equal(t, resp, response)
		assert.NoError(t, err, "unexpected error:", err)

	})
}

func makeServerClient(serverAddr string, svc Service) (server *fasthttp.Server, client Service) {
	client = NewPreparedClient(
		serverAddr,
		serverHost,
		maxConns,
		nil,
		httpserver.NewErrorProcessor(http.StatusInternalServerError, "Iternal Error"),
		httpserver.NewError,
		URIPathClientGetUser,
		URIPathClientPostOrder,
		URIPathClientGetUserCount,
		URIPathClientGetOrders,
		HTTPMethodGetUser,
		HTTPMethodPostOrder,
		HTTPMethodGetUserCount,
		HTTPMethodGetOrders,
	)
	router := httpserver.NewPreparedServer(svc)
	server = &fasthttp.Server{
		Handler:            router.Handler,
		MaxRequestBodySize: maxRequestBodySize,
		ReadTimeout:        serverTimeout,
	}

	go func() {
		err := server.ListenAndServe(serverAddr)
		if err != nil {
			log.Printf("server shut down err: %v", err)
		}
	}()
	return
}

func makeGetUserRequest() *models.GetUserRequest {
	return &models.GetUserRequest{
		Id: ID,
	}
}

func makePostOrderRequest() *models.PostOrderRequest {
	return &models.PostOrderRequest{
		Id: ID,
	}
}

func makeGetUserCountRequest() *models.GetUserCountRequest {
	return &models.GetUserCountRequest{
		Id: ID,
	}
}

func makeResponse() models.Response {
	return models.Response{
		Data: &models.Data{Res: true},
	}
}
