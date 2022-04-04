package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jet/bean"
	"jet/bean/basedic"
	"jet/db"
	"log"
)

func SaveBaseDic(ctx *gin.Context, tx *gorm.DB) error {
	var baseDic basedic.BaseDic
	err := ctx.ShouldBind(&baseDic)
	if err != nil {
		return err
	}

	exec := tx.Save(&baseDic)
	if exec.Error != nil {
		tx.Rollback()
		return errors.New("保存失败")
	}
	return nil
}

func GetBaseDic(ctx *gin.Context) (error, []*basedic.BaseDicVo) {
	var result []*basedic.BaseDicVo
	var id = ctx.Query("Id")
	err := db.Link.Where("id=?", id).Find(&result).Error
	return err, result
}

func SearchBaseDic(ctx *gin.Context) (error, basedic.PageInfo) {
	var result = basedic.PageInfo{}
	var lists []*basedic.BaseDicVo
	pageSearch := bean.PageSearch{}
	err := ctx.ShouldBind(&pageSearch)
	if err != nil {
		log.Fatal(err)
		return err, result
	}
	pageNum := pageSearch.PageNum
	if pageNum < 0 {
		if pageNum < 0 {
			return errors.New("当前页不能小于0"), result
		}
	}
	pageSize := pageSearch.PageSize
	if pageSize < 0 {
		return errors.New("每页条数不能小于0"), result
	}

	id := pageSearch.Item["Id"]
	name := pageSearch.Item["Name"]
	link := db.Link
	if id != nil && id != "" {
		link = link.Where("id = ?", id)
	}

	if name != nil && name != "" {
		link = link.Where("name = ?", name)
	}

	var total int64
	err = link.Model(&lists).Count(&total).Error
	if err != nil {
		log.Fatal(err)
		return err, result
	}

	err = link.Scopes(db.Paginate(pageNum, pageSize)).Find(&lists).Error
	if err != nil {
		log.Fatal(err)
		return err, result
	}
	result.PageSize = pageSearch.PageSize
	result.PageNum = pageSearch.PageNum
	result.Total = total
	result.PageCount = db.CalcPageCount(total, pageSize)
	result.Lists = lists
	return err, result
}

func DeleteBaseDic(ctx *gin.Context, tx *gorm.DB) error {
	id := ctx.Query("Id")
	exec := tx.Delete(basedic.BaseDic{}, "id = ?", id)
	if exec.Error != nil {
		tx.Rollback()
		return errors.New("删除失败")
	}
	return nil
}

func UpdateBaseDic(ctx *gin.Context, tx *gorm.DB) error {
	// 定义接收变量
	var baseDic basedic.BaseDic
	// 对变量进行绑定
	err := ctx.ShouldBind(&baseDic)
	if err != nil {
		return err
	}

	// 执行sql
	exec := tx.Debug().Exec("update base_dic set name = ? where id = ?", baseDic.Name, baseDic.Id)
	// sql执行存在问题，则进行事务回滚
	if exec.Error != nil {
		// 事务回滚
		tx.Rollback()
		return errors.New("更新失败")
	}
	return nil
}
