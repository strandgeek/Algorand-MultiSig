package apiutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Status int
	Error  string `json:"error"`
}

func Abort(ctx *gin.Context, code int) {
	ctx.JSON(code, ApiError{
		Status: code,
		Error:  http.StatusText(code),
	})
}
