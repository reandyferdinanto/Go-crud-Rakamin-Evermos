package entity

import (
	"time"

	"gorm.io/gorm"
)

type Produk struct {

	Id				int8			`gorm:"primaryKey;autoIncrement" json:"id"`
	Nama_produk		string			`gorm:"type:varchar(255);index" json:"nama_produk"`
	Slug			string			`gorm:"type:varchar(255)" json:"slug"`
	Harga_reseller	string			`gorm:"type:varchar(255)" json:"harga_reseller"`
	Harga_konsumen	string			`gorm:"type:varchar(255)" json:"harga_konsumen"`
	Stok			int32			`gorm:"type:varchar(255)" json:"stok"`
	Deskripsi		string			`gorm:"type:varchar(255)" json:"deskripsi"`
	CreatedAt		time.Time		`json:"createdat"`
	UpdatedAt		time.Time		`json:"updatedat"`
	DeletedAt		gorm.DeletedAt 	`json:"-" gorm:"index"`
	Id_toko			int8			`gorm:"type:varchar(255)" json:"id_toko"`
	Toko			Toko			`gorm:"foreignKey:Id_toko" json:"toko"`
	Id_category		int8			`gorm:"type:varchar(255)" json:"id_category"`
	Category		Category		`gorm:"foreignKey:Id_category" json:"category"`

}