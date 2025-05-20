package main

import "testing"

func TestMax(t *testing.T) {
	result := max(5,10)
	expected := 10

	if result != expected{
		t.Errorf("Max(5,10) = %d , ожидается %d",result,expected)
	}
}