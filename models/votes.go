package models

type Votes struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	Candidate int  `json:"candidate"`
}
