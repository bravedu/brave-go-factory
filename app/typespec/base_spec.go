package typespec

type PageReq struct {
	PageSize int `json:"page_size" form:"page_size"`
	Offset   int `json:"offset" form:"offset"`
}
