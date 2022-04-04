package basedic

type BaseDic struct {
	Id   string `json:"Id" db:"id"`
	Name string `json:"Name" db:"name"`
}

func (BaseDic) TableName() string {
	return "base_dic"
}
