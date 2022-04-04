package basedic

type BaseDicVo struct {
	Id   string `db:"id", json:"Id"`
	Name string `db:"name", json:"Name"`
}

func (BaseDicVo) TableName() string {
	return "base_dic"
}
