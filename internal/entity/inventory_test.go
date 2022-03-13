package entity

import (
	"sync"
	"testing"
	"time"
)

func TestInventory_AddCar(t *testing.T) {

	i := NewInventory()
	i.N = []int64{}

	tests := []struct {
		name      string
		inventory *Inventory
		result    int
	}{
		{
			name:      "case 1",
			inventory: i,
			result:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.inventory.AddCar()
			if err != nil {
				t.Fatal(err)
			}
			if len(tt.inventory.N) != tt.result {
				t.Fatal("n is not match result")
			}
		})
	}
}

func TestInventory_SellCar(t *testing.T) {
	i := NewInventory()
	i.N = []int64{1647155235407628000}
	tests := []struct {
		name      string
		inventory *Inventory
		wantCar   int64
	}{
		{
			name:      "case 1",
			inventory: i,
			wantCar:   1647155235407628000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCar, err := tt.inventory.SellCar()
			if err != nil {
				t.Errorf("Inventory.SellCar() err: %v", err)
			}
			if gotCar != tt.wantCar {
				t.Errorf("Inventory.SellCar() = %v, want %v", gotCar, tt.wantCar)
			}
		})
	}
}

func TestInventory_GetN(t *testing.T) {
	i := NewInventory()
	i.N = []int64{1647155235407628000}
	tests := []struct {
		name      string
		inventory *Inventory
		wantN     int
	}{
		{
			name:      "case 1",
			inventory: i,
			wantN:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := tt.inventory.GetN(); gotN != tt.wantN {
				t.Errorf("Inventory.GetN() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestInventory_CalR(t *testing.T) {
	i := NewInventory()
	var m sync.Map
	m.Store(time.Now().UnixNano(), 1231231232)
	i.S = m
	tests := []struct {
		name      string
		inventory *Inventory
		wantN     int
	}{
		{
			name:      "case 1",
			inventory: i,
			wantN:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.inventory.CalR()
			if tt.inventory.R != tt.wantN {
				t.Errorf("Inventory.CalR() = %v, want %v", tt.inventory.R, tt.wantN)
			}
		})
	}
}
