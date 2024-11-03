package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmscatena/Fatec_Sert_SGLab/controllers"
	"github.com/jmscatena/Fatec_Sert_SGLab/database/models/administrativo"
	laboratorios2 "github.com/jmscatena/Fatec_Sert_SGLab/database/models/laboratorios"
	"log"
	"net/http"
	"strconv"
	"time"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
		login := main.Group("/login")
		{
			login.POST("/", gin.BasicAuth(gin.Accounts{
				"admin": "secret",
			}), func(c *gin.Context) {
				// Create a new token object, specifying signing method and the claims
				// you would like it to contain.
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"foo": "bar",
					"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
				})

				// Sign and get the complete encoded token as a string using the secret
				tokenString, err := token.SignedString(hmacSampleSecret)

				fmt.Println(tokenString, err)
				c.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			})
		}
		user := main.Group("user")
		{

			var obj administrativo.Usuario
			user.POST("/", func(context *gin.Context) {
				controllers.Add[administrativo.Usuario](context, &obj)

				// Parse takes the token string and a function for looking up the key. The latter is especially
				// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
				// head of the token to identify which key to use, but the parsed token (head and claims) is provided
				// to the callback, providing flexibility.
				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
					return hmacSampleSecret, nil
				})
				if err != nil {
					log.Fatal(err)
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					fmt.Println(claims["foo"], claims["nbf"])
				} else {
					fmt.Println(err)
				}

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
