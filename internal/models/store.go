package models

import "gorm.io/gorm"

var Shelf []Book

type Book struct {
	gorm.Model
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type UpdatedBook struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
