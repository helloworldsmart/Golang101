package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Author      string
	ISBN        string
	PublishDate string
	Price       string
}

type Customer struct {
	gorm.Model
	Name    string
	Email   string
	Photo   string
	Address string
}

type Order struct {
	gorm.Model
	CustomerID uint
	Total      float64
}

type OrderDetail struct {
	gorm.Model
	OrderID   uint
	BookID    int
	Quantity  int
	UnitPrice float64
}

func main() {
	// TODO:
	dsn := "host=localhost port=5432 user=youruser dbname=yourdb sslmode=disable password=yourpassword"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自動遷移模式
	db.AutoMigrate(&Book{}, &Customer{}, &Order{}, &OrderDetail{})

}
