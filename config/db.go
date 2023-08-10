package config

import (
	cart "ecommerce/features/cart/data"
	product "ecommerce/features/product/data"
	transaction "ecommerce/features/transaction/data"
	user "ecommerce/features/user/data"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(ac AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5433 sslmode=disable TimeZone=Asia/Jakarta",
		ac.DBHost, ac.DBUser, ac.DBPass, ac.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Database connection error : ", err.Error())
		return nil
	}

	return db
}
func Migrate(db *gorm.DB) {
	db.AutoMigrate(user.Users{})
	db.AutoMigrate(product.Products{})
	db.AutoMigrate(cart.Carts{})
	db.AutoMigrate(transaction.Transactions{})
	db.AutoMigrate(transaction_detail.Transaction_Detail{})
}
