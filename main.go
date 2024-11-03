package main

import (
	"github.com/jmscatena/Fatec_Sert_SGLab/database"
)

func main() {
	//http.ListenAndServe(":8000", nil)
	database.Init()
	database.InitDF()
	r := config.NewServer("8000")
	r.Run()
}
