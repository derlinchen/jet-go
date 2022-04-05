package basedic

import "jet/utils"

type BaseDic struct {
	Id   string `json:"Id" db:"id"`
	Name string `json:"Name" db:"name"`
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
