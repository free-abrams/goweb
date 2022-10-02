package routes

import "goweb/app/Http/Controllers/AdminController"

type route []WebRoute

func init() {
	routes := route{
		WebRoute{
			Name:        "admin.lists",
			Method:      "get",
			Pattern:     "/admin/lists",
			HandlerFunc: AdminController.Lists,
		},
		WebRoute{
			Name:        "",
			Method:      "",
			Pattern:     "",
			HandlerFunc: nil,
		},
	}

	for _, v := range routes {
		webRoutes = append(webRoutes, v)
	}
}
