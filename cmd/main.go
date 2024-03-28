package main

import (
	"os"

	"github.com/bootcamp-go/ExamenFinalBE3.git/cmd/docs"
	"github.com/bootcamp-go/ExamenFinalBE3.git/cmd/handler"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/appointment"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/appointmentDNILicense"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/dentist"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/patient"
	"github.com/bootcamp-go/ExamenFinalBE3.git/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Certified Tech Developer - Final Exam BackEnd III
// @version 1.0
// @description Dentist Service
// @termsOfService https://developers.ctd.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support - Danna Velasquez & Jesi Ramirez
// @contact.url https://developers.ctd.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

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

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	patients := r.Group("/patients")
	{
		patients.POST("", middleware.Authentication(), patientHandler.Post())
		patients.GET("", patientHandler.GetAll())
		patients.GET(":id", patientHandler.GetByID())
		patients.PUT(":id", middleware.Authentication(), patientHandler.Put())
		patients.PATCH(":id", middleware.Authentication(), patientHandler.Patch())
		patients.DELETE(":id", middleware.Authentication(), patientHandler.Delete())
	}

	dentists := r.Group("/dentists")
	{
		dentists.POST("", middleware.Authentication(), dentistHandler.Post())
		dentists.GET("", dentistHandler.GetAll())
		dentists.GET(":id", dentistHandler.GetByID())
		dentists.PUT(":id", middleware.Authentication(), dentistHandler.Put())
		dentists.PATCH(":id", middleware.Authentication(), dentistHandler.Patch())
		dentists.DELETE(":id", middleware.Authentication(), dentistHandler.Delete())
	}

	appointments := r.Group("/appointments")
	{
		appointments.POST("", appointmentHandler.Post(), appointmentHandler.Post())
		appointments.GET("", appointmentHandler.GetAll())
		appointments.GET(":id", appointmentHandler.GetByID())
		appointments.PUT(":id", appointmentHandler.Put())
		appointments.PATCH(":id", appointmentHandler.Patch())
		appointments.DELETE(":id", appointmentHandler.Delete())
	}

	appointmentsDNILicense := r.Group("/appointmentsDNI")
	{
		appointmentsDNILicense.POST("", appointmentDNILicenseHandler.Post())
		appointmentsDNILicense.GET("/dni", appointmentDNILicenseHandler.GetByPatientDNI())
	}

	r.Run(":8081")

}
