package interfaces

import (
	"github.com/jmscatena/Fatec_Sert_SGLab/models"
	"gorm.io/gorm"
)

type Tables interface {
	models.Usuario | models.Instituicao | models.Certificado | models.CertVal | models.Evento
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
