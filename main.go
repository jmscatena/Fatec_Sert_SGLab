package main

import (
	"github.com/jmscatena/Fatec_Sert_SGLab/config"
	"github.com/jmscatena/Fatec_Sert_SGLab/database"
)

func main() {
	//http.ListenAndServe(":8000", nil)
	_, _ = database.Init()
	_ = database.InitDF()
	r := config.NewServer("8000")
	r.Run()
}
