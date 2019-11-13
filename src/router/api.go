package router

import (
	"api/system/config"
	"github.com/goxt/dog2/api/framework"
	"github.com/goxt/dog2/util"
)

var Router = map[string]*util.Api{
	// 框架级接口
	"/system/sign/login":    {false, "", framework.Login},
	"/system/sign/logout":   {true, "", framework.Logout},
	"/system/sign/resetPwd": {false, "", framework.ResetPwd},
	"/system/file/upload":   {true, "", framework.Upload},

	// 业务基础接口
	"/system/config/query": {true, "", config.Query},

	// 业务接口
}
