package migrations

import (
	"fmt"
	admin "github.com/jmscatena/Fatec_Sert_SGLab/dto/models/administrativo"
	lab "github.com/jmscatena/Fatec_Sert_SGLab/dto/models/laboratorios"
	"gorm.io/gorm"
)

func RunMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&admin.Usuario{}, &lab.Materiais{}, &lab.Laboratorios{}, &lab.Reservas{}, &lab.GestaoMateriais{})
	if err != nil {
		fmt.Println("Migrating database erro:", err)
		return
	}

}
