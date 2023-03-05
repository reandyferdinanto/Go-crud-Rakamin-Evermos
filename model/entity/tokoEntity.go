package entity

import (
	"time"

	"gorm.io/gorm"
)

type Toko struct {

	Id				int8			`gorm:"primaryKey;autoIncrement" json:"id"`
	Id_user			string			`gorm:"type:varchar(255);index" json:"id_user"`
	User			User			`gorm:"foreignKey:Id_user" json:"user"`
	Nama_toko		string			`gorm:"type:varchar(255)" json:"nama_toko"`
	Url_foto		string			`gorm:"type:varchar(255)" json:"url_foto"`
	CreatedAt		time.Time		`json:"createdat"`
	UpdatedAt		time.Time		`json:"updatedat"`
	DeletedAt		gorm.DeletedAt 	`json:"-" gorm:"index"`

}