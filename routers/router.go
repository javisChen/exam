package routers

import (
	"exam/controllers/auth_handler"
	"exam/controllers/paper_handler"
	"exam/controllers/user_paper_handler"
	"exam/routers/middleware"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	rootGroup := router.Group("/exam")
	rootGroup.Use(middleware.AuthMiddleWare)
	{
		authGroup := rootGroup.Group("/auth/v1")
		{
			authGroup.POST("/login", auth_handler.Login)
		}

		paperGroup := rootGroup.Group("/paper/v1")
		{
			paperGroup.POST("/list", paper_handler.List)
			paperGroup.POST("/create", paper_handler.Create)
			paperGroup.POST("/info", paper_handler.Info)
		}

		appGroup := rootGroup.Group("/app")
		{
			appPaperGroup := appGroup.Group("/paper/v1")
			{
				appPaperGroup.POST("/list", paper_handler.List)
				appPaperGroup.POST("/create", paper_handler.Create)
				appPaperGroup.POST("/info", paper_handler.Info)
			}
			appUserPaperGroup := appGroup.Group("/user-paper/v1")
			{
				appUserPaperGroup.POST("/create", user_paper_handler.Create)
				appUserPaperGroup.POST("/answer", user_paper_handler.Answer)
				appUserPaperGroup.POST("/finish", user_paper_handler.Finish)
			}
		}

	}

	//ns := beego.NewNamespace("/exam",
	//
	//	// 认证
	//	beego.NSNamespace("/auth/v1",
	//		beego.NSRouter("/login", &controllers.LoginController{}, "post:Login"),
	//		beego.NSRouter("/logout", &controllers.LoginController{}, "post:Logout"),
	//	),
	//
	//	// 试卷
	//	beego.NSNamespace("/paper/v1",
	//		// 试卷列表
	//		beego.NSRouter("/list", &controllers.PagerController{}, "post:List"),
	//		// 新建试卷
	//		beego.NSRouter("/create", &controllers.PagerController{}, "post:Create"),
	//		// 试卷详情
	//		beego.NSRouter("/info", &controllers.PagerController{}, "post:Info"),
	//	),
	//	// app端
	//	beego.NSNamespace("/app",
	//
	//		// 试卷
	//		beego.NSNamespace("/paper/v1",
	//			// 试卷列表
	//			beego.NSRouter("/list", &controllers.PagerController{}, "post:List"),
	//			// 试卷详情
	//			beego.NSRouter("/info", &controllers.PagerController{}, "post:Info"),
	//		),
	//
	//		// 用户试卷
	//		beego.NSNamespace("/user-paper/v1",
	//			// 用户开始答卷
	//			beego.NSRouter("/create", &controllers.UserPaperController{}, "post:Create"),
	//			// 用户答题
	//			beego.NSRouter("/answer", &controllers.UserPaperController{}, "post:Answer"),
	//			// 用户交卷
	//			beego.NSRouter("/finish", &controllers.UserPaperController{}, "post:Finish"),
	//		),
	//	),
	//)
	//
	//beego.AddNamespace(ns)

}
