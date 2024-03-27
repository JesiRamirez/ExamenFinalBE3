package main

import (
	"database/sql"
	"log"

	"github.com/bootcamp-go/ExamenFinalBE3.git/cmd/handler"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/appointment"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/appointmentDNILicense"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/dentist"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/patient"
	"github.com/bootcamp-go/ExamenFinalBE3.git/pkg/store/AppointmentStore"
	"github.com/bootcamp-go/ExamenFinalBE3.git/pkg/store/dentistStore"
	"github.com/bootcamp-go/ExamenFinalBE3.git/pkg/store/patient"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//storage := store.NewJsonStore("./patients.json")
	//storageDentist := store.NewJsonStoreDentist("./dentists.json")
	
	bd, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/turnos-odontologia")
	if err != nil{
		log.Fatal(err)
	}

	storage := store.NewSqlStorePatient(bd)
	storageDentist := dentistStore.NewSqlStoreDentist(bd)
	storageAppointment := appointmentStore.NewSqlStoreAppointment(bd)


	//var patientsList = []domain.Patient{}
	//var dentistsList = []domain.Dentist{}
	//var appointmentList = []domain.Appointment{}
	var appointmentsDNILicenseList = []domain.AppointmentDNILicense{}

	repo := patient.NewRepository(storage)
	service := patient.NewService(repo)
	patientHandler := handler.NewPatientHandler(service)

	repoDentist := dentist.NewRepository(storageDentist)
	serviceDentist := dentist.NewService(repoDentist)
	dentistHandler := handler.NewDentistHandler(serviceDentist)

	repoAppointment := appointment.NewRepository(storageAppointment)
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
		patients.PUT(":id", patientHandler.Put())
		patients.PATCH(":id", patientHandler.Patch())
		patients.DELETE(":id", patientHandler.Delete())
	}

	dentists := r.Group("/dentists")
	{
		dentists.POST("", dentistHandler.Post())
		dentists.GET("", dentistHandler.GetAll())
		dentists.GET(":id", dentistHandler.GetByID())
		dentists.PUT(":id", dentistHandler.Put())
		dentists.PATCH(":id", dentistHandler.Patch())
		dentists.DELETE(":id", dentistHandler.Delete())
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
		appointmentsDNILicense.GET("/dni", appointmentDNILicenseHandler.GetByPatientDNI())
	}

	r.Run(":8081")

}
