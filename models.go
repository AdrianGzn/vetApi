package main

import "time"
import "encoding/json"

type Pet struct {
	ID                uint      `gorm:"primaryKey;column:id" json:"id"`
	Group             string    `gorm:"column:group" json:"group"`
	Control           string    `gorm:"column:control" json:"control"`
	Race              string    `gorm:"column:race" json:"race"`
	Age               int       `gorm:"column:age" json:"age"`
	Gender            string    `gorm:"column:gender" json:"gender"`
	Weight            string    `gorm:"column:weight" json:"weight"`
	BodyCondition     int       `gorm:"column:bodyCondition" json:"bodyCondition"`
	Diagnosis         string    `gorm:"column:diagnosis" json:"diagnosis"`
	DegreeLameness    int       `gorm:"column:degreeLameness" json:"degreeLameness"`
	OnsetTimeSymptoms time.Time `gorm:"column:onsetTimeSymptoms" json:"onsetTimeSymptoms"`
	Name              string    `gorm:"column:name" json:"name"`
	Owner             string    `gorm:"column:owner" json:"owner"`
	Color             string    `gorm:"column:color" json:"color"`
	LastAppointment   time.Time `gorm:"column:lastAppointment" json:"lastAppointment"`
	Image             string    `gorm:"column:image" json:"image"`
}

func (Pet) TableName() string {
	return "pet"
}


type Appointment struct {
	ID    uint      `gorm:"primaryKey" json:"id"`
	PetID uint      `gorm:"column:pet_id" json:"pet_id"`
	Date  time.Time `json:"date"`
}

func (Appointment) TableName() string {
	return "appointment"
}

type DataSense struct {
	ID              uint            `gorm:"primaryKey;column:id" json:"id"`
	AppointmentID   uint            `gorm:"column:idAppointment" json:"idAppointment"`
	Type            string          `gorm:"column:type" json:"type"`
	TotalTime       string          `gorm:"column:totalTime" json:"totalTime"`
	FrequencyHZ     int             `gorm:"column:frequencyHZ" json:"frequencyHZ"`
	AmplitudeMV     int             `gorm:"column:amplitudeMV" json:"amplitudeMV"`

	COPN            json.RawMessage `gorm:"column:COPN" json:"COPN"`
	COPC            json.RawMessage `gorm:"column:COPC" json:"COPC"`
	Result          json.RawMessage `gorm:"column:result" json:"result"`

	Gyroscope       json.RawMessage `gorm:"column:gyroscope;type:json" json:"gyroscope"`
	Accelerometer   json.RawMessage `gorm:"column:accelerometer;type:json" json:"accelerometer"`

	SymmetryIndexLF json.RawMessage `gorm:"column:symmetryIndexLF" json:"symmetryIndexLF"`
	SymmetryIndexRF json.RawMessage `gorm:"column:symmetryIndexRF" json:"symmetryIndexRF"`
	SymmetryIndexLB json.RawMessage `gorm:"column:symmetryIndexLB" json:"symmetryIndexLB"`
	SymmetryIndexRB json.RawMessage `gorm:"column:symmetryIndexRB" json:"symmetryIndexRB"`

	WeightDistributionLF json.RawMessage `gorm:"column:weightDistributionLF" json:"weightDistributionLF"`
	WeightDistributionRF json.RawMessage `gorm:"column:weightDistributionRF" json:"weightDistributionRF"`
	WeightDistributionLB json.RawMessage `gorm:"column:weightDistributionLB" json:"weightDistributionLB"`
	WeightDistributionRB json.RawMessage `gorm:"column:weightDistributionRB" json:"weightDistributionRB"`

	VerticalForce   json.RawMessage `gorm:"column:verticalForce" json:"verticalForce"`
	VerticalImpulse string          `gorm:"column:verticalImpulse" json:"verticalImpulse"`
}

func (DataSense) TableName() string {
	return "dataSense"
}

type User struct {
	ID       uint   `gorm:"primaryKey;column:id" json:"id"`
	Name     string `gorm:"column:name;not null" json:"name"`
	Password string `gorm:"column:password;not null" json:"-"`
}

func (User) TableName() string {
	return "user"
}