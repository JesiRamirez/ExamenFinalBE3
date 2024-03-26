package handler

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/appointmentDNILicense"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/gin-gonic/gin"
)

type appointmentDNILicenseHandler struct {
	s appointmentDNILicense.Service
}

// NewAppointmentDNILicenseHandler create a new AppointmentDNILicense controller
func NewAppointmentDNILicenseHandler(s appointmentDNILicense.Service) *appointmentDNILicenseHandler {
	return &appointmentDNILicenseHandler{
		s: s,
	}
}

// POST Create a new dentist
func (h *appointmentDNILicenseHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var appointmentDNILicense domain.AppointmentDNILicense
		err := ctx.ShouldBindJSON(&appointmentDNILicense)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid AppointmentDNILicense", "details": err.Error()})
			return
		}

		p, err := h.s.Create(appointmentDNILicense)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, p)
	}
}

// GET AppointmentDNILicense by ID
func (h *appointmentDNILicenseHandler) GetByPatientDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		dniParam := c.Param("dni")

		appointmentDNILicense, err := h.s.GetByPatientDNI(dniParam)
		if err != nil {
			c.JSON(404, gin.H{"error": "AppointmentDNILicense not found"})
			return
		}
		c.JSON(200, appointmentDNILicense)
	}
}
