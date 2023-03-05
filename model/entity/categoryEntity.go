package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {

	Id				int8			`gorm:"primaryKey;autoIncrement" json:"id"`
	Nama_category	string			`gorm:"type:varchar(255);index" json:"nama_category"`
	CreatedAt		time.Time		`json:"createdat"`
	UpdatedAt		time.Time		`json:"updatedat"`
	DeletedAt		gorm.DeletedAt 	`json:"-" gorm:"index"`

}