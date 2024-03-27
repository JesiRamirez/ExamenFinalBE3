package handler

import (
	"errors"
	"strconv"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/patient"
	"github.com/bootcamp-go/ExamenFinalBE3.git/pkg/web"
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
// Post godoc
// @Summary      Create a new patient
// @Description  Create a new patient in repository
// @Tags         patients
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Patient true "Patient"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /patients [post]
func (h *patientHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var patient domain.Patient
		err := ctx.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid patient"))
			return
		}

		p, err := h.s.Create(patient)
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		web.Success(ctx, 201, p)
	}
}

// GET all patients
// GetAll godoc
// @Summary      Gets all the patients
// @Description  Gets all the patients from the repository
// @Tags         patients
// @Produce      json
// @Success      200 {object}  web.response
// @Router       /patients [get]
func (h *patientHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		patients, _ := h.s.GetAll()
		c.JSON(200, patients)
	}
}

// GET patient by ID
// GetByID godoc
// @Summary      Gets a patient by id
// @Description  Gets a patient by id from the repository
// @Tags         patients
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /patients/{id} [get]
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
// Put godoc
// @Summary      Updates a patient
// @Description  Updates a patient from the repository
// @Tags         patients
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Param        body body domain.Patient true "Patient"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /patients/{id} [put]
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

		p, err := h.s.Update(id, patient)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(201, p)
	}

}

// PATCH patient
// Patch godoc
// @Summary      Updates selected fields
// @Description  Updates selected fields from a patient from the repository
// @Tags         patients
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Param        body body domain.Patient true "Patient"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /patients/{id} [patch]
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

// DELETE elimina un paciente
// Delete godoc
// @Summary      Deletes a patient
// @Description  Deletes a patient from the repository
// @Tags         patients
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Success      204 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /patients/{id} [delete]
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
