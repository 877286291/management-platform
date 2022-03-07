package response

import "encoding/json"

type Response struct {
	Code   int         `json:"code"` // 错误码
	Msg    string      `json:"msg"`  // 错误描述
	Data   interface{} `json:"data"` // 返回数据
	Size   *int        `json:"size,omitempty"`
	Page   *int        `json:"page,omitempty"`
	Total  *int64      `json:"total,omitempty"`
	ErrMsg interface{} `json:"errMsg,omitempty"` // 错误信息

}

func (r Response) WithMsg(message string) Response {
	return Response{
		Code:   r.Code,
		Msg:    message,
		Data:   r.Data,
		ErrMsg: r.ErrMsg,
	}
}
func (r Response) WithData(data interface{}) Response {
	return Response{
		Code:   r.Code,
		Msg:    r.Msg,
		Data:   data,
		ErrMsg: r.ErrMsg,
	}
}
func (r Response) WithErrMsg(err error) Response {
	return Response{
		Code:   r.Code,
		Msg:    r.Msg,
		Data:   r.Data,
		ErrMsg: err.Error(),
	}
}
func (r Response) WithPagination(page, size *int, total *int64) Response {
	return Response{
		Code:   r.Code,
		Msg:    r.Msg,
		Data:   r.Data,
		Page:   page,
		Size:   size,
		Total:  total,
		ErrMsg: r.ErrMsg,
	}
}
func (r *Response) ToString() string {
	raw, _ := json.Marshal(r)
	return string(raw)
}
func response(code int, msg string) *Response {
	return &Response{
		Code:   code,
		Msg:    msg,
		Data:   nil,
		ErrMsg: nil,
	}
}
