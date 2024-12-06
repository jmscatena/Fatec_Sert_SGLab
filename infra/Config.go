package infra

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/jmscatena/Fatec_Sert_SGLab/dto/migrations"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Connection struct {
	Db    *gorm.DB
	NoSql *redis.Client
}

func (c *Connection) InitDB() (*gorm.DB, error) {
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

	c.Db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln("Erro no carregamento do SGBD", err)
	}
	migrations.RunMigrate(c.Db)
	return c.Db, err
}

func (c *Connection) InitNoSQL() (*redis.Client, error) {
	var client *redis.Client
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	c.NoSql = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	var _, err = client.Ping().Result()
	if err != nil {
		log.Fatalln("Erro no carregamento do Redis:", err)
	}
	return c.NoSql, nil
}

type Server struct {
	Port   string
	Server *gin.Engine
}

func (s *Server) NewServer(port string) {
	s.Port = port
	s.Server = gin.Default()
}

func (s *Server) Run() {
	log.Printf("Server running at port: %v", s.Port)

}
