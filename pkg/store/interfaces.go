package store

import "desafioFinalGo/internal/domain"

type StoreInterface interface {
	ReadOdontologo(id int) (*domain.Odontologo, error)
	CreateOdontologo(odontologo domain.Odontologo) error
	UpdateOdontologo(id int, odontologo domain.Odontologo) error
	PatchOdontologo(id int, odontologo domain.Odontologo) error
	DeleteOdontologo(id int) error
	ReadPaciente(id int) (*domain.Paciente, error)
	CreatePaciente(paciente domain.Paciente) error
	UpdatePaciente(id int, paciente domain.Paciente) error
	PatchPaciente(id int, paciente domain.Paciente) error
	DeletePaciente(id int) error
}
