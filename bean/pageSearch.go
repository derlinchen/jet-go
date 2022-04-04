package bean

type PageSearch struct {
	PageNum int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Item map[string] interface{} `json:"item"`
}
