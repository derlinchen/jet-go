package basedic

type PageInfo struct {
	PageNum   int
	PageSize  int
	Total     int64
	PageCount int64
	Lists     []*BaseDicVo
}
