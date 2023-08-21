package handler

type Pagination struct {
	Page     uint64 `form:"page" binding:"required" example:"1"`       // 页码
	PageSize uint64 `form:"page_size" binding:"required" example:"10"` // 每页条数
}
