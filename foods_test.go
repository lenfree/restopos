package main

import (
	"testing"
)

func TestCreateFood(t *testing.T) {
	dbmap.CreateTablesIfNotExists()
	category := CreateCategory("ulam")
	food := CreateFood("lechon", 20.50, category.Id)
	if food.Name != "lechon" {
		t.Errorf("got %s, want lechon", food.Name)
	}

	if food.Price != 20.50 {
		t.Errorf("got %f, want 20.50", food.Price)
	}

	dbmap.DropTables()
}

func TestGetFood(t *testing.T) {
	dbmap.CreateTablesIfNotExists()
	category := CreateCategory("ulam")
	CreateFood("lechon", 20.50, category.Id)
	foodItem, _ := GetFood("lechon")
	if foodItem.Name != "lechon" {
		t.Error("got %s, want lechon", foodItem.Name)
	}

	dbmap.DropTables()
}

func TestDeleteFood(t *testing.T) {
	dbmap.CreateTablesIfNotExists()
	category := CreateCategory("ulam")
	CreateFood("lechon", 20.50, category.Id)
	count := DeleteFood("lechon")
	if count != 1 {
		t.Error("got %d, want 1", count)
	}

	dbmap.DropTables()
}

func TestUpdateFood(t *testing.T) {
	dbmap.CreateTablesIfNotExists()
	category := CreateCategory("ulam")
	foodItem := CreateFood("lechon", 20.50, category.Id)
	count := UpdateFood(foodItem.Name, "adobo", 22.10)

	if count != 1 {
		t.Error("got %d, want 1", count)
	}

	dbmap.DropTables()
}
