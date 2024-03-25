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

func (s *SqlStore) PatchOdontologo(id int, odontologo domain.Odontologo) error {

	var o domain.Odontologo

	query1 := "SELECT * FROM odontologos WHERE id = ?;"

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
