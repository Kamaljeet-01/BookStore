package main

var Shelf []Book

type Book struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type UpdatedBook struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func Find(id int) int {
	for i, b := range Shelf {
		if b.Id == id {
			return i
		}
	}
	return -1
}
