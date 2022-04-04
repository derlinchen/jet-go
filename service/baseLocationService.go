package service

import (
	"jet/bean"

	"github.com/gin-gonic/gin"
)

func SaveBaseLocation(ctx *gin.Context) {
	bean.NewResult(ctx).Success("true")
}
