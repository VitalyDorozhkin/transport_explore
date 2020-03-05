//Package service logging wrapper
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/VitalyDorozhkin/transport_explore/pkg/models"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	logger log.Logger
	svc    Service
}

func (s *loggingMiddleware) GetUser(ctx context.Context, request *models.GetUserRequest) (res models.Response, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetUser",
			"request", request,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetUser(ctx, request)
}

func (s *loggingMiddleware) PostOrder(ctx context.Context, request *models.PostOrderRequest) (res models.Response, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "PostOrder",
			"request", request,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.PostOrder(ctx, request)
}

func (s *loggingMiddleware) GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (res models.Response, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetUserCount",
			"request", request,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetUserCount(ctx, request)
}

func (s *loggingMiddleware) GetOrders(ctx context.Context) (res models.Response, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetOrders",
			
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetOrders(ctx, )
}

func (s *loggingMiddleware) wrap(err error) log.Logger {
	lvl := level.Debug
	if err != nil {
		lvl = level.Error
	}
	return lvl(s.logger)
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(logger log.Logger, svc Service) Service {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
