package constant

import "errors"

var (
	EMAIL_NOT_FOUND   = errors.New("email not found")
	EMAIL_IS_EMPTY    = errors.New("email kosong")
	PASSWORD_IS_EMPTY = errors.New("password kosong")
	INVALID_PASSWORD  = errors.New("password salah")

	NAME_IS_EMPTY     = errors.New("name kosong")
	POWER_IS_EMPTY    = errors.New("power kosong")
	INVALID_DEVICE_ID = errors.New("invalid device ID")
	DEVICE_NOT_FOUND  = errors.New("device not found")
)
