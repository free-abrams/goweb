package jsonResponse

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Code    int    `json:"code"`
	Result  any    `json:"result"`
	Message string `json:"message"`
}

func NewJsonResponse() *JsonResponse {
	return &JsonResponse{}
}

func (jsonClass *JsonResponse) JsonResponse(w http.ResponseWriter, data any) {
	var jsonRes = JsonResponse{
		Code:    200,
		Result:  nil,
		Message: "OK",
	}
	jsonRes.Result = data
	res, _ := json.Marshal(jsonRes)

	w.Write(res)
}
