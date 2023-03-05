package migration

import (
	"fmt"
	"log"
	"tryFiberGo/database"
	"tryFiberGo/model/entity"
)

func RunMigration(){
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Toko{}, &entity.Alamat{}, &entity.Category{},  &entity.DetailTrx{},  &entity.FotoProduk{},  &entity.Log_Produk{},  &entity.Produk{},  &entity.Trx{}  )
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}

