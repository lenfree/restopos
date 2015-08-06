package main

import "testing"

func TestCreateFood(t *testing.T) {
	food := CreateFood("lechon", 20.50)
	if food.Name != "lechon" {
		t.Errorf("got %s, want lechon", food.Name)
	}

	if food.Price != 20.50 {
		t.Errorf("got %f, want 20.50", food.Price)
	}
}
