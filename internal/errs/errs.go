package errs

import "errors"

var (
	ErrSendMessage = errors.New("can't send message\n")
)
