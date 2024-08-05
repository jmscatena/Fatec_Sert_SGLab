package laboratorios

import (
	"errors"
	"github.com/jmscatena/Fatec_Sert_SGLab/models/administrativo"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Laboratorios struct {
	gorm.Model
	ID                  uint64                 `gorm:"primary_key;auto_increment" json:"id"`
	Titulo              string                 `gorm:"not null" json:"titulo"`
	Descricao           string                 `json:"descricao"`
	Quantidade          int16                  `gorm:"not null; default=20" json:"quantidade"`
	ComputadorProfessor bool                   `gorm:"default=true" json:"pc_professor"`
	Rotativo            bool                   `gorm:"default=true" json:"rotativo"`
	CreateUserID        int                    `json:"-"`
	CreatedBy           administrativo.Usuario `gorm:"foreignKey:CreateUserID;references:ID" json:"created_by"`
	CreatedAt           time.Time              `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateUserID        int                    `json:"-"`
	UpdatedBy           administrativo.Usuario `gorm:"foreignKey:UpdateUserID;references:ID" json:"updated_by"`
	UpdatedAt           time.Time              `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Materiais           []*Materiais           `gorm:"many2many:laboratorio_materiais" json:"materiais"`
}

func (p *Laboratorios) Validate() error {

	if p.Titulo == "" || p.Titulo == "null" {
		return errors.New("obrigatório: titulo")
	}
	return nil
}

func (p *Laboratorios) Create(db *gorm.DB) (int64, error) {
	if verr := p.Validate(); verr != nil {
		return -1, verr
	}
	err := db.Debug().Model(&Laboratorios{}).Create(&p).Error
	if err != nil {
		return 0, err
	}
	return int64(p.ID), nil
}

func (p *Laboratorios) Update(db *gorm.DB, uid uint64) (*Laboratorios, error) {
	err := db.Debug().Model(&Laboratorios{}).Where("id = ?", uid).Take(&Laboratorios{}).UpdateColumns(
		map[string]interface{}{
			"Titulo":              p.Titulo,
			"Descricao":           p.Descricao,
			"Quantidade":          p.Quantidade,
			"ComputadorProfessor": p.ComputadorProfessor,
			"Rotativo":            p.Rotativo,
			"Materiais":           p.Materiais,
			"UpdatedBy":           p.UpdatedBy,
			"UpdatedAt":           time.Now()}).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Laboratorios) List(db *gorm.DB) (*[]Laboratorios, error) {
	Laboratorioss := []Laboratorios{}
	//err := db.Debug().Model(&Laboratorios{}).Limit(100).Find(&Laboratorioss).Error
	result := db.Find(&Laboratorioss)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Laboratorioss, nil
}

func (p *Laboratorios) Find(db *gorm.DB, uid uint64) (*Laboratorios, error) {
	err := db.Debug().Model(&Laboratorios{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Laboratorios{}, err
	}
	return p, nil
}

func (p *Laboratorios) FindBy(db *gorm.DB, param string, uid ...interface{}) (*[]Laboratorios, error) {
	Laboratorioss := []Laboratorios{}
	params := strings.Split(param, ";")
	uids := uid[0].([]interface{})
	if len(params) != len(uids) {
		return nil, errors.New("condição inválida")
	}
	result := db.Where(strings.Join(params, " AND "), uids...).Find(&Laboratorioss)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Laboratorioss, nil
}

func (p *Laboratorios) Delete(db *gorm.DB, uid uint64) (int64, error) {
	db = db.Delete(&Laboratorios{}, "id = ? ", uid)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (p *Laboratorios) DeleteBy(db *gorm.DB, cond string, uid uint64) (int64, error) {
	result := db.Delete(&Laboratorios{}, cond+" = ?", uid)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
