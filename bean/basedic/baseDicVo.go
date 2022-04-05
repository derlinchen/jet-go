package basedic

import "jet/utils"

type BaseDicVo struct {
	Id   string
	Name string
}

func NewBaseDicVo() BaseDicVo {
	snow := utils.SnowFlake{}
	return BaseDicVo{
		Id: snow.Generate(),
	}
}

func (BaseDicVo) TableName() string {
	return "base_dic"
}
