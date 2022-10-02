package Controllers

import (
	"encoding/json"
	"goweb/app/Models"
	"goweb/common"
	"net/http"
)

type NewType struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type Books struct {
	Id      int    `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	content := make([]byte, r.ContentLength)
	r.Body.Read(content)

	user := Models.User{}
	json.Unmarshal(content, &user)

	data, _ := Models.SelectOne(user)
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func Json(w http.ResponseWriter, r *http.Request) {
	temp := NewType{
		Type: "aaa",
		Name: "cccc",
	}
	res, _ := json.Marshal(temp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func BookMap(w http.ResponseWriter, r *http.Request) {
	books := make(map[int]*Books)

	book1 := Books{1, "未知作者", "来自外太空的标题", "来自火星的描述"}
	book2 := Books{2, "我是作者", "标题", "描述"}
	books[book1.Id] = &book1
	books[book2.Id] = &book2
	str := common.JsonEncode(books)
	//str, _ := json.Marshal(books)
	datas := make(map[int]*Books)

	string := `{"1":{"id":1,"author":"未知作者","title":"来自外太空的标题","summary":"来自火星的描述"},"2":{"id":2,"author":"我是作者","title":"标题","summary":"描述"}}`
	common.JsonDecode(string, &datas)
	w.Header().Set(`content-type`, `application/json`)
	w.Write(str)
}

func UserPost(w http.ResponseWriter, r *http.Request) {

}
