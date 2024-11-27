package laboratorios

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type Laboratorios struct {
	gorm.Model
	UID                 uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4()" json:"ID"`
	Titulo              string      `gorm:"not null" json:"titulo"`
	Descricao           string      `json:"descricao"`
	Quantidade          int16       `gorm:"not null; default=20" json:"quantidade"`
	ComputadorProfessor bool        `gorm:"default=true" json:"pc_professor"`
	Rotativo            bool        `gorm:"default=false" json:"rotativo"`
	Materiais           []Materiais `gorm:"many2many:laboratorio_materiais" json:"materiais"`
}

func (p *Laboratorios) Validate() error {

	if p.Titulo == "" || p.Titulo == "null" {
		return errors.New("obrigatório: titulo")
	}
	if p.Quantidade == 0 {
		return errors.New("obrigatório: quantidade de computadores")
	}
	return nil
}
func (p *Laboratorios) Prepare(db *gorm.DB) (err error) {
	p.Titulo = html.EscapeString(strings.TrimSpace(p.Titulo))
	p.Descricao = html.EscapeString(strings.TrimSpace(p.Descricao))
	p.Descricao = html.EscapeString(strings.TrimSpace(p.Descricao))
	p.Quantidade = int16(p.Quantidade)
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return
}

func (p *Laboratorios) Create(db *gorm.DB) (uuid.UUID, error) {
	if verr := p.Validate(); verr != nil {
		return uuid.Nil, verr
	}
	p.Prepare(db)
	err := db.Debug().Omit("ID").Create(&p).Error
	if err != nil {
		return uuid.Nil, err
	}
	return p.UID, nil
}

func (p *Laboratorios) Update(db *gorm.DB, uid uuid.UUID) (*Laboratorios, error) {
	p.Prepare(db)
	//err := db.Debug().Model(&Laboratorios{}).Where("id = ?", uid).Take(&Laboratorios{}).UpdateColumns(
	//	map[string]interface{}
	db = db.Model(Laboratorios{}).Where("id = ?", uid).Updates(
		Laboratorios{
			Titulo:              p.Titulo,
			Descricao:           p.Descricao,
			Quantidade:          p.Quantidade,
			ComputadorProfessor: p.ComputadorProfessor,
			Rotativo:            p.Rotativo,
			Materiais:           p.Materiais})
	if db.Error != nil {
		return &Laboratorios{}, db.Error
	}
	return p, nil
}

func (p *Laboratorios) List(db *gorm.DB) (*[]Laboratorios, error) {
	Laboratorioss := []Laboratorios{}
	//err := db.Debug().Model(&Laboratorios{}).Limit(100).Find(&Laboratorioss).Error
	//result := db.Find(&Laboratorioss)
	err := db.Model(&Laboratorios{}).Preload("CreatedBy").Preload("UpdatedBy").Preload("Materiais").Find(&Laboratorioss).Error
	if err != nil {
		return nil, err
	}
	return &Laboratorioss, nil
}

func (p *Laboratorios) Find(db *gorm.DB, uid uuid.UUID) (*Laboratorios, error) {
	err := db.Debug().Model(&Laboratorios{}).Preload("CreatedBy").Preload("UpdatedBy").Preload("Materiais").Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Laboratorios{}, err
	}
	return p, nil
}

func (p *Laboratorios) FindBy(db *gorm.DB, param string, uid ...interface{}) (*[]Laboratorios, error) {
	/*
		Metodo utilizado para pesquisas com outros campos
	*/
	Laboratorioss := []Laboratorios{}
	params := strings.Split(param, ";")
	uids := uid[0].([]interface{})
	if len(params) != len(uids) {
		return nil, errors.New("condição inválida")
	}
	result := db.Model(&Laboratorios{}).Preload("CreatedBy").Preload("UpdatedBy").Preload("Materiais").Where(strings.Join(params, " AND "), uids...).Find(&Laboratorioss)
	//result := db.Joins("CreatedBy", db.Where(strings.Join(params, " AND "), uids...)).Find(&Laboratorioss)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Laboratorioss, nil
}

func (p *Laboratorios) Delete(db *gorm.DB, uid uuid.UUID) (int64, error) {
	db = db.Delete(&Laboratorios{}, "id = ? ", uid)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
