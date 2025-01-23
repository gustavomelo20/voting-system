package models

type Candidate struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
