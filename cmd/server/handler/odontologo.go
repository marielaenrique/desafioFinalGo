package handler

import (
	"desafioFinalGo/internal/domain"
	"desafioFinalGo/internal/odontologo"
	"desafioFinalGo/pkg/web"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OdontologoHandler struct {
	s odontologo.IService
}

func NewOdontologoHandler(s odontologo.IService) *OdontologoHandler {
	return &OdontologoHandler{
		s: s,
	}
}

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
