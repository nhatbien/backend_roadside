package biedeptrai

import "errors"

var (
	ErrorSaveRescueUnitFail   = errors.New("lưu đơn vị cứu hộ thất bại")
	ErrorRescueUnitNotFound   = errors.New("đơn vị cứu hộ không tồn tại")
	ErrorPasswordNotMatch     = errors.New("mật khẩu sai")
	ErrorUpdateRescueUnitFail = errors.New("cập nhật vị trí đơn vị cứu hộ thất bại")
	ErrorOrdersNotFound       = errors.New("không tìm thấy đơn nào cứu hộ")
)
