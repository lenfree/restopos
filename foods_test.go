package main

import (
	"testing"
)

func TestCreateFood(t *testing.T) {
	food := CreateFood("lechon", 20.50)
	if food.Name != "lechon" {
		t.Errorf("got %s, want lechon", food.Name)
	}

	if food.Price != 20.50 {
		t.Errorf("got %f, want 20.50", food.Price)
	}

	DeleteFood("lechon")
}

func TestGetFood(t *testing.T) {
	CreateFood("lechon", 20.50)
	foodItem, _ := GetFood("lechon")
	if foodItem.Name != "lechon" {
		t.Error("got %s, want lechon", foodItem.Name)
	}

	DeleteFood("lechon")
}

func TestDeleteFood(t *testing.T) {
	CreateFood("lechon", 20.50)
	count := DeleteFood("lechon")
	if count != 1 {
		t.Error("got %d, want 1", count)
	}
}
