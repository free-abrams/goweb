package Request

import (
	"encoding/json"
	"log"
	"net/http"
)

var Request *http.Request

func Init(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 打印请求 URL 明细
		url, _ := json.Marshal(r.URL)
		log.Println(string(url))
		Request = r

		next.ServeHTTP(w, r)
	})
}

func JsonParam() []byte {
	content := make([]byte, Request.ContentLength)
	Request.Body.Read(content)

	return content
}
