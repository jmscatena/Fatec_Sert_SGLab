package utils

import (
	"github.com/jmscatena/Fatec_Sert_SGLab/database/models/administrativo"
	laboratorios2 "github.com/jmscatena/Fatec_Sert_SGLab/database/models/laboratorios"
	"gorm.io/gorm"
)

type Tables interface {
	administrativo.Usuario | laboratorios2.GestaoMateriais |
		laboratorios2.Laboratorios | laboratorios2.Reservas | laboratorios2.Materiais
}

type PersistenceHandler[T Tables] interface {
	Create(db *gorm.DB) (int64, error)
	List(db *gorm.DB) (*[]T, error)
	Update(db *gorm.DB, uid uint64) (*T, error)
	Find(db *gorm.DB, uid uint64) (*T, error)
	Delete(db *gorm.DB, uid uint64) (int64, error)

	FindBy(db *gorm.DB, param string, uid ...interface{}) (*[]T, error)
	//DeleteBy(db *gorm.DB, cond string, uid uint64) (int64, error)
}
