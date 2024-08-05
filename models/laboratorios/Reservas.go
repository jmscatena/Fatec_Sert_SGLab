package laboratorios

import (
	"errors"
	"github.com/jmscatena/Fatec_Sert_SGLab/models/administrativo"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Reservas struct {
	gorm.Model
	ID           uint64                 `gorm:"primary_key;auto_increment" json:"id"`
	Laboratorio  Laboratorios           `gorm:"foreignKey: ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"laboratorio"`
	DataInicial  time.Time              `gorm:"default:CURRENT_TIMESTAMP" json:"data_inicio"`
	DataFinal    time.Time              `gorm:"default:CURRENT_TIMESTAMP" json:"data_fim"`
	HoraInicial  time.Time              `gorm:"default:CURRENT_TIMESTAMP" json:"hora_inicio"`
	HoraFinal    time.Time              `gorm:"default:CURRENT_TIMESTAMP" json:"hora_fim"`
	DiaSemana    string                 `gorm:"not null; default=0.0" json:"dia_semana"`
	Rotativo     bool                   `gorm:"default:false" json:"rotativo"`
	Autorizado   bool                   `gorm:"not null" json:"autorizado"`
	AutorizadoBy administrativo.Usuario `gorm:"foreignKey: ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"autorizado_por"`
	AutorizadoAt time.Time              `gorm:"default:CURRENT_TIMESTAMP" json:"autorizado_em"`
	SolicitadoBy administrativo.Usuario `gorm:"foreignKey: ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"solicitado_por"`
	SolicitadoAt time.Time              `gorm:"default:CURRENT_TIMESTAMP" json:"solicitado_em"`
	Ativa        bool                   `gorm:"default:false" json:"ativa"`
	CreatedAt    time.Time              `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (p *Reservas) Validate() error {

	if p.DiaSemana == "" || p.DiaSemana == "null" {
		return errors.New("obrigatório: dia da semana")
	}
	if p.DataInicial.IsZero() {
		return errors.New("obrigatório: data inicial")
	}
	if p.DataFinal.IsZero() {
		return errors.New("obrigatório: data final")
	}
	if p.HoraInicial.IsZero() {
		return errors.New("obrigatório: hora inicial")
	}
	if p.HoraFinal.IsZero() {
		return errors.New("obrigatório: hora final")
	}
	return nil
}

func (p *Reservas) Create(db *gorm.DB) (int64, error) {
	if verr := p.Validate(); verr != nil {
		return -1, verr
	}
	err := db.Debug().Model(&Reservas{}).Create(&p).Error
	if err != nil {
		return 0, err
	}
	return int64(p.ID), nil
}

func (p *Reservas) Update(db *gorm.DB, uid uint64) (*Reservas, error) {
	err := db.Debug().Model(&Reservas{}).Where("id = ?", uid).Take(&Reservas{}).UpdateColumns(
		map[string]interface{}{
			"Laboratorio":  p.Laboratorio,
			"DataInicial":  p.DataInicial,
			"DataFinal":    p.DataFinal,
			"HoraInicial":  p.HoraInicial,
			"HoraFinal":    p.HoraFinal,
			"DiaSemana":    p.DiaSemana,
			"Rotativo":     p.Rotativo,
			"Autorizado":   p.Autorizado,
			"AutorizadoBy": p.AutorizadoBy,
			"SolicitadoBy": p.SolicitadoBy,
			"Ativa":        p.Ativa,
			"UpdatedAt":    time.Now()}).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Reservas) List(db *gorm.DB) (*[]Reservas, error) {
	Reservass := []Reservas{}
	//err := db.Debug().Model(&Reservas{}).Limit(100).Find(&Reservass).Error
	result := db.Find(&Reservass)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Reservass, nil
}

func (p *Reservas) Find(db *gorm.DB, uid uint64) (*Reservas, error) {
	err := db.Debug().Model(&Reservas{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Reservas{}, err
	}
	return p, nil
}

func (p *Reservas) FindBy(db *gorm.DB, param string, uid ...interface{}) (*[]Reservas, error) {
	Reservass := []Reservas{}
	params := strings.Split(param, ";")
	uids := uid[0].([]interface{})
	if len(params) != len(uids) {
		return nil, errors.New("condição inválida")
	}
	result := db.Where(strings.Join(params, " AND "), uids...).Find(&Reservass)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Reservass, nil
}

func (p *Reservas) Delete(db *gorm.DB, uid uint64) (int64, error) {
	db = db.Delete(&Reservas{}, "id = ? ", uid)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (p *Reservas) DeleteBy(db *gorm.DB, cond string, uid uint64) (int64, error) {
	result := db.Delete(&Reservas{}, cond+" = ?", uid)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
