package code

import jsoniter "github.com/json-iterator/go"

var (
	NewUserVoucherId int64 = 1 //新手券id标识

)

type Error struct {
	ErrorNo   int    `json:"errno"`
	ErrorMsg  string `json:"errmsg"`
	UserMsg   string `json:"usermsg"`
	ToastType int    `json:"toastType,omitempty"`
	ErrTitle  string `json:"errorTitle,omitempty"`
}

func (e Error) Error() string {
	dataStr, _ := jsoniter.MarshalToString(&e)
	return dataStr
}
