package biedeptrai

import "errors"

var (
	ErrOrderNotFound = errors.New("order not found")
	ErrOrderRecevied = errors.New("đơn hàng đã có người nhận")
)
