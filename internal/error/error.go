package error

import "errors"

const (
	LoginUserErrorMsg           = "invalid email or password"
	CreateUserBadInputErrorMsg  = "invalid registration data"
	NothingToUpdateUserErrorMsg = "nothing to update"
	BusyUpdateEmailErrorMsg     = "email is busy"
	UserNotFoundErrorMsg        = "not found"
)

var (
	NotFoundError        = errors.New(UserNotFoundErrorMsg)
	NothingToUpdateError = errors.New(NothingToUpdateUserErrorMsg)
	LoginError           = errors.New(LoginUserErrorMsg)
	BusyUpdateEmailError = errors.New(BusyUpdateEmailErrorMsg)
)
