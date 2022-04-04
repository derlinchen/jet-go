package routers

import (
	"jet/bean"
	"jet/controller"
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.NoRoute(NoFound)
	router.NoMethod(NoFound)
	router.Use(Recover)
	baseDic := router.Group("/wms/baseDic")
	{
		baseDic.POST("/saveBaseDic", controller.SaveBaseDic)
		baseDic.GET("/getBaseDic", controller.GetBaseDic)
		baseDic.POST("/searchBaseDic", controller.SearchBaseDic)
		baseDic.DELETE("/deleteBaseDic", controller.DeleteBaseDic)
		baseDic.POST("/updateBaseDic", controller.UpdateBaseDic)
	}
	baseLocation := router.Group("/wms/baseLocation")
	{
		baseLocation.POST("/saveBaseLocation", controller.SaveBaseLocation)
	}
	return router
}

func NoFound(c *gin.Context) {
	bean.NewResult(c).Error("404", "未找到接口")
	return
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			bean.NewResult(c).Error("500", "服务器内部错误")
		}
	}()
	c.Next()
}
