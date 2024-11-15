package helper

import (
	"energia/constant"
	"net/http"
)

func GetResponseCodeFromErr(err error) int {
	switch err {
	case constant.EMAIL_NOT_FOUND,
		constant.EMAIL_IS_EMPTY,
		constant.PASSWORD_IS_EMPTY,
		constant.INVALID_DEVICE_ID,
		constant.DEVICE_NOT_FOUND,
		constant.USER_ID_NOT_FOUND_IN_TOKEN,
		constant.INVALID_TOKEN_CLAIMS,
		constant.INVALID_DATE_FORMAT:
		return http.StatusBadRequest
	case constant.INVALID_PASSWORD:
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}
