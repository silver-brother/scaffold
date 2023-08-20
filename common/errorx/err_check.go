package errorx

import "database/sql"

// CheckDBErr 检查是否有错误
func CheckDBErr(e error) (exists bool, err error) {
	if e == sql.ErrNoRows {
		return false, nil
	}
	if e != nil {
		return true, e
	}
	return false, nil
}
