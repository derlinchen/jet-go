package service

import (
	"github.com/gin-gonic/gin"
	"jet/bean"
	"jet/dao"
	"jet/db"
)

func SaveBaseDic(ctx *gin.Context) {
	// 添加事务
	tx := db.Link.Begin()
	// 对数据操作
	err := dao.SaveBaseDic(ctx, tx)
	// 判断是否操作失败
	if err != nil {
		bean.NewResult(ctx).Error("160", err.Error())
		return
	}
	// 操作成功，返回结果
	bean.NewResult(ctx).Success(true)
	tx.Commit()
}

func GetBaseDic(ctx *gin.Context) {
	err, result := dao.GetBaseDic(ctx)
	if err == nil {
		bean.NewResult(ctx).Success(result)
	} else {
		bean.NewResult(ctx).Error("160", err.Error())
	}
}

func SearchBaseDic(ctx *gin.Context) {
	err, result := dao.SearchBaseDic(ctx)
	if err == nil {
		bean.NewResult(ctx).Success(result)
	} else {
		bean.NewResult(ctx).Error("160", err.Error())
	}
}

func DeleteBaseDic(ctx *gin.Context) {
	tx := db.Link.Begin()
	err := dao.DeleteBaseDic(ctx, tx)
	if err != nil {
		bean.NewResult(ctx).Error("160", err.Error())
		return
	}
	bean.NewResult(ctx).Success(true)
	tx.Commit()
}

func UpdateBaseDic(ctx *gin.Context) {
	// 开启事务
	tx := db.Link.Begin()
	err := dao.UpdateBaseDic(ctx, tx)
	if err != nil {
		bean.NewResult(ctx).Error("160", err.Error())
		return
	}
	bean.NewResult(ctx).Success(true)
	// 事务提交
	tx.Commit()
}
