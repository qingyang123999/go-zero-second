package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// http返回
func Response(r *http.Request, w http.ResponseWriter, res interface{}, err error) {
	if err != nil {
		// 根据不同的错误码，返回不同的错误信息
		body := Body{
			Code: -1,
			Msg:  err.Error(),
			Data: nil,
		}
		httpx.WriteJson(w, http.StatusInternalServerError, body)
		return
	}
	body := Body{
		Code: 0,
		Msg:  "success",
		Data: res,
	}
	httpx.WriteJson(w, http.StatusOK, body)
}
