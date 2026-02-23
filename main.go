package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {

	// Conectar a la base de datos
	ConnectDatabase()

	// Crear router
	r := gin.Default()

	// ======================
	// CORS (OBLIGATORIO)
	// ======================
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		MaxAge: 12 * time.Hour,
	}))


	// ======================
	// PET
	// ======================
	r.GET("/pet/getPets", GetPets)
	r.GET("/pet/getPetById/:id", GetPetByID)
	r.POST("/pet/createPet", CreatePet)
	r.PUT("/pet/updatePet/:id", UpdatePet)
	r.DELETE("/pet/deletePet/:id", DeletePet)

	// ======================
	// APPOINTMENT
	// ======================
	r.GET("/appointment/getAppointments", GetAppointments)
	r.GET("/appointment/getAppointmentsByPet/:id", GetAppointmentsByPet)
	r.POST("/appointment/createAppointment", CreateAppointment)
	r.DELETE("/appointment/deleteAppointment/:id", DeleteAppointment)

	// ======================
	// DATASENSE
	// ======================
	r.GET("/dataSense/getDataSenseByPetId/:id", GetDataSenseByPetId)
	r.GET("/dataSense/getDataSenseByPetIdAndType/:id/:type", GetDataSenseByPetIdAndType)
	r.POST("/dataSense/createDataSense", CreateDataSense)
	r.DELETE("/dataSense/deleteDataSense/:id", DeleteDataSense)
	r.DELETE("/dataSense/deleteDataSenseByAppointment/:id", DeleteDataSenseByAppointment)

	// USER
	r.POST("/user/login", Login)

	// Iniciar servidor
	r.Run(":8080")
}
