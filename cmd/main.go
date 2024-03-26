package main

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/cmd/handler"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/appointment"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/appointmentDNILicense"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/dentist"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/patient"
	"github.com/gin-gonic/gin"
)

func main() {
	var patientsList = []domain.Patient{}
	var dentistsList = []domain.Dentist{}
	var appointmentList = []domain.Appointment{}
	var appointmentsDNILicenseList = []domain.AppointmentDNILicense{}

	repo := patient.NewRepository(patientsList)
	service := patient.NewService(repo)
	patientHandler := handler.NewPatientHandler(service)

	repoDentist := dentist.NewRepository(dentistsList)
	serviceDentist := dentist.NewService(repoDentist)
	dentistHandler := handler.NewDentistHandler(serviceDentist)

	repoAppointment := appointment.NewRepository(appointmentList)
	serviceAppointment := appointment.NewService(repoAppointment)
	appointmentHandler := handler.NewAppointmentHandler(serviceAppointment)

	repoAppointmentDNILicense := appointmentDNILicense.NewRepository(appointmentsDNILicenseList)
	serviceAppointmentDNILicense := appointmentDNILicense.NewService(repoAppointmentDNILicense)
	appointmentDNILicenseHandler := handler.NewAppointmentDNILicenseHandler(serviceAppointmentDNILicense)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	patients := r.Group("/patients")
	{
		patients.POST("", patientHandler.Post())
		patients.GET("", patientHandler.GetAll())
		patients.GET(":id", patientHandler.GetByID())
	}

	dentists := r.Group("/dentists")
	{
		dentists.POST("", dentistHandler.Post())
		dentists.GET("", dentistHandler.GetAll())
		dentists.GET(":id", dentistHandler.GetByID())
	}

	appointments := r.Group("/appointments")
	{
		appointments.POST("", appointmentHandler.Post())
		appointments.GET("", appointmentHandler.GetAll())
		appointments.GET(":id", appointmentHandler.GetByID())
	}

	appointmentsDNILicense := r.Group("/appointmentsDNI")
	{
		appointmentsDNILicense.POST("", appointmentDNILicenseHandler.Post())
		appointmentsDNILicense.GET(":id", appointmentDNILicenseHandler.GetByPatientDNI())
	}

	r.Run(":8081")

}
