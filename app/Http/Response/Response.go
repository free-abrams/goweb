package Response

import (
	"encoding/json"
	jsonResponse "goweb/app/Http/Response/JsonResponnse"
	"net/http"
)

type Response struct {
}

func NewResponse() *Response {
	return &Response{}
}

func (Response *Response) JsonResponse(data any, code int, message string, w http.ResponseWriter) {
	jsonRes := jsonResponse.JsonResponse{
		Code:    code,
		Result:  data,
		Message: message,
	}

	res, _ := json.Marshal(jsonRes)

	w.Write(res)
}
