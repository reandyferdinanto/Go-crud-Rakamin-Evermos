package entity

import (
	"time"

	"gorm.io/gorm"
)

type Log_Produk struct {

	Id					int8			`gorm:"primaryKey;autoIncrement" json:"id"`
	Id_produk			string			`gorm:"type:varchar(255);index" json:"id_produk"`
	Produk				Produk			`gorm:"foreignKey:Id_produk" json:"produk"`
	Nama_produk			string			`gorm:"type:varchar(255)" json:"nama_produk"`
	Slug				string			`gorm:"type:varchar(255)" json:"slug"`
	Harga_reseller		string			`gorm:"type:varchar(255)" json:"harga_reseller"`
	Harga_konsumen		string			`gorm:"type:varchar(255)" json:"harga_konsumen"`
	Deskripsi			string			`gorm:"type:varchar(255)" json:"deskripsi"`
	CreatedAt			time.Time		`json:"createdat"`
	UpdatedAt			time.Time		`json:"updatedat"`
	DeletedAt			gorm.DeletedAt 	`json:"-" gorm:"index"`
	Id_toko				string			`gorm:"type:varchar(255);index" json:"id_toko"`
	Toko				Toko			`gorm:"foreignKey:Id_toko" json:"toko"`
	Id_category			string			`gorm:"type:varchar(255);index" json:"id_category"`
	Category			Category		`gorm:"foreignKey:Id_Category" json:"category"`
}