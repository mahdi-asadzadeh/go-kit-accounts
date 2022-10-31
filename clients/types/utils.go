package types

import "github.com/gin-gonic/gin"

func APIErrorResponse(ctx *gin.Context, StatusCode int, Method string, Error interface{}) {
	jsonResponse := ErrorResponse{
		StatusCode: StatusCode,
		Method:     Method,
		Error:      Error,
	}
	ctx.JSON(StatusCode, jsonResponse)
}
