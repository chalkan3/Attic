package physical_environment

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

func (mw logmw) Create(pe *PhysicalEnvironment) (*PhysicalEnvironment, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Create",
			"type", "Adding a new physical environment",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.Create(pe)
}

func (mw logmw) Update(pe *PhysicalEnvironment) (*PhysicalEnvironment, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Update",
			"type", "Update a physical environment",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.Update(pe)
}

func (mw logmw) Get(pe *PhysicalEnvironment) (*PhysicalEnvironment, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Get",
			"type", "Get a physical environment",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.Get(pe)
}

func (mw logmw) List() ([]PhysicalEnvironment, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "List",
			"type", "List a pe",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.List()
}
