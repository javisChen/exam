package routers

import (
	controllers "exam/controllers"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	ns := beego.NewNamespace("/exam",

		// 认证
		beego.NSNamespace("/auth/v1",
			beego.NSRouter("/login", &controllers.LoginController{}, "post:Login"),
			beego.NSRouter("/logout", &controllers.LoginController{}, "post:Logout"),
		),

		// 试卷
		beego.NSNamespace("/paper/v1",
			// 试卷列表
			beego.NSRouter("/list", &controllers.PagerController{}, "post:List"),
			// 新建试卷
			beego.NSRouter("/create", &controllers.PagerController{}, "post:Create"),
			// 试卷详情
			beego.NSRouter("/info", &controllers.PagerController{}, "post:Info"),
		),
		// app端
		beego.NSNamespace("/app",

			beego.NSCond(func(context *context.Context) bool {
				fmt.Println("false")
				return true
			}),

			// 试卷
			beego.NSNamespace("/paper/v1",
				// 试卷列表
				beego.NSRouter("/list", &controllers.PagerController{}, "post:List"),
				// 试卷详情
				beego.NSRouter("/info", &controllers.PagerController{}, "post:Info"),
			),

			// 用户试卷
			beego.NSNamespace("/user-paper/v1",
				// 用户开始答卷
				beego.NSRouter("/create", &controllers.PagerController{}, "post:List"),
			),
		),
	)

	beego.AddNamespace(ns)

}
