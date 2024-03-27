package handler

import (
	"errors"
	"strconv"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/appointment"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/bootcamp-go/ExamenFinalBE3.git/pkg/web"
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
			web.Failure(ctx, 400, errors.New("invalid appointment"))
			return
		}

		p, err := h.s.Create(appointment)
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		web.Success(ctx, 201, p)
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

// UPDATE appointment
func (h *appointmentHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idString := ctx.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}

		var appointment domain.Appointment
		err = ctx.ShouldBindJSON(&appointment)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid appointment"})
			return
		}

		p, err := h.s.Update(id, appointment)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(201, p)
	}

}

// Patch appointment
func (h *appointmentHandler) Patch() gin.HandlerFunc {
	type Request struct {
		PatientId   string `json:"patient_id,omitempty"`
		DentistId   string `json:"dentist_id,omitempty"`
		Description string `json:"description,omitempty"`
	}

	return func(ctx *gin.Context) {
		var r Request
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(400, gin.H{"error": "invalid json"})
			return
		}
		update := domain.Appointment{
			PatientId:   r.PatientId,
			DentistId:   r.DentistId,
			Description: r.Description,
		}
		p, err := h.s.Patch(id, update)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, p)

	}
}

// DELETE elimina un appointment
func (h *appointmentHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(204, gin.H{"msg": "appointment deleted"})
	}
}
