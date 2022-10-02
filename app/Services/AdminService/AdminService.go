package AdminService

import (
	"goweb/app/Models"
	"goweb/app/Models/Admin"
	"math"
)

type AdminService struct {
}

type Pagination struct {
	Page       int         `json:"page,omitempty"`
	PageSize   int         `json:"page_size"`
	NumPages   int         `json:"num_pages"`
	NumResults int64       `json:"num_results"`
	Data       interface{} `json:"data"`
}

var admin Admin.Admin
var admins []Admin.Admin

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (AdminService *AdminService) Lists(param map[string]int) Pagination {
	offset := (param[`page`] - 1) * param[`page_size`]
	where := make(map[string]interface{})
	where[`status`] = param[`status`]

	_ = Models.DbConn.Model(&admin).Where(where).Order(`id desc`).Limit(param[`page_size`]).Offset(offset).Find(&admins)
	var count int64
	Models.DbConn.Model(&admin).Where(where).Count(&count)
	res := Pagination{
		Page:       param[`page`],
		PageSize:   param[`page_size`],
		NumPages:   int(math.Ceil(float64(count / int64(param[`page_size`])))),
		NumResults: count,
		Data:       admins,
	}

	return res
}

func (AdminService *AdminService) Update(param map[string]interface{}) bool {

	where := make(map[string]interface{})
	where["id"] = 1
	res := Models.DbConn.Model(&admin).Where(where).Updates(param)
	if res.RowsAffected > 0 {
		return true
	} else {
		return true
	}
}

func (AdminService *AdminService) Crate(param map[string]interface{}) bool {
	res := Models.DbConn.Model(&admin).Create(param)
	if res.RowsAffected > 0 {
		return true
	} else {
		return true
	}
}

func (AdminService *AdminService) Delete(param map[string]interface{}) bool {
	res := Models.DbConn.Model(&admin).Delete(param)
	if res.RowsAffected > 0 {
		return true
	} else {
		return true
	}
}
