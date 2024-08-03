package migrations

import (
	"github.com/jmscatena/Fatec_Sert_SGLab/models/administrativo"
	"github.com/jmscatena/Fatec_Sert_SGLab/models/laboratorios"
	"gorm.io/gorm"
)

func RunMigrate(db *gorm.DB) {

	db.AutoMigrate(&administrativo.GestaoMateriais{})
	db.AutoMigrate(&administrativo.Usuario{})
	db.AutoMigrate(&laboratorios.Laboratorios{})
	db.AutoMigrate(&laboratorios.Reservas{})
	db.AutoMigrate(&laboratorios.Materiais{})
}
