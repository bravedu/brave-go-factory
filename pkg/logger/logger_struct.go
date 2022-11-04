package logger

type Result struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

const recoverErrCode = 10004
