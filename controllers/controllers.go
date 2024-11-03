package controllers

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jmscatena/Fatec_Sert_SGLab/services"
	"github.com/jmscatena/Fatec_Sert_SGLab/utils"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "acesso " + http.StatusText(200)})

}

func Add[T utils.Tables](c *gin.Context, o utils.PersistenceHandler[T]) {
	if reflect.TypeOf(o) != nil {
		if err := c.ShouldBindJSON(&o); err != nil {
			fmt.Println("ERRO:", err)
			//msg para deploy
			//c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": "Erro de JSON"})
			//msg para dev
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": err})
			return
		}
		var handler utils.PersistenceHandler[T] = o
		code, cerr := services.New(handler)

		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusConflict), "data": cerr})
		}
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusText(http.StatusCreated), "data": code})
	}
}

func Modify[T utils.Tables](c *gin.Context, o utils.PersistenceHandler[T], uid uint64) {
	if reflect.TypeOf(o) != nil {
		if err := c.ShouldBindJSON(&o); err != nil {
			fmt.Println("ERRO:", err)
			//msg para deploy
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": "Erro de JSON"})
			//msg para dev
			//c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": err})
			return

		}
		var handler utils.PersistenceHandler[T] = o
		code, cerr := services.Update(handler, uid)

		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": cerr})
		}
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusText(http.StatusAccepted), "data": code})
	}
}

func Erase[T utils.Tables](c *gin.Context, o utils.PersistenceHandler[T], uid uint64) {
	if reflect.TypeOf(o) != nil {
		var handler utils.PersistenceHandler[T] = o
		code, cerr := services.Del(handler, uid)

		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": cerr})
		}
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusText(http.StatusNoContent), "data": code})
	}
}

func Get[T utils.Tables](c *gin.Context, o utils.PersistenceHandler[T], uid uint64) {
	if reflect.TypeOf(o) != nil {
		if uid == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": "{}"})
			return
		}
		var handler utils.PersistenceHandler[T] = o
		rec, cerr := services.Get(handler, uid)

		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": cerr})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": rec, "status": http.StatusText(http.StatusOK)})
	}
}

func GetAll[T utils.Tables](c *gin.Context, o utils.PersistenceHandler[T]) {
	if reflect.TypeOf(o) != nil {
		var handler utils.PersistenceHandler[T] = o
		rec, cerr := services.GetAll(handler)
		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest),
				"data": cerr})
		}
		c.JSON(http.StatusOK, gin.H{"data": rec, "status": http.StatusText(http.StatusOK)})
	}
}

func GetBy[T utils.Tables](c *gin.Context, o utils.PersistenceHandler[T], param string, uid ...interface{}) {
	if reflect.TypeOf(o) != nil {
		if len(uid) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusText(http.StatusNotFound), "data": "No Data"})
		}
		var handler utils.PersistenceHandler[T] = o
		rec, cerr := services.GetBy(handler, param, uid)

		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": cerr})
		}
		c.JSON(http.StatusOK, gin.H{"data": rec, "status": http.StatusText(http.StatusOK)})
	}
}
