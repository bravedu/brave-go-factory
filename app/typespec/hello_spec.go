package typespec

type HelloListReq struct {
	PageReq
	Id int `json:"id" form:"id" binding:"required"`
}

type HelloListResp struct {
	List []HelloList `json:"list"`
}

type HelloList struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
