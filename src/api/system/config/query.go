package config

import (
	"github.com/goxt/dog2/env"
	"github.com/goxt/dog2/util"
)

// 数据结构
type query struct {
	*util.Base
	Name string
}

// 构造函数
func Query(base *util.Base) util.ApiInterface {
	this := &query{Base: base}
	this.Init(this)
	return this
}

// 业务逻辑
func (this *query) Handler() int {

	var data = map[string]interface{}{
		"version":         env.Config.App.Version,
		"appName":         env.Config.App.AppName,
		"appCnName":       env.Config.App.AppCnName,
		"sessionLifeTime": int(env.Config.App.SessionLifeTime),
	}

	return this.SuccessWithData(data)
}
