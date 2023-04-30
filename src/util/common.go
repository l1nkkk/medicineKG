package util

type HttpResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	HttpOK = iota + 1
	HttpParaError
	HttpInternalError
)
