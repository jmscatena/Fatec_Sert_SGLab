package routes

import (
	"github.com/jmscatena/Fatec_Sert_SGLab/models/laboratorios"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmscatena/Fatec_Sert_SGLab/controllers"
	"github.com/jmscatena/Fatec_Sert_SGLab/models/administrativo"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
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
			var obj laboratorios.Materiais
			mat.POST("/", func(context *gin.Context) {
				controllers.Add[laboratorios.Materiais](context, &obj)
			})
			mat.GET("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Get[laboratorios.Materiais](context, &obj, uint64(uid))
			})

			mat.GET("/", func(context *gin.Context) {
				controllers.GetAll[laboratorios.Materiais](context, &obj)
			})
			mat.PATCH("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Modify[laboratorios.Materiais](context, &obj, uint64(uid))
			})
			mat.DELETE("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Erase[laboratorios.Materiais](context, &obj, uint64(uid))
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
			var obj laboratorios.Laboratorios
			lab.POST("/", func(context *gin.Context) {
				controllers.Add[laboratorios.Laboratorios](context, &obj)
			})
			lab.GET("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Get[laboratorios.Laboratorios](context, &obj, uint64(uid))
			})

			lab.GET("/", func(context *gin.Context) {
				controllers.GetAll[laboratorios.Laboratorios](context, &obj)
			})
			lab.PATCH("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Modify[laboratorios.Laboratorios](context, &obj, uint64(uid))
			})
			lab.DELETE("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Erase[laboratorios.Laboratorios](context, &obj, uint64(uid))
			})
		}
		res := main.Group("reservas")
		{
			var obj laboratorios.Reservas
			res.POST("/", func(context *gin.Context) {
				controllers.Add[laboratorios.Reservas](context, &obj)
			})
			res.GET("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Get[laboratorios.Reservas](context, &obj, uint64(uid))
			})

			res.GET("/", func(context *gin.Context) {
				controllers.GetAll[laboratorios.Reservas](context, &obj)
			})
			res.PATCH("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Modify[laboratorios.Reservas](context, &obj, uint64(uid))
			})
			res.DELETE("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Erase[laboratorios.Reservas](context, &obj, uint64(uid))
			})
		}
		ges := main.Group("gestao")
		{
			var obj laboratorios.GestaoMateriais
			ges.POST("/", func(context *gin.Context) {
				controllers.Add[laboratorios.GestaoMateriais](context, &obj)
			})
			ges.GET("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Get[laboratorios.GestaoMateriais](context, &obj, uint64(uid))
			})

			ges.GET("/", func(context *gin.Context) {
				controllers.GetAll[laboratorios.GestaoMateriais](context, &obj)
			})
			ges.PATCH("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Modify[laboratorios.GestaoMateriais](context, &obj, uint64(uid))
			})
			ges.DELETE("/:id", func(context *gin.Context) {
				uid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
				controllers.Erase[laboratorios.GestaoMateriais](context, &obj, uint64(uid))
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
