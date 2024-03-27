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


// @Summary Obtener turno por id
// @Description Obtiene un turno por su id
// @Tags Turnos
// @Accept json
// @Produce json
// @Param id path int true "ID del turno"
// @Success 200 {object} domain.Turno
// @Failure 400 {string} string "ID inválido"
// @Failure 404 {string} string "Turno no encontrado"
// @Router /odontologos/{id} [get]
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

 
// @Summary Obtener turno por id
// @Description Obtiene un turno por su id
// @Tags Turnos
// @Accept json
// @Produce json
// @Param id path int true "ID del turno"
// @Success 200 {object} domain.Turno
// @Failure 400 {string} string "JSON inválido"
// @Failure 404 {string} string "Turno no encontrado"
// @Router /turnos/{id} [get]
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

// Put actualiza un turno por su ID.
// @Summary Actualizar turno por ID
// @Description Actualiza un turno existente por su ID.
// @Tags Turnos
// @Accept json
// @Produce json
// @Param id path int true "ID del turno a actualizar"
// @Param turno body domain.Turno true "Datos del turno a actualizar"
// @Success 200 {object} domain.Turno "Turno actualizado"
// @Failure 400 {string} string "ID inválido o JSON inválido"
// @Failure 404 {string} string "Turno no encontrado"
// @Router /turnos/{id} [put]
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

// Patch actualiza un turno por su ID.
// @Summary Actualizar turno por ID
// @Description Actualiza un turno existente por su ID.
// @Tags Turnos
// @Accept json
// @Produce json
// @Param id path int true "ID del turno a actualizar"
// @Param request body Request true "Datos del turno a actualizar (campos opcionales)"
// @Success 200 {object} domain.Turno "Turno actualizado"
// @Failure 400 {string} string "ID inválido o JSON inválido"
// @Failure 404 {string} string "Turno no encontrado"
// @Router /turnos/{id} [Patch]
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

// Delete elimina un turno por su ID.
// @Summary Eliminar turno por ID
// @Description Elimina un turno existente por su ID.
// @Tags Turnos
// @Accept json
// @Produce json
// @Param id path int true "ID del turno a eliminar"
// @Success 204 {string} string "Turno eliminado"
// @Failure 400 {string} string "ID inválido"
// @Failure 404 {string} string "Turno no encontrado"
// @Router /turnos/{id} [delete]
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
