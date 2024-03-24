package store

import (
	"database/sql"
	"desafioFinalGo/internal/domain"
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

	query := "SELECT * FROM odontologos WHERE id = ?;"

	row := s.db.QueryRow(query, id)
	err := row.Scan(&odontologo.Id, &odontologo.Nombre, &odontologo.Apellido, &odontologo.Matricula)
	if err != nil {
		return nil, err
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
