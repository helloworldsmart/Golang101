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
	Price       float64
	Stock       int
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

	// 添加書籍
	book := Book{
		Title:       "The Go Programming Language",
		Author:      "Alan A. A. Donovan & Brian W. Kernighan",
		PublishDate: "2015-10-26",
		Price:       46.99,
		Stock:       10,
	}
	err = AddBook(db, book)
	if err != nil {
		panic("failed to add book")
	}

	// 刪除書籍
	//err = DeleteBook(db, book.ID)
	//if err != nil {
	//	panic("failed to delete book")
	//}

	// 更新書籍庫存，假設增加 5 本
	err = UpdateBookStock(db, book.ID, 5)
	if err != nil {
		panic("failed to update book stock")
	}
}

func AddBook(db *gorm.DB, book Book) error {
	result := db.Create(&book)
	return result.Error
}

func DeleteBook(db *gorm.DB, bookID uint) error {
	result := db.Delete(&Book{}, bookID)
	return result.Error
}

func UpdateBookStock(db *gorm.DB, bookID uint, quantity int) error {
	var book Book
	// 搜索書籍
	result := db.First(&book, bookID)
	if result.Error != nil {
		return result.Error
	}
	// 更新庫存
	book.Stock += quantity
	result = db.Save(&book)
	return result.Error
}
