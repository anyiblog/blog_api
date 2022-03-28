package serializer

// Response 基础序列化器
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

const (
	SystemOk = 0

	SystemError = 1

	RedisSms = 3

	// TokenRedisTime Token保存时间（天）
	TokenRedisTime = 5

	// AuthFailed 未登录
	AuthFailed = 401

	// AuthIllegal 认证非法，拒绝访问
	AuthIllegal = 403

	// NotBindPhone 未绑定手机
	NotBindPhone = 410

	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
)
