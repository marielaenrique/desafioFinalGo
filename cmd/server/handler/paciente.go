package handler

import (
	"desafioFinalGo/internal/domain"
	"desafioFinalGo/internal/paciente"
	"desafioFinalGo/pkg/web"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PacienteHandler struct {
	s paciente.IService
}

func NewPacienteHandler(s paciente.IService) *PacienteHandler {
	return &PacienteHandler{
		s: s,
	}
}

// @Summary Obtener paciente por id
// @Description Obtiene un paciente por su id
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param id path int true "ID del paciente"
// @Success 200 {object} domain.Paciente
// @Failure 400 {string} string "ID inválido"
// @Failure 404 {string} string "Paciente no encontrado"
// @Router /pacientes/{id} [get]
func (h *PacienteHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("Id inválido"))
			return
		}
		paciente, err := h.s.GetByID(id)
		if err != nil {
			web.FailureResponse(c, 404, errors.New("Paciente no encontrado"))
			return
		}
		web.SuccessResponse(c, 200, &paciente)
	}
}

// @Summary Obtener paciente por id
// @Description Obtiene un paciente por su id
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param id path int true "ID del paciente"
// @Success 200 {object} domain.Paciente
// @Failure 400 {string} string "Json invalido"
// @Failure 404 {string} string "Json invalido"
// @Router /pacientes/{id} [get]
func (h *PacienteHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paciente domain.Paciente
		err := c.ShouldBindJSON(&paciente)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("JSON inválido"))
			return
		}
		p, err := h.s.Create(paciente)
		if err != nil {
			web.FailureResponse(c, 400, err)
			return
		}
		web.SuccessResponse(c, 200, p)
	}
}

// Put actualiza un paciente por su ID.
// @Summary Actualizar paciente por ID
// @Description Actualiza un paciente existente por su ID.
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param id path int true "ID del paciente a actualizar"
// @Param odontologo body domain.Paciente true "Datos del odontólogo a actualizar"
// @Success 200 {object} domain.Paciente "Odontólogo actualizado"
// @Failure 400 {string} string "ID inválido o JSON inválido"
// @Failure 404 {string} string "Paciente no encontrado"
// @Router /pacientes/{id} [put]
func (h *PacienteHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("Id inválido"))
			return
		}
		var paciente domain.Paciente
		err = c.ShouldBindJSON(&paciente)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("JSON inválido"))
			return
		}
		p, err := h.s.Update(id, paciente)
		if err != nil {
			web.FailureResponse(c, 400, err)
			return
		}
		web.SuccessResponse(c, 200, p)
	}
}
// Patch actualiza un paciente por su ID.
// @Summary Actualizar paciente por ID
// @Description Actualiza un paciente existente por su ID.
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param id path int true "ID del paciente a actualizar"
// @Param request body Request true "Datos del paciente a actualizar (campos opcionales)"
// @Success 200 {object} domain.Paciente "Paciente actualizado"
// @Failure 400 {string} string "ID inválido o JSON inválido"
// @Failure 404 {string} string "Paciente no encontrado"
// @Router /pacientes/{id} [Patch]
func (h *PacienteHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Nombre      string `json:"nombre,omitempty"`
		Apellido    string `json:"apellido,omitempty"`
		Domicilio   string `json:"domicilio,omitempty"`
		DNI         int    `json:"dni,omitempty"`
		AltaSistema string `json:"altasistema,omitempty"`
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
		pacienteUpdated := domain.Paciente{
			Nombre:      r.Nombre,
			Apellido:    r.Apellido,
			Domicilio:   r.Domicilio,
			DNI:         r.DNI,
			AltaSistema: r.AltaSistema,
		}
		p, err := h.s.Patch(id, pacienteUpdated)
		if err != nil {
			web.FailureResponse(c, 400, err)
			return
		}
		web.SuccessResponse(c, 200, p)
	}
}

// Delete elimina un paciente por su ID.
// @Summary Eliminar paciente por ID
// @Description Elimina un paciente existente por su ID.
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param id path int true "ID del paciente a eliminar"
// @Success 204 {string} string "Paciente eliminado"
// @Failure 400 {string} string "ID inválido"
// @Failure 404 {string} string "Paciente no encontrado"
// @Router /pacientes/{id} [delete]
func (h *PacienteHandler) Delete() gin.HandlerFunc {
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
