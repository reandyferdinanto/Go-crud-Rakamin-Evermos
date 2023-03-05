package entity

import (
	"time"

	"gorm.io/gorm"
)

type Trx struct {

	Id					int8			`gorm:"primaryKey;autoIncrement" json:"id"`
	Id_user				string			`gorm:"type:varchar(255);index" json:"id_user"`
	User				User			`gorm:"foreignKey:Id_user" json:"user"`
	Alamat_Pengiriman	string			`gorm:"type:varchar(255)" json:"alamat_pengiriman"`
	Harga_total			string			`gorm:"type:varchar(255)" json:"harga_total"`
	Kode_invoice		string			`gorm:"type:varchar(255)" json:"kode_invoice"`
	Method_bayar		string			`gorm:"type:varchar(255)" json:"method_bayar"`
	CreatedAt			time.Time		`json:"createdat"`
	UpdatedAt			time.Time		`json:"updatedat"`
	DeletedAt			gorm.DeletedAt 	`json:"-" gorm:"index"`

}