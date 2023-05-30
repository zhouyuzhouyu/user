package views

import (
	"net/http"
	"reflect"

	"zy-apilabs/internal/code"

	"github.com/gin-gonic/gin"
)

type BaseView map[string]any

// Option custom setup config
type Option func(*option)

type option struct {
	ErrorNo   int
	ErrorMsg  string
	UserMsg   string
	ToastType int
	ErrTitle  string
}

func WithCode(code int) Option {
	return func(opt *option) {
		opt.ErrorNo = code
	}
}
func WithType(toastType int) Option {
	return func(opt *option) {
		opt.ToastType = toastType
	}
}
func WithError(err string) Option {
	return func(opt *option) {
		opt.ErrorMsg = err
	}
}
func WithMessage(msg string) Option {
	return func(opt *option) {
		opt.UserMsg = msg
	}
}
func WithTitle(title string) Option {
	return func(opt *option) {
		opt.ErrTitle = title
	}
}
func Success(values ...any) BaseView {
	view := make(BaseView)
	view.AddError(code.Error{ErrorNo: http.StatusOK, ErrorMsg: "SUCC"})
	if len(values) > 0 {
		view.AddData(values[0])
	}
	return view
}
func (view BaseView) Add(key string, value any) BaseView {
	view[key] = value
	return view
}
func (view BaseView) AddData(value any) BaseView {
	view["retData"] = value
	return view
}
func (view BaseView) AddError(value any) BaseView {
	view["error"] = value
	return view
}
func Failed(options ...Option) BaseView {
	view := make(BaseView)
	opt := new(option)
	for _, f := range options {
		f(opt)
	}
	if opt.ErrorNo == 0 {
		opt.ErrorNo = http.StatusOK
	}
	base := code.Error{
		ErrorNo:  opt.ErrorNo,
		ErrorMsg: opt.ErrorMsg,
		UserMsg:  opt.UserMsg,
	}
	view.AddError(base)
	return view
}

func FailedByMsgType(ctx *gin.Context, msgCode int, msgType string) BaseView {
	base := make(BaseView)
	base.Add("error", code.Error{ErrorNo: msgCode, ErrorMsg: msgType, UserMsg: msgType})
	return base
}
func FailedByError(ctx *gin.Context, err error, userMsgType string) BaseView {
	if err == nil {
		return FailedByMsgType(ctx, http.StatusForbidden, userMsgType)
	}
	aErr, ok := err.(code.Error)
	if ok {
		var userMsg string
		if aErr.UserMsg == "common error" {
			userMsg = userMsgType
		} else {
			userMsg = aErr.UserMsg
		}
		view := make(BaseView)
		view.Add("error", code.Error{ErrorNo: aErr.ErrorNo, ErrorMsg: aErr.ErrorMsg, UserMsg: userMsg})
		return view
	}
	return FailedByMsgType(ctx, http.StatusForbidden, userMsgType)
}

func FailedByErrCode(ctx *gin.Context, msgCode int) BaseView {
	base := make(BaseView)
	base.Add("error", code.Error{ErrorNo: msgCode})
	return base
}

func (view BaseView) AddList(key, listKey string, value any) BaseView {
	reflVal := reflect.ValueOf(value)
	if reflVal.Kind() == reflect.Slice {
		listMap := make(map[string]any)
		listMap["num"] = reflVal.Len()
		listMap[listKey] = value
		view[key] = listMap
		return view
	}

	return view.Add(key, value)
}
