package Controllers

import (
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "用户列表")
}
