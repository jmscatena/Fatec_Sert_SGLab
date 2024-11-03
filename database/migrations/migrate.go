package migrations

import (
	"fmt"
	"github.com/jmscatena/Fatec_Sert_SGLab/database/models/administrativo"
	laboratorios2 "github.com/jmscatena/Fatec_Sert_SGLab/database/models/laboratorios"
	"gorm.io/gorm"
)

func RunMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&administrativo.Usuario{}, &laboratorios2.Materiais{}, &laboratorios2.Laboratorios{}, &laboratorios2.Reservas{}, &laboratorios2.GestaoMateriais{})
	if err != nil {
		fmt.Println("Migrating database erro:", err)
		return
	}

}
