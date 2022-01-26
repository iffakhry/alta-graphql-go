package common

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Code    int    `json:"code"`
		Status  string `json:"status"`
		Message string `json:"message"`
		// Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, message string, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Code = http.StatusOK
	response.Meta.Status = "success"
	response.Meta.Message = message
	response.Data = param

	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, message string) error {
	response := BaseResponse{}
	response.Meta.Code = status
	response.Meta.Status = "failed"
	response.Meta.Message = message

	return c.JSON(status, response)
}

// func NewErrorResponse(c echo.Context, status int, message string, err error) error {
// 	response := BaseResponse{}
// 	response.Meta.Code = status
// 	response.Meta.Status = "failed"
// 	response.Meta.Message = message
// 	response.Meta.Messages = []string{err.Error()}

// 	return c.JSON(status, response)
// }
