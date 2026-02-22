package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


// ======================
// PET HANDLERS
// ======================

func GetPets(c *gin.Context) {
	var pets []Pet
	DB.Find(&pets)
	c.JSON(http.StatusOK, pets)
}

func GetPetByID(c *gin.Context) {
	id := c.Param("id")

	var pet Pet
	if err := DB.First(&pet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}

	c.JSON(http.StatusOK, pet)
}

func CreatePet(c *gin.Context) {
	var pet Pet

	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := DB.Create(&pet)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No rows inserted",
		})
		return
	}

	c.JSON(http.StatusCreated, pet)
}


func UpdatePet(c *gin.Context) {
	id := c.Param("id")

	var pet Pet
	if err := DB.First(&pet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}

	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Save(&pet)
	c.JSON(http.StatusOK, pet)
}

func DeletePet(c *gin.Context) {
	id := c.Param("id")

	if err := DB.Delete(&Pet{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error deleting pet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet deleted"})
}


// ======================
// APPOINTMENT HANDLERS
// ======================

func GetAppointments(c *gin.Context) {
	var appointments []Appointment
	DB.Find(&appointments)
	c.JSON(http.StatusOK, appointments)
}

func GetAppointmentsByPet(c *gin.Context) {
	id := c.Param("id")

	var appointments []Appointment
	DB.Where("pet_id = ?", id).Find(&appointments)

	c.JSON(http.StatusOK, appointments)
}

func CreateAppointment(c *gin.Context) {
	var appointment Appointment

	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Create(&appointment)
	c.JSON(http.StatusOK, appointment)
}

func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")

	// Primero, eliminar todos los dataSense asociados a esta cita
	if err := DB.Where("idAppointment = ?", id).Delete(&DataSense{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Luego eliminar la cita
	if err := DB.Delete(&Appointment{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}


// ======================
// DATASENSE HANDLERS
// ======================

func GetDataSenseByPetId(c *gin.Context) {
	petID := c.Param("id")

	var data []DataSense

	subQuery := DB.
		Table("appointment").
		Select("id").
		Where("pet_id = ?", petID)

	result := DB.
		Where("idAppointment IN (?)", subQuery).
		Find(&data)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}


func GetDataSenseByPetIdAndType(c *gin.Context) {
	petID := c.Param("id")
	typeParam := c.Param("type")

	var data []DataSense

	subQuery := DB.
		Table("appointment").
		Select("id").
		Where("pet_id = ?", petID)

	result := DB.
		Where("idAppointment IN (?) AND type = ?", subQuery, typeParam).
		Find(&data)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}


func CreateDataSense(c *gin.Context) {
	var data DataSense

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Create(&data)
	c.JSON(http.StatusOK, data)
}

func DeleteDataSense(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := DB.Delete(&DataSense{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error deleting DataSense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "DataSense deleted"})
}


func DeleteDataSenseByAppointment(c *gin.Context) {
	appointmentID := c.Param("id")

	if err := DB.Where("idAppointment = ?", appointmentID).Delete(&DataSense{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All DataSense for appointment deleted successfully"})
}