package basedic

import "jet/utils"

type BaseDic struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}

func NewBaseDic() BaseDic {
	snow, _ := utils.NewSnowFlake()
	return BaseDic{
		Id: snow.Generate(),
	}
}

func (BaseDic) TableName() string {
	return "base_dic"
}
