package router

import (
	"github.com/goxt/dog2/util"
	"github.com/kataras/iris"
)

type this struct {
	ctx  iris.Context
	path string
	api  *util.Api
}

func Handler(ctx iris.Context) {

	t := this{ctx: ctx}
	isView := false

	// 定义全局异常处理
	defer catchException(t.ctx, &isView)

	// 定义this对象，并获取路由和api对象
	t.getApi()

	// 处理跨域
	t.crossDomain()

	// OPTIONS请求处理
	if t.ctx.Method() == "OPTIONS" {
		util.JsonResponse(util.JsonSuccess("支持跨域访问"), t.ctx)
		return
	}

	// 生成基础API，并获取会话数据
	needLogin := t.api.Login
	baseApi := util.NewBaseApi(needLogin, t.ctx)
	defer func() {
		isView = baseApi.IsView
	}()

	// 校验登录
	if needLogin && !baseApi.Session.IsLogged() {
		util.ThrowBiz(util.MsgNoLogin, util.CodeNoLogin)
	}

	// 操作校验权限
	if needLogin && t.api.AuthKey != "" && !baseApi.HasAuth(t.api.AuthKey) {
		util.ThrowBiz(util.MsgNoAuth, util.CodeNoAuth)
	}

	// 执行接口，失败则关闭数据库连接
	defer baseApi.CloseConnection()
	ins := t.api.Controller(baseApi)
	resType := ins.Handler()

	// 更新会话数据
	baseApi.Session.UpdateSession()

	// 响应数据 (不同类型的数据)
	switch resType {
	case 1:
		util.JsonResponse(baseApi.ResData, baseApi.Ctx)
	case 2:
		return
	}
}

func catchException(ctx iris.Context, isView *bool) {
	e := recover()
	if e == nil {
		return
	}

	switch v := e.(type) {
	case util.BizException:
		if *isView {
			util.BizExceptionViewHandler(v, ctx)
		} else {
			util.BizExceptionHandler(v, ctx)
		}
	case util.SysException:
		if *isView {
			util.SysExceptionViewHandler(ctx)
		} else {
			util.SysExceptionHandler(ctx)
		}
	case string:
		util.LogException(v)
		if *isView {
			util.SysExceptionViewHandler(ctx)
		} else {
			util.SysExceptionHandler(ctx)
		}
	default:
		util.LogException(e.(error).Error())
		if *isView {
			util.SysExceptionViewHandler(ctx)
		} else {
			util.SysExceptionHandler(ctx)
		}
	}
}

func (t *this) getApi() {

	t.path = t.ctx.Path()
	t.api = Router[t.path]

	if t.api == nil {
		util.ThrowBiz("接口不存在")
	}
}

func (t *this) crossDomain() {
	allowHeader := "Origin, Content-Type, Accept, Token"
	allowMethod := "POST, GET, OPTIONS"
	t.ctx.ResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")
	t.ctx.ResponseWriter().Header().Set("Access-Control-Allow-Methods", allowMethod)
	t.ctx.ResponseWriter().Header().Set("Access-Control-Allow-Headers", allowHeader)
	t.ctx.ResponseWriter().Header().Set("Access-Control-Max-Age", "3600")
}
