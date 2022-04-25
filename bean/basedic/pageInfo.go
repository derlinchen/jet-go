package basedic

type PageInfo struct {
	PageNum   int          `json:"pageNum"`
	PageSize  int          `json:"pageSize"`
	Total     int64        `json:"total"`
	PageCount int64        `json:"pageCount"`
	Lists     []*BaseDicVo `json:"lists"`
}
