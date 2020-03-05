//Package service instrumenting wrapper
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package service

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/VitalyDorozhkin/transport_explore/pkg/models"
)

// instrumentingMiddleware wraps Service and enables request metrics
type instrumentingMiddleware struct {
	reqCount    metrics.Counter
	reqDuration metrics.Histogram
	svc         Service
}

func (s *instrumentingMiddleware) GetUser(ctx context.Context, request *models.GetUserRequest) (res models.Response, err error) {
	defer s.recordMetrics("GetUser", time.Now(), err)
	return s.svc.GetUser(ctx, request)
}

func (s *instrumentingMiddleware) PostOrder(ctx context.Context, request *models.PostOrderRequest) (res models.Response, err error) {
	defer s.recordMetrics("PostOrder", time.Now(), err)
	return s.svc.PostOrder(ctx, request)
}

func (s *instrumentingMiddleware) GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (res models.Response, err error) {
	defer s.recordMetrics("GetUserCount", time.Now(), err)
	return s.svc.GetUserCount(ctx, request)
}

func (s *instrumentingMiddleware) GetOrders(ctx context.Context,  ) (res models.Response, err error) {
	defer s.recordMetrics("GetOrders", time.Now(), err)
	return s.svc.GetOrders(ctx, )
}

func (s *instrumentingMiddleware) recordMetrics(method string, startTime time.Time, err error) {
	labels := []string{
		"method", method,
		"error", strconv.FormatBool(err != nil),
	}
	s.reqCount.With(labels...).Add(1)
	s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
}

// NewInstrumentingMiddleware ...
func NewInstrumentingMiddleware(reqCount metrics.Counter, reqDuration metrics.Histogram, svc Service) Service {
	return &instrumentingMiddleware{
		reqCount:    reqCount,
		reqDuration: reqDuration,
		svc:         svc,
	}
}
