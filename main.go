package main

import (
	"github.com/jmscatena/Fatec_Sert_SGLab/database"
	"github.com/jmscatena/Fatec_Sert_SGLab/server"
	"net/http"
)

func main() {
	http.ListenAndServe(":8000", nil)
	database.Init()
	r := server.NewServer("8000")
	r.Run()
}
