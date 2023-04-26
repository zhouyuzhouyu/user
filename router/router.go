package router

import (
	"net/http"

	v1 "com.zoyu/user/handler/v1"
	"com.zoyu/user/middleware"
)

func Router() http.Handler {
	// 创建路由器
	r := http.NewServeMux()
	// 注册路由
	for _, route := range v1.Routes {
		// 添加HTTP方法中间件
		h := middleware.Chain(route.Handler, middleware.Method(route.Method), middleware.Logging())
		r.HandleFunc(route.Path, h)
	}

	return r
}
