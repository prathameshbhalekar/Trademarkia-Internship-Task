package models

type Users struct {
	ID       uint   `json:"id" gorm:"primary_key;not null"`
	Name     string `json:"name" gorm:"not null"`
	Location int    `json:"location"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
}
