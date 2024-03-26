package store

import (
	"database/sql"
	"desafioFinalGo/internal/domain"
	"encoding/json"
	"log"
)

type SqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &SqlStore{
		db: db,
	}
}

func (s *SqlStore) ReadOdontologo(id int) (*domain.Odontologo, error) {
	var odontologo domain.Odontologo
	var turnosData sql.NullString
	query := "SELECT * FROM odontologos WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&odontologo.Id, &odontologo.Nombre, &odontologo.Apellido, &odontologo.Matricula, &odontologo.turnosData)
	if err != nil {
		return nil, err
	}
	if turnosData.Valid {
	if err := json.Unmarshal([]byte(turnosData.String), &odontologo.Turnos); err != nil {
	return nil, err
	}
	}
	return &odontologo, nil
}

func (s *SqlStore) CreateOdontologo(odontologo domain.Odontologo) error {
	query := "INSERT INTO odontologos (nombre, apellido, matricula) VALUES (?, ?, ?);"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(odontologo.Nombre, odontologo.Apellido, odontologo.Matricula)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlStore) UpdateOdontologo(id int, odontologo domain.Odontologo) error {
	query := "UPDATE odontologos SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(odontologo.Nombre, odontologo.Apellido, odontologo.Matricula, id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlStore) PatchOdontologo(id int, odontologo domain.Odontologo) error {

	var o domain.Odontologo

	query1 := "SELECT id, nombre, apellido, matricula FROM odontologos WHERE id = ?;"

	row := s.db.QueryRow(query1, id)
	err := row.Scan(&o.Id, &o.Nombre, &o.Apellido, &o.Matricula)
	if err != nil {
		return err
	}

	if odontologo.Nombre != "" {
		o.Nombre = odontologo.Nombre
	}
	if odontologo.Apellido != "" {
		o.Apellido = odontologo.Apellido
	}
	if odontologo.Matricula != 0 {
		o.Matricula = odontologo.Matricula
	}

	query2 := "UPDATE odontologos SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?;"

	stmt, err := s.db.Prepare(query2)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(o.Nombre, o.Apellido, o.Matricula, id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil

}

func (s SqlStore) DeleteOdontologo(id int) error {

	query := "DELETE FROM odontologos WHERE id = ?;"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil

}

// Pacientes

func (s *SqlStore) ReadPaciente(id int) (*domain.Paciente, error) {
	var paciente domain.Paciente
	var turnosData sql.NullString

	query := "SELECT * FROM pacientes WHERE id = ?;"

	row := s.db.QueryRow(query, id)
	err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.AltaSistema, &turnosData)
	if err != nil {
		return nil, err
	}

	if turnosData.Valid {
		if err := json.Unmarshal([]byte(turnosData.String), &paciente.Turnos); err != nil {
			return nil, err
		}
	}

	return &paciente, nil
}

func (s *SqlStore) CreatePaciente(paciente domain.Paciente) error {
	query := "INSERT INTO pacientes (nombre, apellido, domicilio, dni, altasistema) VALUES (?, ?, ?,?,?);"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.DNI, paciente.AltaSistema)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlStore) UpdatePaciente(id int, paciente domain.Paciente) error {
	query := "UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, altasistema = ? WHERE id = ?"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.DNI, paciente.AltaSistema, id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlStore) PatchPaciente(id int, paciente domain.Paciente) error {

	var p domain.Paciente

	query1 := "SELECT id, nombre, apellido, domicilio, dni, altasistema FROM pacientes WHERE id = ?;"

	row := s.db.QueryRow(query1, id)
	err := row.Scan(&p.Id, &p.Nombre, &p.Apellido, &p.Domicilio, &p.DNI, &p.AltaSistema)
	if err != nil {
		return err
	}

	if paciente.Nombre != "" {
		p.Nombre = paciente.Nombre
	}
	if paciente.Apellido != "" {
		p.Apellido = paciente.Apellido
	}
	if paciente.Domicilio != "" {
		p.Domicilio = paciente.Domicilio
	}
	if paciente.DNI != 0 {
		p.DNI = paciente.DNI
	}
	if paciente.AltaSistema != "" {
		p.AltaSistema = paciente.AltaSistema
	}
	query2 := "UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, altasistema = ? WHERE id = ?"

	stmt, err := s.db.Prepare(query2)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(p.Nombre, p.Apellido, p.Domicilio, p.DNI, p.AltaSistema, id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil

}

func (s SqlStore) DeletePaciente(id int) error {

	query := "DELETE FROM pacientes WHERE id = ?;"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil

}
