package base

import (
	"energia/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

// BaseResponse is the base response for all the endpoints
// @Description BaseResponse is the base response for all the endpoints
// @Param Status bool true "Status of the response"
// @Param Message string true "Message of the response"
// @Param Data any true "Data of the response"
type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "success",
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, err error) error {
	return c.JSON(helper.GetResponseCodeFromErr(err), BaseResponse{
		Status:  false,
		Message: err.Error(),
	})
}
