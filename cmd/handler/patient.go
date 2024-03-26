package handler

import (
	"errors"
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

// UPDATE patient
func (h *patientHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idString := ctx.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}

		var patient domain.Patient
		err = ctx.ShouldBindJSON(&patient)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid patient"})
			return
		}

		valid, err := validateEmptys(&patient)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		p, err := h.s.Update(id, patient)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(201, p)
	}

}

// Patch patient
func (h *patientHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name     string `json:"name,omitempty"`
		Lastname string `json:"last_name,omitempty"`
		Address  string `json:"adress,omitempty"`
		DNI      string `json:"dni,omitempty"`
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
		update := domain.Patient{
			Name:     r.Name,
			Lastname: r.Lastname,
			Address:  r.Address,
			DNI:      r.DNI,
		}
		p, err := h.s.Patch(id, update)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, p)

	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptys(patient *domain.Patient) (bool, error) {
	switch {
	case patient.Name == "" || patient.Lastname == "" || patient.Address == "" || patient.DNI == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil

}

// DELETE elimina un paciente
func (h *patientHandler) Delete() gin.HandlerFunc {
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
		ctx.JSON(204, gin.H{"msg": "patient deleted"})
	}
}
