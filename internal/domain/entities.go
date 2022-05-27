package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Entity struct {
	ID string `validate:"required,uuid4"`
}

func (e Entity) Validate(v *validator.Validate) error {
	return v.Struct(e)
}

type Movie struct {
	Entity
	Title      string    `validate:"required"`
	ReleasedAt time.Time `validate:"required"`
}

func (m Movie) Validate(v *validator.Validate) error {
	return v.Struct(m)
}
