package handler

import (
	"strconv"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/patient"
	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	s patient.Service
}

// NewPatientHandler create a new patient controller
func NewPatientHandler(s patient.Service) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

// POST Create a new patient
func (h *patientHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var patient domain.Patient
		err := ctx.ShouldBindJSON(&patient)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid patient"})
			return
		}

		p, err := h.s.Create(patient)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, p)
	}
}

// GET all patients
func (h *patientHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		patients, _ := h.s.GetAll()
		c.JSON(200, patients)
	}
}

// GET patient by ID
func (h *patientHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		patient, err := h.s.GetByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "patient not found"})
			return
		}
		c.JSON(200, patient)
	}
}
