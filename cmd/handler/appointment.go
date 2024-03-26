package handler

import (
	"strconv"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/appointment"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	s appointment.Service
}

// NewAppointmentHandler create a new appointment controller
func NewAppointmentHandler(s appointment.Service) *appointmentHandler {
	return &appointmentHandler{
		s: s,
	}
}

// POST Create a new dentist
func (h *appointmentHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var appointment domain.Appointment
		err := ctx.ShouldBindJSON(&appointment)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid appointment", "details": err.Error()})
			return
		}

		p, err := h.s.Create(appointment)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, p)
	}
}

// GET all appointment
func (h *appointmentHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		appointment, _ := h.s.GetAll()
		c.JSON(200, appointment)
	}
}

// GET appointment by ID
func (h *appointmentHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		appointment, err := h.s.GetByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "appointment not found"})
			return
		}
		c.JSON(200, appointment)
	}
}
