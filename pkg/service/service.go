package service

import (
	"context"
	"errors"

	"github.com/VitalyDorozhkin/transport_explore/pkg/models"
)

type Service interface {
	GetUser(ctx context.Context, request *models.GetUserRequest) (res models.Response, err error)
	PostOrder(ctx context.Context, request *models.PostOrderRequest) (res models.Response, err error)
	GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (res models.Response, err error)
	GetOrders(ctx context.Context) (res models.Response, err error)
}

type service struct {
}

func (s *service) GetUser(ctx context.Context, request *models.GetUserRequest) (res models.Response, err error) {
	if request.Id > 0 {
		res.Data = &models.Data{Res: true}
		return
	}
	res.Error = true
	res.ErrorText = "bad id"
	err = errors.New("bad id")
	return
}

func (s *service) PostOrder(ctx context.Context, request *models.PostOrderRequest) (res models.Response, err error) {
	if request.Id > 0 {
		res.Data = &models.Data{Res: true}
		return
	}
	res.Error = true
	res.ErrorText = "bad id"
	err = errors.New("bad id")
	return
}

func (s *service) GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (res models.Response, err error) {
	if request.Id > 0 {
		res.Data = &models.Data{Res: true}
		return
	}
	res.Error = true
	res.ErrorText = "bad id"
	err = errors.New("bad id")
	return
}

func (s *service) GetOrders(ctx context.Context) (res models.Response, err error) {
	return
}

func NewService() Service {
	return &service{}
}
