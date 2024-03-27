package domain

type Turno struct {
	Id           int    `json:"id"`
	Descripcion  string `json:"descripcion" binding:"required"`
	FechaHora    string `json:"fechaHora" binding:"required"`
	PacienteId   int    `json:"pacienteId"`
	OdontologoId int    `json:"odontologoId"`
}
