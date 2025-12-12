package model

import "time"

type InfoType string

const (
	TypeHistory   InfoType = "HISTORY"
	TypeCuriosity InfoType = "CURIOSITY"
	TypeFestivity InfoType = "FESTIVITY"
)

type MunicipalityInfo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"size:200;not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	Type      InfoType  `gorm:"size:50;index" json:"type"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
