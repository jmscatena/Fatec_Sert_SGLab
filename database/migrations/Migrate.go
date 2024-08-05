package migrations

import (
	"fmt"
	"github.com/jmscatena/Fatec_Sert_SGLab/models/administrativo"
	"github.com/jmscatena/Fatec_Sert_SGLab/models/laboratorios"
	"gorm.io/gorm"
)

func RunMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&administrativo.Usuario{}, &laboratorios.Materiais{}, &laboratorios.Laboratorios{},
		&laboratorios.Reservas{}, &laboratorios.GestaoMateriais{})
	if err != nil {
		return
	}
	fmt.Println("Migrating database erro:", err)
}
