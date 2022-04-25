package basedic

import "jet/utils"

type BaseDicVo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewBaseDicVo() BaseDicVo {
	snow, _ := utils.NewSnowFlake()
	return BaseDicVo{
		Id: snow.Generate(),
	}
}

func (BaseDicVo) TableName() string {
	return "base_dic"
}
