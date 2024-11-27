package laboratorios

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"html"
	"log"
	"strings"
	"time"
)

type Materiais struct {
	gorm.Model
	UID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"ID"`
	Titulo     string    `gorm:"not null" json:"titulo"`
	Quantidade float64   `gorm:"not null; default=0.0" json:"quantidade"`
	Medida     string    `gorm:"not null" json:"medida"`

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
	p.Quantidade = float64(p.Quantidade)
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := p.Validate()
	if err != nil {
		log.Fatalf("Error during validation:%v", err)
	}
}
func (p *Materiais) Create(db *gorm.DB) (uuid.UUID, error) {
	if verr := p.Validate(); verr != nil {
		return uuid.Nil, verr
	}
	p.Prepare()
	err := db.Debug().Omit("ID").Create(&p).Error
	if err != nil {
		return uuid.Nil, err
	}
	return p.UID, nil
}
func (p *Materiais) Update(db *gorm.DB, uid uuid.UUID) (*Materiais, error) {
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
func (p *Materiais) Find(db *gorm.DB, uid uuid.UUID) (*Materiais, error) {
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
func (p *Materiais) Delete(db *gorm.DB, uid uuid.UUID) (int64, error) {
	db = db.Delete(&Materiais{}, "id = ? ", uid)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
func (p *Materiais) DeleteBy(db *gorm.DB, cond string, uid uuid.UUID) (int64, error) {
	result := db.Delete(&Materiais{}, cond+" = ?", uid)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
