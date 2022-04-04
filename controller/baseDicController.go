package controller

import (
	"jet/service"

	"github.com/gin-gonic/gin"
)

func SaveBaseDic(ctx *gin.Context) {
	service.SaveBaseDic(ctx)
}

func GetBaseDic(ctx *gin.Context) {
	service.GetBaseDic(ctx)
}

func SearchBaseDic(ctx *gin.Context) {
	service.SearchBaseDic(ctx)
}

func DeleteBaseDic(ctx *gin.Context) {
	service.DeleteBaseDic(ctx)
}

func UpdateBaseDic(ctx *gin.Context) {
	service.UpdateBaseDic(ctx)
}
