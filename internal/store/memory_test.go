package store

import (
	"errors"
	"testing"

	"github.com/mrckurz/CI-CD-MCM/internal/model"
)

func TestCreateAndGet(t *testing.T) {
	s := NewMemoryStore()
	p := model.Product{Name: "Test Product", Price: 10.99}

	// Create
	created := s.Create(p)
	if created.ID == 0 {
		t.Fatal("expected assigned ID to be non-zero")
	}

	// Get
	found, err := s.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error getting product: %v", err)
	}

	if found.Name != p.Name || found.Price != p.Price {
		t.Errorf("expected %+v, got %+v", p, found)
	}
}

func TestUpdateProduct(t *testing.T) {
	s := NewMemoryStore()
	// Initial create
	p := s.Create(model.Product{Name: "Old Name", Price: 5.00})

	// Update
	p.Name = "New Name"
	p.Price = 15.00
	updated, err := s.Update(p.ID, p)
	if err != nil {
		t.Fatalf("unexpected error updating product: %v", err)
	}

	// Verify return value
	if updated.Name != "New Name" {
		t.Errorf("expected name 'New Name', got '%s'", updated.Name)
	}

	// Verify persistence
	stored, _ := s.GetByID(p.ID)
	if stored.Price != 15.00 {
		t.Errorf("expected price 15.00, got %f", stored.Price)
	}
}

func TestDeleteProduct(t *testing.T) {
	s := NewMemoryStore()
	p := s.Create(model.Product{Name: "To Delete", Price: 1.00})

	// Delete
	err := s.Delete(p.ID)
	if err != nil {
		t.Fatalf("unexpected error deleting product: %v", err)
	}

	// Verify GetByID returns error
	_, err = s.GetByID(p.ID)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound after deletion, got %v", err)
	}
}

func TestGetByIDNotFound(t *testing.T) {
	s := NewMemoryStore()

	// Table-driven test
	tests := []struct {
		name string
		id   int
	}{
		{"Empty Store", 1},
		{"Non-existent ID", 999},
		{"Negative ID", -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetByID(tt.id)
			if !errors.Is(err, ErrNotFound) {
				t.Errorf("%s: expected ErrNotFound, got %v", tt.name, err)
			}
		})
	}
}

func TestGetAllEmpty(t *testing.T) {
	s := NewMemoryStore()
	products := s.GetAll()
	if len(products) != 0 {
		t.Errorf("expected 0 products, got %d", len(products))
	}
}

func TestDeleteNonExistent(t *testing.T) {
	s := NewMemoryStore()
	err := s.Delete(999)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}
