package routes

import (
	admincontrollers "admin/adminControllers"
	studentcontrollers "admin/studentControllers"

	"github.com/gin-gonic/gin"
)

func AuthStudentsRoutes(router *gin.Engine) {

	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	router.GET("/", studentcontrollers.IndexHandler)

	router.GET("/signup", studentcontrollers.SignUp)
	router.POST("/signup", studentcontrollers.PostSignUp)
	router.GET("/succesfull", studentcontrollers.HandleLogin)

	router.GET("/login", studentcontrollers.Login)
	router.POST("/login", studentcontrollers.PostLogin)
	router.GET("/home", studentcontrollers.Home)
	router.GET("/logout", studentcontrollers.Logout)

}

func AuthAdminRoutes(router *gin.Engine) {

	router.GET("/admin", admincontrollers.Home)
	router.GET("/delete/:Id/:Table", admincontrollers.Delete)
	router.POST("/adddepartment", admincontrollers.AddDepartment)

	router.GET("/addadmin", admincontrollers.AddAdmin)
	router.POST("/addadmin", admincontrollers.AddAdmin)
	router.POST("/adminlogin", admincontrollers.AdminLogin)
	 router.POST("/ajax", admincontrollers.SendAjax)
	router.POST("/updatestudent", admincontrollers.UpdateStudent)
	router.POST("/addstudent", admincontrollers.AddStudent)
	 router.POST("/search", admincontrollers.Search)
	router.GET("/adminlogout", admincontrollers.Logout)
}
