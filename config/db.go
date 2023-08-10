package config

import (
	"fmt"
	"log"

	cart "github.com/wussh/tokokentang/features/cart/data"
	product "github.com/wussh/tokokentang/features/product/data"
	transaction "github.com/wussh/tokokentang/features/transaction/data"
	transaction_detail "github.com/wussh/tokokentang/features/transaction_detail/data"
	user "github.com/wussh/tokokentang/features/user/data"

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
