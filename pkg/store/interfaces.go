package store

import "desafioFinalGo/internal/domain"

type StoreInterface interface {
	ReadOdontologo(id int) (*domain.Odontologo, error)
	CreateOdontologo(odontologo domain.Odontologo) error
	UpdateOdontologo(id int, odontologo domain.Odontologo) error
	PatchOdontologo(id int, odontologo domain.Odontologo) error
	DeleteOdontologo(id int) error
}
