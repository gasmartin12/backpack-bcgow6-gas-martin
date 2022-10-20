package web

import "strconv"

type Response struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) (int, Response) {
	if code < 300 {
		return code, Response{strconv.FormatInt(int64(code), 10), data, ""}
	}
	return code, Response{strconv.FormatInt(int64(code), 10), nil, err}
}
