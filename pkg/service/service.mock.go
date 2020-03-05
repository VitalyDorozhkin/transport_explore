package service

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/VitalyDorozhkin/transport_explore/pkg/models"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetUser(ctx context.Context, request *models.GetUserRequest) (res models.Response, err error) {
	args := m.Called(context.Background(),request)
	if a, ok := args.Get(0).(models.Response); ok{
		return a, args.Error(1)
	}
	return res, args.Error(1)
}

func (m *MockService) PostOrder(ctx context.Context, request *models.PostOrderRequest) (res models.Response, err error) {
	args := m.Called(context.Background(),request)
	if a, ok := args.Get(0).(models.Response); ok{
		return a, args.Error(1)
	}
	return res, args.Error(1)
}

func (m *MockService) GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (res models.Response, err error) {
	args := m.Called(context.Background(),request)
	if a, ok := args.Get(0).(models.Response); ok{
		return a, args.Error(1)
	}
	return res, args.Error(1)
}

func (m *MockService) GetOrders(ctx context.Context) (res models.Response, err error) {
	args := m.Called(context.Background())
	if a, ok := args.Get(0).(models.Response); ok{
		return a, args.Error(1)
	}
	return res, args.Error(1)
}
