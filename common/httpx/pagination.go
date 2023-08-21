package httpx

func Pagination(page, pageSize uint64) (limit, offset int32) {
	if page < 1 {
		page = 1
	}

	if pageSize > 30 {
		pageSize = 30
	}
	offset = int32((page - 1) * pageSize)
	return int32(pageSize), offset
}
