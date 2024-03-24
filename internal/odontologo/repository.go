package odontologo

import (
	"desafioFinalGo/internal/domain"
	"desafioFinalGo/pkg/store"
	"errors"
)

type IRepository interface {
	GetByID(id int) (*domain.Odontologo, error)
	Create(odontologo domain.Odontologo) (domain.Odontologo, error)
	Update(id int, odontologo domain.Odontologo) (domain.Odontologo, error)
}

type Repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) IRepository {
	return &Repository{storage}
}

func (r *Repository) GetByID(id int) (*domain.Odontologo, error) {
	odontologo, err := r.storage.ReadOdontologo(id)
	if err != nil {
		return nil, errors.New("Odontólogo no encontrado")
	}
	return odontologo, nil
}

func (r *Repository) Create(odontologo domain.Odontologo) (domain.Odontologo, error) {
	err := r.storage.CreateOdontologo(odontologo)
	if err != nil {
		return domain.Odontologo{}, errors.New("Error al crear un nuevo odontólogo")
	}
	return odontologo, nil
}

func (r *Repository) Update(id int, odontologo domain.Odontologo) (domain.Odontologo, error) {
	err := r.storage.UpdateOdontologo(id, odontologo)
	if err != nil {
		return domain.Odontologo{}, errors.New("Error al actualizar odontólogo")
	}
	return odontologo, nil
}
