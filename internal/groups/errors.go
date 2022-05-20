package groups

import "errors"

var (
	ErrUserIDNotExist = errors.New("User don't Exist")
	ErrPeIDNotExist   = errors.New("Physical Environment don't Exist")
	ErrUserDontMatch  = errors.New("User don't match with Physical Environment")
)
