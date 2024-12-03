package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmscatena/Fatec_Sert_SGLab/database/models/administrativo"
	"github.com/jmscatena/Fatec_Sert_SGLab/middleware"
	"github.com/jmscatena/Fatec_Sert_SGLab/services"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
		login := main.Group("login")
		{
			login.POST("/", services.Login())
		}
		userRoute := main.Group("user", services.Authenticate())
		{
			var user administrativo.Usuario
			userRoute.POST("/", func(context *gin.Context) {
				middleware.Add[administrativo.Usuario](context, &user)
			})
			userRoute.GET("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				condition := "Id=?"
				middleware.Get[administrativo.Usuario](context, &user, condition, uid.String())
			})

			userRoute.GET("/", func(context *gin.Context) {
				middleware.GetAll[administrativo.Usuario](context, &user)
			})
			userRoute.GET("/admin/", func(context *gin.Context) {
				//colocar as configuracoes para os params q virao do frontend
				params := "admin=?;ativo=?"
				middleware.Get[administrativo.Usuario](context, &user, params, "false; true")
			})

			userRoute.PATCH("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Modify[administrativo.Usuario](context, &user, uid)
			})
			userRoute.DELETE("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Erase[administrativo.Usuario](context, &user, uid)
			})

		}
	}

	return router
}
