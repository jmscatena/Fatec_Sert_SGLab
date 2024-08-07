package tests

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/jmscatena/Fatec_Sert_SGLab/controllers"
	"github.com/jmscatena/Fatec_Sert_SGLab/models/administrativo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsuario(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(responseRecorder)
	var obj administrativo.Usuario
	engine.POST("/", func(context *gin.Context) {
		controllers.Add[administrativo.Usuario](context, &obj)
	})
	requestBody := `{"nome":"Teste de nome","email":"teste@awc.com","senha":"1234"}`

	ctx.Request = httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer([]byte(requestBody)))
	engine.ServeHTTP(responseRecorder, ctx.Request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, requestBody, responseRecorder.Body.String())
}
