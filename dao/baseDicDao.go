package dao

import (
	"errors"
	"jet/bean"
	"jet/bean/basedic"
	"jet/db"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func GetBaseDic(ctx *gin.Context) ([]*basedic.BaseDicVo, error) {
	var result []*basedic.BaseDicVo
	var id = ctx.Query("Id")
	err := db.Link.Where("id=?", id).Find(&result).Error
	return result, err
}

func SearchBaseDic(ctx *gin.Context) (basedic.PageInfo, error) {
	var result = basedic.PageInfo{}
	var lists []*basedic.BaseDicVo
	pageSearch := bean.PageSearch{}
	err := ctx.ShouldBind(&pageSearch)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	pageNum := pageSearch.PageNum
	if pageNum < 0 {
		return result, errors.New("当前页不能小于0")
	}

	pageSize := pageSearch.PageSize
	if pageSize < 0 {
		return result, errors.New("每页条数不能小于0")
	}

	link := db.Link

	id := pageSearch.Item["Id"]
	if id != nil && id != "" {
		link = link.Where("id = ?", id)
	}

	name := pageSearch.Item["Name"]
	if name != nil && name != "" {
		link = link.Where("name = ?", name)
	}

	var total int64
	err = link.Model(&lists).Count(&total).Error
	if err != nil {
		log.Fatal(err)
		return result, err
	}

	err = link.Scopes(db.Paginate(pageNum, pageSize)).Find(&lists).Error
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	result.PageSize = pageSearch.PageSize
	result.PageNum = pageSearch.PageNum
	result.Total = total
	result.PageCount = db.CalcPageCount(total, pageSize)
	result.Lists = lists
	return result, err
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

	err = tx.Debug().Save(&baseDic).Error
	if err != nil {
		tx.Rollback()
		return errors.New("更新失败")
	}
	return nil
}
