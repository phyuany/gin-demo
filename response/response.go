package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
