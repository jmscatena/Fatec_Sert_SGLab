package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmscatena/Fatec_Sert_SGLab/controllers"
	"github.com/jmscatena/Fatec_Sert_SGLab/database/models/administrativo"
	laboratorios2 "github.com/jmscatena/Fatec_Sert_SGLab/database/models/laboratorios"
	"strconv"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
		/*login := main.Group("/login")
		{

		}*/
		user := main.Group("user")
		{

			var obj administrativo.Usuario
			user.POST("/", func(context *gin.Context) {
				controllers.Add[administrativo.Usuario](context, &obj)
			})
			user.GET("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Get[administrativo.Usuario](context, &obj, uint64(uid))
			})

			user.GET("/", func(context *gin.Context) {
				controllers.GetAll[administrativo.Usuario](context, &obj)
			})
			user.GET("/admin/", func(context *gin.Context) {
				//colocar as configuracoes para os params q virao do frontend
				params := "admin=?;ativo=?"
				controllers.GetBy[administrativo.Usuario](context, &obj, params, false, true)
			})

			user.PATCH("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Modify[administrativo.Usuario](context, &obj, uint64(uid))
			})
			user.DELETE("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Erase[administrativo.Usuario](context, &obj, uint64(uid))
			})

		}
		mat := main.Group("materiais")
		{
			var obj laboratorios2.Materiais
			mat.POST("/", func(context *gin.Context) {
				controllers.Add[laboratorios2.Materiais](context, &obj)
			})
			mat.GET("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Get[laboratorios2.Materiais](context, &obj, uint64(uid))
			})

			mat.GET("/", func(context *gin.Context) {
				controllers.GetAll[laboratorios2.Materiais](context, &obj)
			})
			mat.PATCH("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Modify[laboratorios2.Materiais](context, &obj, uint64(uid))
			})
			mat.DELETE("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Erase[laboratorios2.Materiais](context, &obj, uint64(uid))
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
				controllers.Add[laboratorios2.Laboratorios](context, &obj)
			})
			lab.GET("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Get[laboratorios2.Laboratorios](context, &obj, uint64(uid))
			})

			lab.GET("/", func(context *gin.Context) {
				controllers.GetAll[laboratorios2.Laboratorios](context, &obj)
			})
			lab.PATCH("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Modify[laboratorios2.Laboratorios](context, &obj, uint64(uid))
			})
			lab.DELETE("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Erase[laboratorios2.Laboratorios](context, &obj, uint64(uid))
			})
		}
		res := main.Group("reservas")
		{
			var obj laboratorios2.Reservas
			res.POST("/", func(context *gin.Context) {
				controllers.Add[laboratorios2.Reservas](context, &obj)
			})
			res.GET("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Get[laboratorios2.Reservas](context, &obj, uint64(uid))
			})

			res.GET("/", func(context *gin.Context) {
				controllers.GetAll[laboratorios2.Reservas](context, &obj)
			})
			res.PATCH("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Modify[laboratorios2.Reservas](context, &obj, uint64(uid))
			})
			res.DELETE("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Erase[laboratorios2.Reservas](context, &obj, uint64(uid))
			})
		}
		ges := main.Group("gestao")
		{
			var obj laboratorios2.GestaoMateriais
			ges.POST("/", func(context *gin.Context) {
				controllers.Add[laboratorios2.GestaoMateriais](context, &obj)
			})
			ges.GET("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Get[laboratorios2.GestaoMateriais](context, &obj, uint64(uid))
			})

			ges.GET("/", func(context *gin.Context) {
				controllers.GetAll[laboratorios2.GestaoMateriais](context, &obj)
			})
			ges.PATCH("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Modify[laboratorios2.GestaoMateriais](context, &obj, uint64(uid))
			})
			ges.DELETE("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Erase[laboratorios2.GestaoMateriais](context, &obj, uint64(uid))
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
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Get[models.Instituicao](context, &obj, uint64(uid))
				})
				inst.PATCH("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Modify[models.Instituicao](context, &obj, uint64(uid))
				})
				inst.DELETE("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Erase[models.Instituicao](context, &obj, uint64(uid))
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
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Get[models.Evento](context, &obj, uint64(uid))
				})
				event.PATCH("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Modify[models.Evento](context, &obj, uint64(uid))
				})
				event.DELETE("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Erase[models.Evento](context, &obj, uint64(uid))
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
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Get[models.Certificado](context, &obj, uint64(uid))
				})
				cert.PATCH("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Modify[models.Certificado](context, &obj, uint64(uid))
				})
				cert.DELETE("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Erase[models.Certificado](context, &obj, uint64(uid))
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
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Get[models.CertVal](context, &obj, uint64(uid))
				})
				certval.PATCH("/", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Modify[models.CertVal](context, &obj, uint64(uid))
				})
				certval.DELETE("/:id", func(context *gin.Context) {
					uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
					controllers.Erase[models.CertVal](context, &obj, uint64(uid))
				})
			}
		*/
		//main.POST("login", controllers.Login)
	}

	return router
}
