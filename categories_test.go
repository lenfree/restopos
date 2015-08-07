package main

import (
	"testing"
)

func TestCreateCategory(t *testing.T) {
	dbmap.CreateTablesIfNotExists()
	category := CreateCategory("ulam")
	if category.Name != "ulam" {
		t.Errorf("got %s, want ulam", category.Name)
	}

	dbmap.DropTables()
}

func TestGetCategory(t *testing.T) {
	dbmap.CreateTablesIfNotExists()
	CreateCategory("ulam")
	categoryItem, _ := GetCategory("ulam")
	if categoryItem.Name != "ulam" {
		t.Error("got %s, want ulam", categoryItem.Name)
	}

	dbmap.DropTables()
}

func TestDeleteCategory(t *testing.T) {
	dbmap.CreateTablesIfNotExists()
	CreateCategory("ulam")
	count := DeleteCategory("ulam")
	if count != 1 {
		t.Error("got %d, want 1", count)
	}
}

func TestUpdateCategory(t *testing.T) {
	dbmap.CreateTablesIfNotExists()
	CreateCategory("ulam")
	CategoryItem, _ := GetCategory("ulam")

	CategoryItem.Name = "tomada"

	count, _ := dbmap.Update(&CategoryItem)

	if count != 1 {
		t.Error("got %d, want 1", count)
	}

	dbmap.DropTables()
}
