package domain

type Turno struct {
	Id           int    `json:"id"`
	Descripcion  string `json:"descripcion" binding:"required"`
	FechaHora    string `json:"fechaHora" binding:"required"`
	PacienteId   int    `json:"pacienteId" binding:"required"`
	OdontologoId int    `json:"odontologoId" binding:"required"`
}
