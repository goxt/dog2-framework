package config

const (
	// 项目名
	AppName = "framework"

	// 是否启动Session功能
	OpenSession = true

	// 会话超时时间 (单位:秒)
	SessionLifeTime = 1800

	// 上传文件大小限制 (单位:M)
	FileUploadMax = 20
)

// 业务自定义配置对象，启动服务的时候从json文件中加载
var Config = struct {
}{}
