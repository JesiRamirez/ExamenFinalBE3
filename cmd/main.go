package main

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/cmd/handler"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/patient"
	"github.com/gin-gonic/gin"
)

func main() {
	var patientsList = []domain.Patient{}

	repo := patient.NewRepository(patientsList)
	service := patient.NewService(repo)
	patientHandler := handler.NewPatientHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	patients := r.Group("/patients")
	{
		patients.POST("", patientHandler.Post())
		patients.GET("", patientHandler.GetAll())
		patients.GET(":id", patientHandler.GetByID())
	}
	r.Run(":8081")
}
