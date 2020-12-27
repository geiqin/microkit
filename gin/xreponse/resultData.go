package xreponse

import "github.com/geiqin/microkit/protobuf/common"

type ResultData struct {
	Entity interface{}   `json:"entity,omitempty"`
	Info   *common.Info  `json:"info,omitempty"`
	Items  interface{}   `json:"items,omitempty"`
	Pager  *common.Pager `json:"pager,omitempty"`
	Error  *common.Error `json:"error,omitempty"`
}

func NewError(code int32, message string) *common.Error {
	return &common.Error{
		Code:    code,
		Message: message,
	}
}
