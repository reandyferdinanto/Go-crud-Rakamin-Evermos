package entity

import (
	"time"

	"gorm.io/gorm"
)

type DetailTrx struct {

	Id				int8			`gorm:"primaryKey;autoIncrement" json:"id"`
	Id_trx			string			`gorm:"type:varchar(255);index" json:"id_trx"`
	Trx				Trx				`gorm:"foreignKey:Id_trx" json:"trx"`
	Id_log_produk	int8			`gorm:"type:varchar(255);index" json:"id_log_produk"`
	Log_produk		Log_Produk		`gorm:"foreignKey:Id_log_produk" json:"log_produk"`
	Id_toko			int8			`gorm:"type:varchar(255);index" json:"id_toko"`
	Toko			Toko			`gorm:"foreignKey:Id_toko" json:"toko"`
	Kuantitas		string			`gorm:"type:varchar(255)" json:"kuantitas"`
	Harga_total		string			`gorm:"type:varchar(255)" json:"harga_total"`
	CreatedAt		time.Time		`json:"createdat"`
	UpdatedAt		time.Time		`json:"updatedat"`
	DeletedAt		gorm.DeletedAt 	`json:"-" gorm:"index"`

}