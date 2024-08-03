package main

import (
	"github.com/jmscatena/Fatec_Sert_SGLab/database"
	"github.com/jmscatena/Fatec_Sert_SGLab/server"
)

func main() {
	database.Init()
	r := server.NewServer()
	r.Run()
}
