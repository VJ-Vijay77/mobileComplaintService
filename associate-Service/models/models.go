package models

import "gorm.io/gorm"

type Associate struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Phone     int    `json:"phone"`
}

type Complaint struct {
	Tokenno       int    `gorm:"primaryKey"`
	UserID        uint   `json:"id,omitempty"`
	Brand         string `json:"brand,omitempty"`
	Model         string `json:"model,omitempty"`
	PurchaseDate  string `json:"purchase_date,omitempty"`
	Complaints    string `json:"complaint,omitempty"`
	Complaintdate string `json:"complaintdate,omitempty"`
	Status        string `json:"status,omitempty"`
	Resolvedate   string `json:"resolvedate,omitempty"`
	Associateid   uint   `json:"associateid,omitempty"`
}
