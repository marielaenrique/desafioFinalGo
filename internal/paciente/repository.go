package paciente

import (
	"desafioFinalGo/internal/domain"
	"desafioFinalGo/pkg/store"
	"errors"
)

type IRepository interface {
	GetByID(id int) (*domain.Paciente, error)
	Create(paciente domain.Paciente) (domain.Paciente, error)
	Update(id int, paciente domain.Paciente) (domain.Paciente, error)
	Patch(id int, paciente domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

type Repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) IRepository {
	return &Repository{storage}
}

func (r *Repository) GetByID(id int) (*domain.Paciente, error) {
	paciente, err := r.storage.ReadPaciente(id)
	if err != nil {
		return nil, errors.New("Paciente no encontrado")
	}
	return paciente, nil
}

func (r *Repository) Create(paciente domain.Paciente) (domain.Paciente, error) {
	err := r.storage.CreatePaciente(paciente)
	if err != nil {
		return domain.Paciente{}, errors.New("Error al crear un nuevo paciente")
	}
	return paciente, nil
}

func (r *Repository) Update(id int, paciente domain.Paciente) (domain.Paciente, error) {
	err := r.storage.UpdatePaciente(id, paciente)
	if err != nil {
		return domain.Paciente{}, errors.New("Error al actualizar paciente")
	}
	return paciente, nil
}

func (r *Repository) Patch(id int, paciente domain.Paciente) (domain.Paciente, error) {
	err := r.storage.PatchPaciente(id, paciente)
	if err != nil {
		return domain.Paciente{}, errors.New("Error al actualizar paciente")
	}
	return paciente, nil
}

func (r *Repository) Delete(id int) error {
	err := r.storage.DeletePaciente(id)
	if err != nil {
		return err
	}
	return nil
}
