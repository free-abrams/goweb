package AdminRequest

import (
	"encoding/json"
	"net/http"
)

type AdminRequest struct {
}

func (adminRequest *AdminRequest) handel(r *http.Request) {
	param := make(map[string]string)
	content := make([]byte, r.ContentLength)
	r.Body.Read(content)
	json.Unmarshal(content, &param)
}
