package controllers

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jmscatena/Fatec_Sert_SGLab/interfaces"
	"github.com/jmscatena/Fatec_Sert_SGLab/services"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "acesso " + http.StatusText(200)})

}

func Add[T interfaces.Tables](c *gin.Context, o interfaces.PersistenceHandler[T]) {
	if reflect.TypeOf(o) != nil {
		if err := c.ShouldBindJSON(&o); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": err})
		}
		var handler interfaces.PersistenceHandler[T] = o
		code, cerr := services.New(handler)

		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": cerr})
		}
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusText(http.StatusCreated), "data": code})
	}
}

func Modify[T interfaces.Tables](c *gin.Context, o interfaces.PersistenceHandler[T], uid uint64) {
	if reflect.TypeOf(o) != nil {
		if err := c.ShouldBindJSON(&o); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": err})
		}
		var handler interfaces.PersistenceHandler[T] = o
		code, cerr := services.Update(handler, uid)

		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": cerr})
		}
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusText(http.StatusCreated), "data": code})
	}
}

func Erase[T interfaces.Tables](c *gin.Context, o interfaces.PersistenceHandler[T], uid uint64) {
	if reflect.TypeOf(o) != nil {
		var handler interfaces.PersistenceHandler[T] = o
		code, cerr := services.Del(handler, uid)

		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": cerr})
		}
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusText(http.StatusCreated), "data": code})
	}
}

func Get[T interfaces.Tables](c *gin.Context, o interfaces.PersistenceHandler[T], uid uint64) {
	if reflect.TypeOf(o) != nil {
		if uid == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusText(http.StatusNotFound), "data": "No Data"})
		}
		var handler interfaces.PersistenceHandler[T] = o
		rec, cerr := services.Get(handler, uid)

		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": cerr})
		}
		c.JSON(http.StatusOK, gin.H{"data": rec, "status": http.StatusText(http.StatusOK)})
	}
}

func GetAll[T interfaces.Tables](c *gin.Context, o interfaces.PersistenceHandler[T]) {
	if reflect.TypeOf(o) != nil {
		var handler interfaces.PersistenceHandler[T] = o
		rec, cerr := services.GetAll(handler)
		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest),
				"data": cerr})
		}
		c.JSON(http.StatusOK, gin.H{"data": rec, "status": http.StatusText(http.StatusOK)})
	}
}

func GetBy[T interfaces.Tables](c *gin.Context, o interfaces.PersistenceHandler[T], param string, uid ...interface{}) {
	if reflect.TypeOf(o) != nil {
		if len(uid) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusText(http.StatusNotFound), "data": "No Data"})
		}
		var handler interfaces.PersistenceHandler[T] = o
		rec, cerr := services.GetBy(handler, param, uid)

		if cerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusText(http.StatusBadRequest), "data": cerr})
		}
		c.JSON(http.StatusOK, gin.H{"data": rec, "status": http.StatusText(http.StatusOK)})
	}
}
