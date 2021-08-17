package mysql

// 根据页码和分页大小计算偏移量
func calculatePageAndOffset(page, size int) int {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * size
	return offset
}
