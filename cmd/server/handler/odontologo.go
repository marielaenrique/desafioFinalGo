package handler

import (
	"desafioFinalGo/internal/domain"
	"desafioFinalGo/internal/odontologo"
	"desafioFinalGo/pkg/web"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Nombre    string `json:"nombre,omitempty"`
	Apellido  string `json:"apellido,omitempty"`
	Matricula int    `json:"matricula,omitempty"`
}

type OdontologoHandler struct {
	s odontologo.IService
}

func NewOdontologoHandler(s odontologo.IService) *OdontologoHandler {
	return &OdontologoHandler{
		s: s,
	}
}

// @Summary Obtener odontólogo por id
// @Description Obtiene un odontólogo por su id
// @Tags Odontologos
// @Accept json
// @Produce json
// @Param id path int true "ID del odontólogo"
// @Success 200 {object} domain.Odontologo
// @Failure 400 {string} string "ID inválido"
// @Failure 404 {string} string "Odontólogo no encontrado"
// @Router /odontologos/{id} [get]
func (h *OdontologoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("Id inválido"))
			return
		}
		odontologo, err := h.s.GetByID(id)
		if err != nil {
			web.FailureResponse(c, 404, errors.New("Odóntologo no encontrado"))
			return
		}
		web.SuccessResponse(c, 200, &odontologo)
	}
}

// @Summary Obtener odontólogo por id
// @Description Obtiene un odontólogo por su id
// @Tags Odontologos
// @Accept json
// @Produce json
// @Param id path int true "ID del odontólogo"
// @Success 200 {object} domain.Odontologo
// @Failure 400 {string} string "Json invalido"
// @Failure 404 {string} string "Json invalido"
// @Router /odontologos/{id} [post]
func (h *OdontologoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var odontologo domain.Odontologo
		err := c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("JSON inválido"))
			return
		}
		o, err := h.s.Create(odontologo)
		if err != nil {
			web.FailureResponse(c, 400, err)
			return
		}
		web.SuccessResponse(c, 200, o)
	}
}

// Put actualiza un odontólogo por su ID.
// @Summary Actualizar odontólogo por ID
// @Description Actualiza un odontólogo existente por su ID.
// @Tags Odontologos
// @Accept json
// @Produce json
// @Param id path int true "ID del odontólogo a actualizar"
// @Param odontologo body domain.Odontologo true "Datos del odontólogo a actualizar"
// @Success 200 {object} domain.Odontologo "Odontólogo actualizado"
// @Failure 400 {string} string "ID inválido o JSON inválido"
// @Failure 404 {string} string "Odontólogo no encontrado"
// @Router /odontologos/{id} [put]
func (h *OdontologoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("Id inválido"))
			return
		}
		var odontologo domain.Odontologo
		err = c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.FailureResponse(c, 400, errors.New("JSON inválido"))
			return
		}
		o, err := h.s.Update(id, odontologo)
		if err != nil {
			web.FailureResponse(c, 400, err)
			return
		}
		web.SuccessResponse(c, 200, o)
	}
}

// Patch actualiza un odontólogo por su ID.
// @Summary Actualizar odontólogo por ID
// @Description Actualiza un odontólogo existente por su ID.
// @Tags Odontologos
// @Accept json
// @Produce json
// @Param id path int true "ID del odontólogo a actualizar"
// @Param request body Request true "Datos del odontólogo a actualizar (campos opcionales)"
// @Success 200 {object} domain.Odontologo "Odontólogo actualizado"
// @Failure 400 {string} string "ID inválido o JSON inválido"
// @Failure 404 {string} string "Odontólogo no encontrado"
// @Router /odontologos/{id} [Patch]
func (h *OdontologoHandler) Patch() gin.HandlerFunc {

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
		odontologoUpdated := domain.Odontologo{
			Nombre:    r.Nombre,
			Apellido:  r.Apellido,
			Matricula: r.Matricula,
		}
		o, err := h.s.Patch(id, odontologoUpdated)
		if err != nil {
			web.FailureResponse(c, 400, err)
			return
		}
		web.SuccessResponse(c, 200, o)
	}
}

// Delete elimina un odontólogo por su ID.
// @Summary Eliminar odontólogo por ID
// @Description Elimina un odontólogo existente por su ID.
// @Tags Odontologos
// @Accept json
// @Produce json
// @Param id path int true "ID del odontólogo a eliminar"
// @Success 204 {string} string "Odontólogo eliminado"
// @Failure 400 {string} string "ID inválido"
// @Failure 404 {string} string "Odontólogo no encontrado"
// @Router /odontologos/{id} [delete]
func (h *OdontologoHandler) Delete() gin.HandlerFunc {
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
