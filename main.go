package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// Conectar a la base de datos
	ConnectDatabase()

	// Crear router
	r := gin.Default()

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

	// ======================
	// DATASENSE
	// ======================
	r.GET("/dataSense/getDataSenseByPetId/:id", GetDataSenseByPetId)
	r.GET("/dataSense/getDataSenseByPetIdAndType/:id/:type", GetDataSenseByPetIdAndType)
	r.POST("/dataSense/createDataSense", CreateDataSense)
	r.DELETE("/dataSense/deleteDataSense/:id", DeleteDataSense)

	// Iniciar servidor
	r.Run(":8080")
}
