package users

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

func (mw logmw) Create(user *User) (*User, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Create",
			"type", "Adding a new user",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.Create(user)
}

func (mw logmw) Update(user *User) (*User, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Update",
			"type", "Update a user",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.Update(user)
}

func (mw logmw) Get(user *User) (*User, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Get",
			"type", "Get a user",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.Get(user)
}

func (mw logmw) List() ([]User, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "List",
			"type", "List a user",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.List()
}
