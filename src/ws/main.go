package ws

import (
	"github.com/goxt/dog2/dogWebSocket"
	"github.com/goxt/dog2/util"
	"net/http"
)

func Start() {

	// 定义账号在其他地方登录后，强制退出时发送Code值
	dogWebSocket.CodeForcedLogout = "1000"

	// 重新定义底层包的一些方法
	dogWebSocket.Receive = receive
	dogWebSocket.GetUid = getUid
	dogWebSocket.Register = register
	dogWebSocket.UnRegister = unRegister

	// 启动
	dogWebSocket.Run()
}

// 根据请求对象，获取用户ID
func getUid(res http.ResponseWriter, req *http.Request) uint64 {
	return util.GetUidByTokenInSession(res, req)
}

// 连接成功，并加入到底层队列前，触发的回调函数
func register(id uint64) bool {
	return true
}

// 注销前触发的回调函数
func unRegister(id uint64) bool {
	return true
}