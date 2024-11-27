package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmscatena/Fatec_Sert_SGLab/database/models/administrativo"
	laboratorios2 "github.com/jmscatena/Fatec_Sert_SGLab/database/models/laboratorios"
	"github.com/jmscatena/Fatec_Sert_SGLab/middleware"
	"github.com/jmscatena/Fatec_Sert_SGLab/services"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
		/*login := main.Group("/login")
		{

		}*/
		userRoute := main.Group("user", services.Authenticate())
		{
			var user administrativo.Usuario
			userRoute.POST("/", func(context *gin.Context) {
				middleware.Add[administrativo.Usuario](context, &user)
			})
			userRoute.GET("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Get[administrativo.Usuario](context, &user, uid)
			})

			userRoute.GET("/", func(context *gin.Context) {
				middleware.GetAll[administrativo.Usuario](context, &user)
			})
			userRoute.GET("/admin/", func(context *gin.Context) {
				//colocar as configuracoes para os params q virao do frontend
				params := "admin=?;ativo=?"
				middleware.GetBy[administrativo.Usuario](context, &user, params, false, true)
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

		matRoute := main.Group("materiais")
		{
			var mat laboratorios2.Materiais
			matRoute.POST("/", func(context *gin.Context) {
				middleware.Add[laboratorios2.Materiais](context, &mat)
			})
			matRoute.GET("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Get[laboratorios2.Materiais](context, &mat, uid)
			})

			matRoute.GET("/", func(context *gin.Context) {
				middleware.GetAll[laboratorios2.Materiais](context, &mat)
			})
			matRoute.PATCH("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Modify[laboratorios2.Materiais](context, &mat, uid)
			})
			matRoute.DELETE("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Erase[laboratorios2.Materiais](context, &mat, uid)
			})

			/*
				mat.GET("/admin/", func(context *gin.Context) {
					//colocar as configuracoes para os params q virao do frontend
					params := "admin=?;ativo=?"
					controllers.GetBy[administrativo.Usuario](context, &obj, params, false, true)
				})
			*/

		}
		lab := main.Group("laboratorios")
		{
			var obj laboratorios2.Laboratorios
			lab.POST("/", func(context *gin.Context) {
				middleware.Add[laboratorios2.Laboratorios](context, &obj)
			})
			lab.GET("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Get[laboratorios2.Laboratorios](context, &obj, uid)
			})

			lab.GET("/", func(context *gin.Context) {
				middleware.GetAll[laboratorios2.Laboratorios](context, &obj)
			})
			lab.PATCH("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Modify[laboratorios2.Laboratorios](context, &obj, uid)
			})
			lab.DELETE("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Erase[laboratorios2.Laboratorios](context, &obj, uid)
			})
		}
		res := main.Group("reservas")
		{
			var obj laboratorios2.Reservas
			res.POST("/", func(context *gin.Context) {
				middleware.Add[laboratorios2.Reservas](context, &obj)
			})
			res.GET("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Get[laboratorios2.Reservas](context, &obj, uid)
			})

			res.GET("/", func(context *gin.Context) {
				middleware.GetAll[laboratorios2.Reservas](context, &obj)
			})
			res.PATCH("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Modify[laboratorios2.Reservas](context, &obj, uid)
			})
			res.DELETE("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Erase[laboratorios2.Reservas](context, &obj, uid)
			})
		}
		ges := main.Group("gestao")
		{
			var obj laboratorios2.GestaoMateriais
			ges.POST("/", func(context *gin.Context) {
				middleware.Add[laboratorios2.GestaoMateriais](context, &obj)
			})
			ges.GET("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Get[laboratorios2.GestaoMateriais](context, &obj, uid)
			})

			ges.GET("/", func(context *gin.Context) {
				middleware.GetAll[laboratorios2.GestaoMateriais](context, &obj)
			})
			ges.PATCH("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Modify[laboratorios2.GestaoMateriais](context, &obj, uid)
			})
			ges.DELETE("/:id", func(context *gin.Context) {
				uid, _ := uuid.Parse(context.Param("id"))
				middleware.Erase[laboratorios2.GestaoMateriais](context, &obj, uid)
			})
		}
		/*
			inst := main.Group("instituicao")
			{
				var obj models.Instituicao
				inst.POST("/", func(context *gin.Context) {
					controllers.Add[models.Instituicao](context, &obj)
				})
				inst.GET("/", func(context *gin.Context) {
					controllers.GetAll[models.Instituicao](context, &obj)
				})
				inst.GET("/:id", func(context *gin.Context) {
					uid, _ := uuid.Parse(context.Param("id"))
					controllers.Get[models.Instituicao](context, &obj, uuid)
				})
				inst.PATCH("/:id", func(context *gin.Context) {
					uid, _ := uuid.Parse(context.Param("id"))
					controllers.Modify[models.Instituicao](context, &obj, uuid)
				})
				inst.DELETE("/:id", func(context *gin.Context) {
					uid, _ := uuid.Parse(context.Param("id"))
					controllers.Erase[models.Instituicao](context, &obj, uuid)
				})
			}
			event := main.Group("evento")
			{
				var obj models.Evento
				event.POST("/", func(context *gin.Context) {
					controllers.Add[models.Evento](context, &obj)
				})
				event.GET("/", func(context *gin.Context) {
					controllers.GetAll[models.Evento](context, &obj)
				})
				event.GET("/:id", func(context *gin.Context) {
					uid, _ := uuid.Parse(context.Param("id"))
					controllers.Get[models.Evento](context, &obj, uuid)
				})
				event.PATCH("/:id", func(context *gin.Context) {
					uid, _ := uuid.Parse(context.Param("id"))
					controllers.Modify[models.Evento](context, &obj, uuid)
				})
				event.DELETE("/:id", func(context *gin.Context) {
					uid, _ := uuid.Parse(context.Param("id"))
					controllers.Erase[models.Evento](context, &obj, uuid)
				})
			}
			cert := main.Group("cert")
			{
				var obj models.Certificado
				cert.POST("/", func(context *gin.Context) {
					controllers.Add[models.Certificado](context, &obj)
				})
				cert.GET("/", func(context *gin.Context) {
					controllers.GetAll[models.Certificado](context, &obj)
				})
				cert.GET("/:id", func(context *gin.Context) {
					uid, _ := uuid.Parse(context.Param("id"))
					controllers.Get[models.Certificado](context, &obj, uuid)
				})
				cert.PATCH("/:id", func(context *gin.Context) {
					uid, _ := uuid.Parse(context.Param("id"))
					controllers.Modify[models.Certificado](context, &obj, uuid)
				})
				cert.DELETE("/:id", func(context *gin.Context) {
					uid, _ := uuid.Parse(context.Param("id"))
					controllers.Erase[models.Certificado](context, &obj, uuid)
				})
			}
			certval := main.Group("valida")
			{
				var obj models.CertVal
				certval.POST("/", func(context *gin.Context) {
					controllers.Add[models.CertVal](context, &obj)
				})
				certval.GET("/", func(context *gin.Context) {
					controllers.GetAll[models.CertVal](context, &obj)
				})
				certval.GET("/:id", func(context *gin.Context) {
					uid, _ := uuid.Parse(context.Param("id"))
					controllers.Get[models.CertVal](context, &obj, uuid)
				})
				certval.PATCH("/", func(context *gin.Context) {
					uid, _ := uuid.Parse(context.Param("id"))
					controllers.Modify[models.CertVal](context, &obj, uuid)
				})
				certval.DELETE("/:id", func(context *gin.Context) {
					uid, _ := uuid.Parse(context.Param("id"))
					controllers.Erase[models.CertVal](context, &obj, uuid)
				})
			}
		*/
		//main.POST("login", controllers.Login)
	}

	return router
}
