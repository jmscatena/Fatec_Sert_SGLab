package config

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/jmscatena/Fatec_Sert_SGLab/dto/migrations"
	"github.com/jmscatena/Fatec_Sert_SGLab/routes"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer(port string) Server {
	return Server{
		port:   port,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}

type CONNECTION struct {
	Conn gorm.DB
}

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {

		log.Fatalf("Error Loading Configuration File")
	}

	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbase := os.Getenv("DB")
	dbServer := os.Getenv("DBSERVER")
	dbPort := os.Getenv("DBPORT")
	dbURL := "postgres://" + dbUser + ":" + dbPass + "@" + dbServer + ":" + dbPort + "/" + dbase

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln("Erro no carregamento do SGBD", err)
	}
	migrations.RunMigrate(db)
	return db, err
}

func InitDF() (*redis.Client, error) {
	var client *redis.Client
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalln("Erro no carregamento do Redis:", err)
	}
	return client, nil
}
