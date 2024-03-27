package store

import "desafioFinalGo/internal/domain"

type OdontologoStoreInterface interface {
	ReadOdontologo(id int) (*domain.Odontologo, error)
	CreateOdontologo(odontologo domain.Odontologo) error
	UpdateOdontologo(id int, odontologo domain.Odontologo) error
	PatchOdontologo(id int, odontologo domain.Odontologo) error
	DeleteOdontologo(id int) error
}

type PacienteStoreInterface interface {
	ReadPaciente(id int) (*domain.Paciente, error)
	CreatePaciente(paciente domain.Paciente) error
	UpdatePaciente(id int, paciente domain.Paciente) error
	PatchPaciente(id int, paciente domain.Paciente) error
	DeletePaciente(id int) error
}

type TurnoStoreInterface interface {
	ReadTurno(id int) (*domain.Turno, error)
	CreateTurno(turno domain.Turno) (domain.Turno, error)
	UpdateTurno(id int, turno domain.Turno) (domain.Turno, error)
	PatchTurno(id int, turno domain.Turno) (domain.Turno, error)
	DeleteTurno(id int) error
	CreateTurnoDniMatricula(turno domain.Turno, dniPaciente int, matriculaOdontologo int) (domain.Turno, error)
}
