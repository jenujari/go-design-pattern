package main

import "testing"

func TestGetSportsFactory(t *testing.T) {
	tests := []struct {
		brand string
		logo  string
	}{
		{"adidas", "adidas"},
		{"nike", "nike"},
	}

	for _, tt := range tests {
		f, _ := GetSportsFactory(tt.brand)
		if f == nil {
			t.Fatalf("factory for %s returned nil", tt.brand)
		}
		shoe := f.makeShoe()
		if shoe.getLogo() != tt.logo {
			t.Errorf("%s shoe logo expected %s got %s", tt.brand, tt.logo, shoe.getLogo())
		}
		if shoe.getSize() != 14 {
			t.Errorf("%s shoe size expected 14 got %d", tt.brand, shoe.getSize())
		}

		shirt := f.makeShirt()
		if shirt.getLogo() != tt.logo {
			t.Errorf("%s shirt logo expected %s got %s", tt.brand, tt.logo, shirt.getLogo())
		}
		if shirt.getSize() != 14 {
			t.Errorf("%s shirt size expected 14 got %d", tt.brand, shirt.getSize())
		}
	}
}

func TestGetSportsFactoryUnknown(t *testing.T) {
	_, err := GetSportsFactory("unknown")
	if err == nil {
		t.Fatalf("expected error for unknown brand, got nil")
	}
}
