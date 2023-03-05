package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {

	Id				int8			`gorm:"primaryKey;autoIncrement" json:"id"`
	Name			string			`gorm:"type:varchar(255)" json:"name"`
	Password		string			`json:"password" gorm:"column:password"`
	No_telp			string 			`gorm:"type:varchar(255);uniqueIndex" json:"no_telp"`
	Tanggal_lahir	time.Time		`gorm:"type:date" json:"tanggal_lahir"`
	Jenis_kelamin	string			`gorm:"type:varchar(255)" json:"jenis_kelamin"`
	Tentang			string			`json:"tentang"`
	Pekerjaan		string			`gorm:"type:varchar(255)" json:"pekerjaan"`
	Email			string			`gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Id_provinsi		string			`gorm:"type:varchar(255)" json:"id_provinsi"`
	Id_kota			string			`gorm:"type:varchar(255)" json:"id_kota"`	
	IsAdmin			bool			`json:"isAdmin"`	
	CreatedAt		time.Time		`json:"createdat"`
	UpdatedAt		time.Time		`json:"updatedat"`
	DeletedAt		gorm.DeletedAt 	`json:"-" gorm:"index"`

}