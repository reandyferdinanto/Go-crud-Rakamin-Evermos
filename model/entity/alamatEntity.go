package entity

import (
	"time"

	"gorm.io/gorm"
)

type Alamat struct {

	Id				int8			`gorm:"primaryKey;autoIncrement" json:"id"`
	Id_user			string			`gorm:"type:varchar(255);index" json:"id_user"`
	User			User			`gorm:"foreignKey:Id_user" json:"user"`
	Judul_alamat	string			`gorm:"type:varchar(255)" json:"judul_alamat"`
	Nama_penerima	string			`gorm:"type:varchar(255)" json:"nama_penerima"`
	No_telp			string			`gorm:"type:varchar(255)" json:"no_telp"`
	Detail_Alamat	string			`gorm:"type:varchar(255)" json:"detail_alamat"`
	CreatedAt		time.Time		`json:"createdat"`
	UpdatedAt		time.Time		`json:"updatedat"`
	DeletedAt		gorm.DeletedAt 	`json:"-" gorm:"index"`

}