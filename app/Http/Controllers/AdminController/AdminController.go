package AdminController

import (
	"encoding/json"
	"goweb/app/Http/Response"
	"goweb/app/Services/AdminService"
	"net/http"
)

var adminService = AdminService.NewAdminService()

var response = Response.NewResponse()

func Lists(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]int)
	content := make([]byte, r.ContentLength)
	r.Body.Read(content)
	json.Unmarshal(content, &param)
	//fmt.Fprintf(w, "%#v\n", param)

	res := adminService.Lists(param)

	response.JsonResponse(res, 200, `Ok`, w)
}
