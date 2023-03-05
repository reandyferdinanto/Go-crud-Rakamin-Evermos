package request

import "time"

type UserCreateRequest struct {
	Name          string    `json:"name" validate:"required"`
	Password      string    `json:"password" validate:"required"`
	No_telp       string    `json:"no_telp"`
	Tanggal_lahir time.Time `gorm:"type:date" json:"tanggal_lahir"`
	Jenis_kelamin string    `json:"jenis_kelamin"`
	Tentang       string    `json:"tentang"`
	Pekerjaan     string    `json:"pekerjaan"`
	Email         string    `json:"email" validate:"required,email"`
	Id_provinsi   string    `json:"id_provinsi"`
	Id_kota       string    `json:"id_kota"`
	IsAdmin       bool      `json:"isAdmin"`
}

type UserUpdateRequest struct {
	Name    			string 		`json:"name"`
	No_telp       		string    	`json:"no_telp"`
	Tanggal_lahir 		time.Time 	`gorm:"type:date" json:"tanggal_lahir"`
	Jenis_kelamin 		string    	`json:"jenis_kelamin"`
	Tentang       		string    	`json:"tentang"`
	Pekerjaan     		string    	`json:"pekerjaan"`
}

type UserEmailRequest struct {
	Email string `json:"email" validate:"required"`
}