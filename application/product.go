package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string `valid:"uuidv4"`
	Name   string `valid:"required"`
	Status string `valid:"required,in(enabled|disabled)"`
	Price  float64 `valid:"float,optional,range(0|9999999999)"`
}

func NewProduct() *Product {
	product := Product{
		ID:     uuid.New().String(),
		Status: DISABLED,
	}

	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	return govalidator.ValidateStruct(p)
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("Price must be greater than 0")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return errors.New("Price must be 0")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}