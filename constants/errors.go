package constants

import "errors"

var (
	MissWechatInitParamErr = errors.New("missing wechat init parameter")
	MissAlipayInitParamErr = errors.New("missing alipay init parameter")
	MissPayPalInitParamErr = errors.New("missing paypal init parameter")
	MissParamErr           = errors.New("missing required parameter")
	MarshalErr             = errors.New("marshal error")
	UnmarshalErr           = errors.New("unmarshal error")
	SignatureErr           = errors.New("signature error")
	VerifySignatureErr     = errors.New("verify signature error")
	CertNotMatchErr        = errors.New("cert not match error")
	GetSignDataErr         = errors.New("get signature data error")
)

/**
自定义错误代码
*/

const (
	Success              = 0
	ProgramIntError      = 500
	RequestTokenFail     = 10000
	InvalidUnameOrPasswd = 10001
	TokenIsExpired       = 10002
	LoginFail            = 10003
	ServerError          = 10004
	BadRequest           = 10005
	ParamsValidateFail   = 10006
	EmergencyClose       = 20001
)

var CustomStatusText = map[int]string{
	Success:              "成功",
	ProgramIntError:      "业务异常", //500错误
	RequestTokenFail:     "request token fail",
	InvalidUnameOrPasswd: "invalid username or password",
	TokenIsExpired:       "token失效",
	LoginFail:            "登录失败",
	ServerError:          "请求内容失败",
	BadRequest:           "无效的请求",
	ParamsValidateFail:   "参数校验不合规",
	EmergencyClose:       "项目暂时无法访问,静待恢复",
}
