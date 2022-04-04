package basedic

type BaseDicVo struct {
	Id   string
	Name string
}

func (BaseDicVo) TableName() string {
	return "base_dic"
}
