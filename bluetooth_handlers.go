package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ================================================================
//  NUEVAS RUTAS PARA DATOS BLUETOOTH EN TIEMPO REAL
// ================================================================

// BluetoothSensorData es el formato de datos que envía el Bluetooth Bridge
type BluetoothSensorData struct {
	// Datos de pesos (ESP32 #1)
	WeightDistributionLF float64            `json:"weightDistributionLF"`
	WeightDistributionRF float64            `json:"weightDistributionRF"`
	WeightDistributionLB float64            `json:"weightDistributionLB"`
	WeightDistributionRB float64            `json:"weightDistributionRB"`
	TotalWeight          float64            `json:"totalWeight"`
	COP                  json.RawMessage    `json:"cop"`

	// Datos de rotación (ESP32 #2)
	Gyroscope     json.RawMessage `json:"gyroscope"`
	Accelerometer json.RawMessage `json:"accelerometer"`
	Angles        json.RawMessage `json:"angles"`
	Temperature   float64         `json:"temperature"`

	// Metadatos
	Timestamp string `json:"timestamp"`
}

// BluetoothDataRequest es el formato para guardar datos en una cita específica
type BluetoothDataRequest struct {
	AppointmentID uint                  `json:"appointmentId" binding:"required"`
	SensorData    BluetoothSensorData   `json:"sensorData" binding:"required"`
	Type          string                `json:"type"`
	TotalTime     string                `json:"totalTime"`
	FrequencyHZ   int                   `json:"frequencyHZ"`
	AmplitudeMV   int                   `json:"amplitudeMV"`
}

// ================================================================
//  HANDLER: Guardar datos Bluetooth en tiempo real
// ================================================================
func SaveBluetoothData(c *gin.Context) {
	var req BluetoothDataRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[ERROR] Error parseando datos Bluetooth: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear objeto DataSense con los datos recibidos
	dataSense := DataSense{
		AppointmentID: req.AppointmentID,
		Type:          req.Type,
		TotalTime:     req.TotalTime,
		FrequencyHZ:   req.FrequencyHZ,
		AmplitudeMV:   req.AmplitudeMV,

		// Datos de pesos
		WeightDistributionLF: json.RawMessage(`{"value": ` + floatToString(req.SensorData.WeightDistributionLF) + `}`),
		WeightDistributionRF: json.RawMessage(`{"value": ` + floatToString(req.SensorData.WeightDistributionRF) + `}`),
		WeightDistributionLB: json.RawMessage(`{"value": ` + floatToString(req.SensorData.WeightDistributionLB) + `}`),
		WeightDistributionRB: json.RawMessage(`{"value": ` + floatToString(req.SensorData.WeightDistributionRB) + `}`),

		// Centro de presión
		COPN: req.SensorData.COP,
		COPC: req.SensorData.COP,

		// Datos de rotación
		Gyroscope:     req.SensorData.Gyroscope,
		Accelerometer: req.SensorData.Accelerometer,

		// Fuerza vertical (usando temperatura como placeholder)
		VerticalForce: json.RawMessage(`{"value": ` + floatToString(req.SensorData.Temperature) + `}`),
	}

	// Guardar en base de datos
	result := DB.Create(&dataSense)

	if result.Error != nil {
		log.Printf("[ERROR] Error guardando datos en BD: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	log.Printf("[OK] Datos Bluetooth guardados - AppointmentID: %d", req.AppointmentID)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Datos guardados exitosamente",
		"id":      dataSense.ID,
	})
}

// ================================================================
//  HANDLER: Guardar streaming continuo de datos Bluetooth
// ================================================================
func SaveBluetoothStreamData(c *gin.Context) {
	var data BluetoothSensorData

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Este endpoint es para guardar datos en tiempo real sin appointmentID específico
	// Útil para almacenamiento temporal o visualización en vivo

	log.Printf("[STREAM] Datos recibidos - Pesos: LF=%.2f, RF=%.2f, LB=%.2f, RB=%.2f",
		data.WeightDistributionLF, data.WeightDistributionRF,
		data.WeightDistributionLB, data.WeightDistributionRB)

	c.JSON(http.StatusOK, gin.H{"status": "datos recibidos"})
}

// ================================================================
//  HANDLER: Obtener últimos datos de un perro
// ================================================================
func GetLatestBluetoothData(c *gin.Context) {
	petID := c.Param("id")

	var latestData DataSense

	// Obtener el último registro de datos para este perro
	result := DB.
		Joins("JOIN appointment ON dataSense.idAppointment = appointment.id").
		Where("appointment.pet_id = ?", petID).
		Order("dataSense.id DESC").
		First(&latestData)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No data found"})
		return
	}

	c.JSON(http.StatusOK, latestData)
}

// ================================================================
//  HANDLER: Iniciar nueva cita con datos Bluetooth
// ================================================================
type StartTestRequest struct {
	PetID       uint   `json:"petId" binding:"required"`
	TestType    string `json:"testType"`
	Duration    int    `json:"duration"` // en segundos
}

type StartTestResponse struct {
	AppointmentID uint   `json:"appointmentId"`
	Message       string `json:"message"`
}

func StartBluetoothTest(c *gin.Context) {
	var req StartTestRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear nueva cita
	appointment := Appointment{
		PetID: req.PetID,
		Date:  nowTime(),
	}

	result := DB.Create(&appointment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	log.Printf("[TEST_START] Nueva prueba iniciada - PetID: %d, AppointmentID: %d", req.PetID, appointment.ID)

	c.JSON(http.StatusCreated, StartTestResponse{
		AppointmentID: appointment.ID,
		Message:       "Prueba iniciada - comienza la recolección de datos",
	})
}

// ================================================================
//  HANDLER: Finalizar prueba Bluetooth
// ================================================================
type EndTestRequest struct {
	AppointmentID uint   `json:"appointmentId" binding:"required"`
	Notes         string `json:"notes"`
}

func EndBluetoothTest(c *gin.Context) {
	var req EndTestRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar que la cita existe
	var appointment Appointment
	if err := DB.First(&appointment, req.AppointmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	log.Printf("[TEST_END] Prueba finalizada - AppointmentID: %d", req.AppointmentID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Prueba finalizada",
		"notes":   req.Notes,
	})
}

// ================================================================
//  HANDLER: Obtener estado de conexión Bluetooth
// ================================================================
func GetBluetoothStatus(c *gin.Context) {
	// Este endpoint podría verificar el estado de conexión del Bluetooth Bridge
	// Por ahora, devuelve un estado simulado

	status := gin.H{
		"connected":     true,
		"esp32_weights": true,
		"esp32_rotation": true,
		"last_data":     "2024-12-19T10:30:45Z",
		"data_rate":     "20 Hz",
	}

	c.JSON(http.StatusOK, status)
}

// ================================================================
//  FUNCIONES AUXILIARES
// ================================================================

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func nowTime() time.Time {
	return time.Now()
}
