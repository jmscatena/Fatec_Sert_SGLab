package laboratorios

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jmscatena/Fatec_Sert_SGLab/database/models/administrativo"
	"gorm.io/gorm"
	"strings"
	"time"
)

type GestaoMateriais struct {
	// Esta faltando os materiais
	gorm.Model
	UID        uuid.UUID              `gorm:"type:uuid;default:uuid_generate_v4()" json:"ID"`
	ReservaID  uint64                 `json:"-"`
	Reserva    Reservas               `gorm:"references:ID" json:"reserva"`
	Disponivel bool                   `gorm:"default:false" json:"disponivel"`
	CompraEm   time.Time              `json:"compra_em"`
	UsuarioID  uint64                 `json:"-"`
	CreatedBy  administrativo.Usuario `gorm:"foreignKey:UsuarioID" json:"created_by"`
}

func (p *GestaoMateriais) Validate() error {
	return nil
}

func (p *GestaoMateriais) Create(db *gorm.DB) (uuid.UUID, error) {
	if verr := p.Validate(); verr != nil {
		return uuid.Nil, verr
	}
	err := db.Debug().Omit("ID").Create(&p).Error
	if err != nil {
		return uuid.Nil, err
	}
	return p.UID, nil
}

func (p *GestaoMateriais) Update(db *gorm.DB, uid uuid.UUID) (*GestaoMateriais, error) {
	db = db.Debug().Model(GestaoMateriais{}).Where("id = ?", uid).Updates(GestaoMateriais{
		Reserva:    p.Reserva,
		Disponivel: p.Disponivel,
		CompraEm:   p.CompraEm,
		CreatedBy:  p.CreatedBy})

	if db.Error != nil {
		return nil, db.Error
	}
	return p, nil
}

func (p *GestaoMateriais) List(db *gorm.DB) (*[]GestaoMateriais, error) {
	GestaoMateriaiss := []GestaoMateriais{}
	err := db.Debug().Model(&GestaoMateriais{}).Limit(100).Find(&GestaoMateriaiss).Error
	//result := db.Find(&GestaoMateriaiss)
	if err != nil {
		return nil, err
	}
	return &GestaoMateriaiss, nil
}

func (p *GestaoMateriais) Find(db *gorm.DB, uid uuid.UUID) (*GestaoMateriais, error) {
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

func (p *GestaoMateriais) Delete(db *gorm.DB, uid uuid.UUID) (int64, error) {
	db = db.Delete(&GestaoMateriais{}, "id = ? ", uid)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (p *GestaoMateriais) DeleteBy(db *gorm.DB, cond string, uid uuid.UUID) (int64, error) {
	result := db.Delete(&GestaoMateriais{}, cond+" = ?", uid)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
