package controller

import (
	"jet/service"

	"github.com/gin-gonic/gin"
)

func SaveBaseLocation(ctx *gin.Context) {
	service.SaveBaseLocation(ctx)
}
