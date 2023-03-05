package entity

import (
	"time"

	"gorm.io/gorm"
)

type FotoProduk struct {

	Id				int8			`gorm:"primaryKey;autoIncrement" json:"id"`
	Id_produk		int8			`gorm:"type:varchar(255);index" json:"id_produk"`
	Produk			Produk			`gorm:"foreignKey:Id_produk" json:"produk"`
	Url				string			`gorm:"type:varchar(255)" json:"url"`
	CreatedAt		time.Time		`json:"createdat"`
	UpdatedAt		time.Time		`json:"updatedat"`
	DeletedAt		gorm.DeletedAt 	`json:"-" gorm:"index"`

}