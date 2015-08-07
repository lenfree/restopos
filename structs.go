package main

type Food struct {
	Id         int64 `db:"food_id"`
	Created    int64
	Name       string
	Price      float32
	CategoryId int64
}

type Category struct {
	Id   int64 `db:"category_id"`
	Name string
}
