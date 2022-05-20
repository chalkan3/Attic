package items

import (
	"time"

	"github.com/go-kit/log"
)

type logmw struct {
	logger log.Logger
	svc    Service
}

func NewLogMW(logger log.Logger, next Service) logmw {
	return logmw{logger, next}
}

func (mw logmw) Create(item *Item) (*Item, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Create",
			"type", "Adding a new physical environment",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.Create(item)
}

func (mw logmw) Update(item *Item) (*Item, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Update",
			"type", "Update a physical environment",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.Update(item)
}

func (mw logmw) Get(item *Item) (*Item, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Get",
			"type", "Get a physical environment",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.Get(item)
}

func (mw logmw) List() ([]Item, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "List",
			"type", "List a pe",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.List()
}
