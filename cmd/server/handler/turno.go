package handler

import (
	"desafioFinalGo/internal/domain"
	"desafioFinalGo/internal/turno"
	"desafioFinalGo/pkg/web"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TurnoHandler struct {
	s turno.IService
}

func NewTurnoHandler(s turno.IService) *TurnoHandler {
	return &TurnoHandler{
		s: s,
	}
}

func (h *TurnoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("Id inválido"))
			return
		}
		turno, err := h.s.GetByID(id)
		if err != nil {
			web.FailureResponse(c, 404, errors.New("Turno no encontrado"))
			return
		}
		web.SuccessResponse(c, 200, &turno)
	}
}

func (h *TurnoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.Turno

		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("JSON inválido"))
			return
		}

		t, err := h.s.Create(turno)
		if err != nil {
			web.FailureResponse(c, 400, err)
			return
		}
		web.SuccessResponse(c, 200, t)
	}

}

func (h *TurnoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("Id inválido"))
			return
		}
		var turno domain.Turno
		err = c.ShouldBindJSON(&turno)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("JSON inválido"))
			return
		}
		t, err := h.s.Update(id, turno)
		if err != nil {
			web.FailureResponse(c, 400, err)
			return
		}
		web.SuccessResponse(c, 200, t)
	}
}

func (h *TurnoHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Descripcion  string `json:"descripcion,omitempty"`
		FechaHora    string `json:"fechaHora,omitempty"`
		PacienteId   int    `json:"pacienteId,omitempty"`
		OdontologoId int    `json:"odontologoId,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("Id inválido"))
			return
		}
		err = c.ShouldBindJSON(&r)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("JSON inválido"))
			return
		}
		turnoUpdated := domain.Turno{
			Descripcion:  r.Descripcion,
			FechaHora:    r.FechaHora,
			PacienteId:   r.PacienteId,
			OdontologoId: r.OdontologoId,
		}
		t, err := h.s.Patch(id, turnoUpdated)
		if err != nil {
			web.FailureResponse(c, 400, err)
			return
		}
		web.SuccessResponse(c, 200, t)
	}
}

func (h *TurnoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("Id inválido"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.FailureResponse(c, 400, err)
			return
		}
		web.SuccessResponse(c, 204, nil)
	}
}

func (h *TurnoHandler) PostDniMatricula() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.Turno

		dniPaciente, err := strconv.Atoi(c.Query("dni"))
		if err != nil {
			web.FailureResponse(c, 400, errors.New("Dni debe ser un número entero"))
			return
		}

		matriculaOdontologo, err := strconv.Atoi(c.Query("matricula"))
		if err != nil {
			web.FailureResponse(c, 400, errors.New("Matrícula debe ser un número entero"))
			return
		}

		err = c.ShouldBindJSON(&turno)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("JSON inválido"))
			return
		}

		t, err := h.s.CreateTurnoDniMatricula(turno, dniPaciente, matriculaOdontologo)
		if err != nil {
			web.FailureResponse(c, 400, err)
			return
		}
		web.SuccessResponse(c, 200, t)
	}
}

func (h *TurnoHandler) GetByDni() gin.HandlerFunc {
	return func(c *gin.Context) {

		dniPaciente, err := strconv.Atoi(c.Query("dni"))
		if err != nil {
			web.FailureResponse(c, 400, errors.New("Dni debe ser un número entero"))
			return
		}

		turno, err := h.s.GetTurnoByDni(dniPaciente)
		if err != nil {
			web.FailureResponse(c, 404, errors.New("Turno no encontrado"))
			return
		}
		web.SuccessResponse(c, 200, &turno)
	}
}
