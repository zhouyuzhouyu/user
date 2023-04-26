package v1

import (
	"com.zoyu/user/middleware"
)

// 路由列表
var Routes = []middleware.Route{
	{"/registe", "POST", registe},
	{"/login", "POST", login},
	{"/logout", "POST", logout},
}
