package model

import (
	"time"

	"gorm.io/gorm"
)

// WorkerTypes es un slice de WorkerType
type WorkerTypes []WorkerType

// WorkerType es el modelo para el tipo de trabajador (OPERATIVO / ADMINISTRATIVO)
type WorkerType struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(200);not null" json:"name"`
	Observations *string `gorm:"type:varchar(250)" json:"observations"`
	Status       uint    `gorm:"not null;default:1" json:"status"`
	HandledBy    uint    `gorm:"not null" json:"handled_by"`
	Staff        Staff
}

// WorkAreas es un slice de WorkArea
type WorkAreas []WorkArea

// WorkArea es el modelo para área de trabajo
type WorkArea struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(200);not null" json:"name"`
	Observations *string `gorm:"type:varchar(250)" json:"observations"`
	Status       uint    `gorm:"not null;default:1" json:"status"`
	HandledBy    uint    `gorm:"not null" json:"handled_by"`
	Staff        Staff
}

// Locals es un slice de Local
type Locals []Local

// Campus es el modelo para listar las sedes de la organización
type Local struct {
	gorm.Model
	Name            string  `gorm:"type:varchar(200);not null" json:"name"`
	Address         string  `gorm:"type:varchar(250);not null" json:"address"`
	Phone           string  `gorm:"type:varchar(20);not null" json:"phone"`
	Observations    *string `gorm:"type:varchar(250)" json:"observations"`
	Status          uint    `gorm:"not null;default:1" json:"status"`
	HandledBy       uint    `gorm:"not null" json:"handled_by"`
	Staff           Staff
	BusinessByLocal BusinessByLocals
}

// Staffs es un slice de Staff
type Staffs []Staff

// Staff es el modelo para el registro del personal de la organización
type Staff struct {
	gorm.Model
	FirstName     string    `gorm:"type:varchar(100);not null" json:"first_name"`
	SecondName    *string   `gorm:"type:varchar(100)" json:"second_name"`
	LastName1     string    `gorm:"type:varchar(100);not null" json:"last_name1"`
	LastName2     string    `gorm:"type:varchar(100);not null" json:"last_name2"`
	Dni           string    `gorm:"type:varchar(20);not null unique" json:"dni"`
	BirthDate     time.Time `gorm:"type:timestamp;not null" json:"birth_date"`
	Phone         string    `gorm:"type:varchar(20);not null" json:"phone"`
	Email         string    `gorm:"type:varchar(200);not null" json:"email"`
	BusinessID    uint      `json:"business_id"`
	LocalID       uint      `json:"local_id"`
	Status        uint      `gorm:"not null;default:1" json:"status"`
	HandledBy     uint      `gorm:"not null" json:"handled_by"`
	PensionFundID uint
	WorkerTypeID  uint
	WorkAreaID    uint
	UserID        uint
}

// Businesses es un slice de Business
type Businesses []Business

// Business es el modelo para el registro de las razones sociales asociadas a la organización
type Business struct {
	gorm.Model
	Name            string  `gorm:"type:varchar(200);not null" json:"name"`
	Ruc             string  `gorm:"type:varchar(20);not null unique" json:"ruc"`
	Status          uint    `gorm:"not null;default:1" json:"status"`
	Observations    *string `gorm:"type:varchar(250)" json:"observations"`
	HandledBy       uint    `gorm:"not null" json:"handled_by"`
	Staff           Staff
	BusinessByLocal BusinessByLocals
}

// BusinessByLocals es un slice de BusinessByLocal
type BusinessByLocals []BusinessByLocal

// BusinessByLocal es el modelo para el registro de las sedes agrupadas por razon social
type BusinessByLocal struct {
	gorm.Model
	Status     uint `gorm:"not null;default:1" json:"status"`
	HandledBy  uint `gorm:"not null" json:"handled_by"`
	BusinessID uint
	LocalID    uint
}

// Users es un slice de User
type Users []User

// User es el modelo para el login del usuario - Se creará automáticamente al registrar al personal
type User struct {
	gorm.Model
	DniUser   string `gorm:"type:varchar(20);not null" json:"dni_user"`
	Password  string `gorm:"type:varchar(200);not null" json:"password"`
	Status    uint   `gorm:"not null;default:1" json:"status"`
	HandledBy uint   `gorm:"not null" json:"handled_by"`
	Staff     Staff
}

// PensionFunds es un slice de PensionFund
type PensionFunds []PensionFund

// PensionFund es el modelo para el registro de los fondos de pensiones
type PensionFund struct {
	gorm.Model
	Name                 string  `gorm:"type:varchar(200);not null" json:"name"`
	FixedCommission      float64 `gorm:"type:decimal(10,2);not null" json:"fixed_commission"`
	FlowCommission       float64 `gorm:"type:decimal(10,2);not null" json:"flow_commission"`
	MixFlowCommission    float64 `gorm:"type:decimal(10,2);not null" json:"mix_flow_commission"`
	MixBalanceCommission float64 `gorm:"type:decimal(10,2);not null" json:"mix_balance_commission"`
	PremiumInsurance     float64 `gorm:"type:decimal(10,2);not null" json:"premium_insurance"`
	Input                float64 `gorm:"type:decimal(10,2);not null" json:"input"`
	MaxRemuneration      float64 `gorm:"type:decimal(10,2);not null" json:"max_remuneration"`
	Observations         *string `gorm:"type:varchar(250)" json:"observations"`
	Status               uint    `gorm:"not null;default:1" json:"status"`
	HandledBy            uint    `gorm:"not null" json:"handled_by"`
	Staff                Staff
}
