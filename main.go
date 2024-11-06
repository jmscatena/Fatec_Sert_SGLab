package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmscatena/Fatec_Sert_SGLab/config"
	"github.com/jmscatena/Fatec_Sert_SGLab/database"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	//http.ListenAndServe(":8000", nil)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error Loading Configuration File")
	}
	gin.SetMode(os.Getenv("SET_MODE"))
	_, _ = database.Init()
	_ = database.InitDF()
	r := config.NewServer("8000")
	r.Run()
}
