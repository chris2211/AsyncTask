package common

import "errors"

var (
	ERR_LOCK_ALREADY_REQUIRED = errors.New("lock already used")

	ERR_NO_LOCAL_IP_FOUND = errors.New("can not found network IP")
)
