package router

import (
	ac "go-ibooking/internal/adapter/v1/admin/controllers"
	p "go-ibooking/internal/enum/permission"
	t "go-ibooking/internal/model/technical"
)

func addAuthRoute(c ac.AuthController) t.Routes {
	return t.Routes{
		// login
		t.Route{
			Name: "Login", Method: "POST", Path: "login", Action: string(p.NONE), Subject: string(p.AUTH), HandlerFunc: c.Login,
		},
		// refresh
		t.Route{
			Name: "Refresh", Method: "POST", Path: "refresh", Action: string(p.NONE), Subject: string(p.AUTH), HandlerFunc: c.Refresh,
		},
	}
}
