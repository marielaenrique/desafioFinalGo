package turno

import (
	"desafioFinalGo/internal/domain"
	"desafioFinalGo/pkg/store"
	"errors"
)

type IRepository interface {
	GetByID(id int) (*domain.Turno, error)
	Create(turno domain.Turno) (domain.Turno, error)
	Update(id int, turno domain.Turno) (domain.Turno, error)
	Patch(id int, turno domain.Turno) (domain.Turno, error)
	Delete(id int) error
}

type Repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) IRepository {
	return &Repository{storage}
}

func (r *Repository) GetByID(id int) (*domain.Turno, error) {
	turno, err := r.storage.ReadTurno(id)
	if err != nil {
		return nil, errors.New("Turno no encontrado")
	}
	return turno, nil

}

func (r *Repository) Create(turno domain.Turno) (domain.Turno, error) {
	t, err := r.storage.CreateTurno(turno)
	if err != nil {
		return domain.Turno{}, errors.New("Error al crear un nuevo turno")
	}
	return t, nil
}

func (r *Repository) Update(id int, turno domain.Turno) (domain.Turno, error) {
	t, err := r.storage.UpdateTurno(id, turno)
	if err != nil {
		return domain.Turno{}, errors.New("Error actualizando turno")
	}
	return t, nil
}

func (r *Repository) Patch(id int, turno domain.Turno) (domain.Turno, error) {
	t, err := r.storage.PatchTurno(id, turno)
	if err != nil {
		return domain.Turno{}, errors.New("Error actualizando turno")
	}
	return t, nil
}

func (r *Repository) Delete(id int) error {
	err := r.storage.DeleteTurno(id)
	if err != nil {
		return err
	}
	return nil
}
