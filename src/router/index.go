package router

import (
	"github.com/goxt/dog2/util"
	"github.com/kataras/iris"
)

func Index(ctx iris.Context) {
	_, _ = ctx.HTML(util.WelCome)
}
