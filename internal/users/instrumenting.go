package users

import (
	"time"
)

type instrumentingMiddleware struct {
	metrics *Metrics
	next    Service
}

func NewInstrumentingMiddleware(metrics *Metrics, next Service) *instrumentingMiddleware {
	return &instrumentingMiddleware{
		metrics: metrics,
		next:    next,
	}
}

func (mw instrumentingMiddleware) Create(user *User) (*User, error) {
	defer func(begin time.Time) {
		lvsRequests := []string{"method", "create", "type", "requestsCount"}
		lvsLatency := []string{"method", "create", "type", "requestsLatency"}
		mw.metrics.latency.With(lvsLatency...).Observe(float64(time.Since(begin).Seconds()))
		mw.metrics.requestCounter.With(lvsRequests...).Add(1)

	}(time.Now())
	return mw.next.Create(user)
}

func (mw instrumentingMiddleware) Update(user *User) (*User, error) {
	defer func(begin time.Time) {
		lvsRequests := []string{"method", "update", "type", "requestsCount"}
		lvsLatency := []string{"method", "update", "type", "requestsLatency"}
		mw.metrics.latency.With(lvsLatency...).Observe(float64(time.Since(begin).Seconds()))
		mw.metrics.requestCounter.With(lvsRequests...).Add(1)

	}(time.Now())
	return mw.next.Update(user)
}

func (mw instrumentingMiddleware) Get(user *User) (*User, error) {
	defer func(begin time.Time) {
		lvsRequests := []string{"method", "get", "type", "requestsCount"}
		lvsLatency := []string{"method", "get", "type", "requestsLatency"}
		mw.metrics.latency.With(lvsLatency...).Observe(float64(time.Since(begin).Seconds()))
		mw.metrics.requestCounter.With(lvsRequests...).Add(1)

	}(time.Now())
	return mw.next.Get(user)
}

func (mw instrumentingMiddleware) List() ([]User, error) {
	defer func(begin time.Time) {
		lvsRequests := []string{"method", "List", "type", "requestsCount"}
		lvsLatency := []string{"method", "List", "type", "requestsLatency"}
		mw.metrics.latency.With(lvsLatency...).Observe(float64(time.Since(begin).Seconds()))
		mw.metrics.requestCounter.With(lvsRequests...).Add(1)

	}(time.Now())
	return mw.next.List()
}
