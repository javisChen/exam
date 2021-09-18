package routers

import (
	controllers "exam/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/exam",

		// 认证
		beego.NSNamespace("/auth/v1/",
			beego.NSRouter("/login", &controllers.LoginController{}, "post:Login"),
			beego.NSRouter("/logout", &controllers.LoginController{}, "post:Logout"),
		),

		// 试卷
		beego.NSNamespace("/paper/v1/",
			beego.NSRouter("/list", &controllers.PagerController{}, "post:List"),
		),
	)

	beego.AddNamespace(ns)

}
