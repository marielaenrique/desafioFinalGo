package store

import (
	"desafioFinalGo/internal/domain"
	"encoding/json"
	"log"
)

func (s *SqlStore) ReadTurno(id int) (*domain.Turno, error) {
	var turno domain.Turno
	query := "SELECT * FROM turnos WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&turno.Id, &turno.Descripcion, &turno.FechaHora, &turno.PacienteId, &turno.OdontologoId)
	if err != nil {
		return nil, err
	}
	return &turno, nil
}

func (s *SqlStore) CreateTurno(turno domain.Turno) (domain.Turno, error) {

	query1 := "INSERT INTO turnos (descripcion, fechaHora, pacienteId, odontologoId) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query1)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(turno.Descripcion, turno.FechaHora, turno.PacienteId, turno.OdontologoId)
	if err != nil {
		log.Fatal(err)
	}

	turnoID, err := res.LastInsertId()
	if err != nil {
		return domain.Turno{}, err
	}

	turno.Id = int(turnoID)

	p, err := s.ReadPaciente(turno.PacienteId)
	if err != nil {
		log.Fatal(err)
	}

	p.Turnos = append(p.Turnos, turno)

	turnosJSON1, err := json.Marshal(p.Turnos)
	if err != nil {
		log.Fatal(err)
	}

	query2 := "UPDATE pacientes SET turnos = ? WHERE id = ?"
	_, err = s.db.Exec(query2, string(turnosJSON1), turno.PacienteId)
	if err != nil {
		log.Fatal(err)
	}

	o, err := s.ReadOdontologo(turno.OdontologoId)
	if err != nil {
		log.Fatal(err)
	}

	o.Turnos = append(o.Turnos, turno)

	turnosJSON2, err := json.Marshal(o.Turnos)
	if err != nil {
		log.Fatal(err)
	}

	query3 := "UPDATE odontologos SET turnos = ? WHERE id = ?"
	_, err = s.db.Exec(query3, string(turnosJSON2), turno.OdontologoId)
	if err != nil {
		log.Fatal(err)
	}

	return turno, nil
}

func (s *SqlStore) UpdateTurno(id int, turno domain.Turno) (domain.Turno, error) {

	query1 := "UPDATE turnos SET descripcion = ?, fechaHora = ?, pacienteId = ?, odontologoId = ? WHERE id = ?"
	stmt, err := s.db.Prepare(query1)
	if err != nil {
		return domain.Turno{}, err
	}
	_, err = stmt.Exec(turno.Descripcion, turno.FechaHora, turno.PacienteId, turno.OdontologoId, id)
	if err != nil {
		return domain.Turno{}, err
	}

	turno.Id = id

	err = s.updateTurnoPaciente(turno.PacienteId, turno)
	if err != nil {
		return domain.Turno{}, err
	}

	err = s.updateTurnoOdontologo(turno.OdontologoId, turno)
	if err != nil {
		return domain.Turno{}, err
	}

	return turno, nil
}

func (s *SqlStore) PatchTurno(id int, turno domain.Turno) (domain.Turno, error) {

	var t domain.Turno

	query1 := "SELECT id, descripcion, fechaHora, pacienteId, odontologoId FROM turnos WHERE id = ?;"

	row := s.db.QueryRow(query1, id)
	err := row.Scan(&t.Id, &t.Descripcion, &t.FechaHora, &t.PacienteId, &t.OdontologoId)
	if err != nil {
		return domain.Turno{}, err
	}

	if turno.Descripcion != "" {
		t.Descripcion = turno.Descripcion
	}
	if turno.FechaHora != "" {
		t.FechaHora = turno.FechaHora
	}
	if turno.PacienteId != 0 {
		t.PacienteId = turno.PacienteId
	}
	if turno.OdontologoId != 0 {
		t.OdontologoId = turno.OdontologoId
	}

	query2 := "UPDATE turnos SET descripcion = ?, fechaHora = ?, pacienteId = ?, odontologoId = ? WHERE id = ?"

	stmt, err := s.db.Prepare(query2)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(t.Descripcion, t.FechaHora, t.PacienteId, t.OdontologoId, id)
	if err != nil {
		log.Fatal(err)
	}

	turno.Id = id

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Turno{}, err
	}

	err = s.updateTurnoPaciente(t.PacienteId, t)
	if err != nil {
		return domain.Turno{}, err
	}

	err = s.updateTurnoOdontologo(t.OdontologoId, t)
	if err != nil {
		return domain.Turno{}, err
	}

	return t, nil

}

func (s *SqlStore) updateTurnoPaciente(pacienteID int, turno domain.Turno) error {

	paciente, err := s.ReadPaciente(pacienteID)
	if err != nil {
		return err
	}

	for i, t := range paciente.Turnos {
		if t.Id == turno.Id {
			paciente.Turnos = append(paciente.Turnos[:i], paciente.Turnos[i+1:]...)
			break
		}
	}

	paciente.Turnos = append(paciente.Turnos,
		domain.Turno{
			Id:           turno.Id,
			Descripcion:  turno.Descripcion,
			FechaHora:    turno.FechaHora,
			PacienteId:   turno.PacienteId,
			OdontologoId: turno.OdontologoId,
		})

	turnosJSON, err := json.Marshal(paciente.Turnos)
	if err != nil {
		return err
	}

	query := "UPDATE pacientes SET turnos = ? WHERE id = ?"
	_, err = s.db.Exec(query, string(turnosJSON), pacienteID)
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlStore) updateTurnoOdontologo(odontologoID int, turno domain.Turno) error {

	odontologo, err := s.ReadOdontologo(odontologoID)
	if err != nil {
		return err
	}

	for i, t := range odontologo.Turnos {
		if t.Id == turno.Id {
			odontologo.Turnos = append(odontologo.Turnos[:i], odontologo.Turnos[i+1:]...)
			break
		}
	}

	odontologo.Turnos = append(odontologo.Turnos,
		domain.Turno{
			Id:           turno.Id,
			Descripcion:  turno.Descripcion,
			FechaHora:    turno.FechaHora,
			PacienteId:   turno.PacienteId,
			OdontologoId: turno.OdontologoId,
		})

	turnosJSON, err := json.Marshal(odontologo.Turnos)
	if err != nil {
		return err
	}

	query := "UPDATE odontologos SET turnos = ? WHERE id = ?"
	_, err = s.db.Exec(query, string(turnosJSON), odontologoID)
	if err != nil {
		return err
	}

	return nil
}
