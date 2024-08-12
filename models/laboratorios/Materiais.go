package laboratorios

import (
	"errors"
	"gorm.io/gorm"
	"html"
	"log"
	"strings"
	"time"
)

type Materiais struct {
	gorm.Model
	ID         uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Titulo     string  `gorm:"not null" json:"titulo"`
	Quantidade float64 `gorm:"not null; default=0.0" json:"quantidade"`
	Medida     string  `gorm:"not null" json:"medida"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Materiais) Validate() error {

	if p.Titulo == "" || p.Titulo == "null" {
		return errors.New("obrigatório: titulo")
	}
	if p.Medida == "" || p.Medida == "null" {
		return errors.New("obrigatório: tipo de medida")
	}
	return nil
}

func (p *Materiais) Prepare() {
	p.Titulo = html.EscapeString(strings.TrimSpace(p.Titulo))
	p.Medida = html.EscapeString(strings.TrimSpace(p.Medida))
	p.Quantidade = p.Quantidade
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := p.Validate()
	if err != nil {
		log.Fatalf("Error during validation:%v", err)
	}
}
func (p *Materiais) Create(db *gorm.DB) (int64, error) {
	if verr := p.Validate(); verr != nil {
		return -1, verr
	}
	err := db.Debug().Create(&p).Error
	if err != nil {
		return 0, err
	}
	return int64(p.ID), nil
}

func (p *Materiais) Update(db *gorm.DB, uid uint64) (*Materiais, error) {
	db = db.Debug().Model(&Materiais{}).Where("id = ?", uid).Updates(Materiais{
		Titulo:     p.Titulo,
		Quantidade: p.Quantidade,
		Medida:     p.Medida})
	if db.Error != nil {
		return nil, db.Error
	}
	return p, nil
}

func (p *Materiais) List(db *gorm.DB) (*[]Materiais, error) {
	Materiaiss := []Materiais{}
	err := db.Debug().Model(&Materiais{}).Limit(100).Find(&Materiaiss).Error
	//result := db.Find(&Materiaiss)
	if err != nil {
		return nil, err
	}
	return &Materiaiss, nil
}

func (p *Materiais) Find(db *gorm.DB, uid uint64) (*Materiais, error) {
	err := db.Debug().Model(&Materiais{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Materiais{}, err
	}
	return p, nil
}

func (p *Materiais) FindBy(db *gorm.DB, param string, uid ...interface{}) (*[]Materiais, error) {
	Materiaiss := []Materiais{}
	params := strings.Split(param, ";")
	uids := uid[0].([]interface{})
	if len(params) != len(uids) {
		return nil, errors.New("condição inválida")
	}
	result := db.Where(strings.Join(params, " AND "), uids...).Find(&Materiaiss)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Materiaiss, nil
}

func (p *Materiais) Delete(db *gorm.DB, uid uint64) (int64, error) {
	db = db.Delete(&Materiais{}, "id = ? ", uid)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (p *Materiais) DeleteBy(db *gorm.DB, cond string, uid uint64) (int64, error) {
	result := db.Delete(&Materiais{}, cond+" = ?", uid)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
