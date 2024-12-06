package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmscatena/Fatec_Sert_SGLab/infra"
	"github.com/jmscatena/Fatec_Sert_SGLab/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error Loading Configuration File")
	}
	gin.SetMode(os.Getenv("SET_MODE"))
	DbConn := infra.Connection{Db: nil, NoSql: nil}
	_, err = DbConn.InitDB()
	if err != nil {
		log.Fatalf("Error Loading Database Connection")
	}
	_, err = DbConn.InitNoSQL()
	if err != nil {
		log.Fatalf("Error Loading Redis Connection")
	}
	server := infra.Server{}
	server.NewServer("8000")
	router := routes.ConfigRoutes(server.Server, DbConn)
	log.Fatal(router.Run(":" + server.Port))

}
