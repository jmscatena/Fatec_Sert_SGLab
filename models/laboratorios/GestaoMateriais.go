package laboratorios

import (
	"errors"
	"github.com/jmscatena/Fatec_Sert_SGLab/models/administrativo"
	"gorm.io/gorm"
	"strings"
	"time"
)

type GestaoMateriais struct {
	gorm.Model
	ID         uint64                 `gorm:"primary_key;auto_increment" json:"id"`
	Reserva    Reservas               `gorm:"foreignKey:ID" json:"reserva"`
	Disponivel bool                   `gorm:"default:false" json:"disponivel"`
	CompraEm   time.Time              `json:"compra_em"`
	CreatedBy  administrativo.Usuario `gorm:"foreignKey:ID" json:"created_by"`
	CreatedAt  time.Time              `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (p *GestaoMateriais) Validate() error {
	return nil
}

func (p *GestaoMateriais) Create(db *gorm.DB) (int64, error) {
	if verr := p.Validate(); verr != nil {
		return -1, verr
	}
	err := db.Debug().Model(&GestaoMateriais{}).Create(&p).Error
	if err != nil {
		return 0, err
	}
	return int64(p.ID), nil
}

func (p *GestaoMateriais) Update(db *gorm.DB, uid uint64) (*GestaoMateriais, error) {
	err := db.Debug().Model(&GestaoMateriais{}).Where("id = ?", uid).Take(&GestaoMateriais{}).UpdateColumns(
		map[string]interface{}{
			"reserva":    p.Reserva,
			"disponivel": p.Disponivel,
			"compra_em":  p.CompraEm,
			"criado_por": p.CreatedBy}).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *GestaoMateriais) List(db *gorm.DB) (*[]GestaoMateriais, error) {
	GestaoMateriaiss := []GestaoMateriais{}
	//err := db.Debug().Model(&GestaoMateriais{}).Limit(100).Find(&GestaoMateriaiss).Error
	result := db.Find(&GestaoMateriaiss)
	if result.Error != nil {
		return nil, result.Error
	}
	return &GestaoMateriaiss, nil
}

func (p *GestaoMateriais) Find(db *gorm.DB, uid uint64) (*GestaoMateriais, error) {
	err := db.Debug().Model(&GestaoMateriais{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &GestaoMateriais{}, err
	}
	return p, nil
}

func (p *GestaoMateriais) FindBy(db *gorm.DB, param string, uid ...interface{}) (*[]GestaoMateriais, error) {
	GestaoMateriaiss := []GestaoMateriais{}
	params := strings.Split(param, ";")
	uids := uid[0].([]interface{})
	if len(params) != len(uids) {
		return nil, errors.New("condição inválida")
	}
	result := db.Where(strings.Join(params, " AND "), uids...).Find(&GestaoMateriaiss)
	if result.Error != nil {
		return nil, result.Error
	}
	return &GestaoMateriaiss, nil
}

func (p *GestaoMateriais) Delete(db *gorm.DB, uid uint64) (int64, error) {
	db = db.Delete(&GestaoMateriais{}, "id = ? ", uid)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (p *GestaoMateriais) DeleteBy(db *gorm.DB, cond string, uid uint64) (int64, error) {
	result := db.Delete(&GestaoMateriais{}, cond+" = ?", uid)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
