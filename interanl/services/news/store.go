package news

import (
	"context"

	"github.com/google/uuid"
)

// Storer - Storer interface.
type Storer interface {
	// Create news from post request body
	Create(context.Context, PostReqBody) (PostReqBody, error)
	// FindByID news by its ID.
	FindByID(context.Context, uuid.UUID) (PostReqBody, error)
	// FindAll returns all news in the store
	FindAll() ([]PostReqBody, error)
	// DeleteByID deletes a news item by its ID
	DeleteByID(context.Context, uuid.UUID) error
	// UpdateByID updates a news resource by its ID
	UpdateByID(context.Context, PostReqBody) error
}

// Store -  Store type structure.
type Store struct{}

// Create - implements NewsStorer.
func (n *Store) Create(context.Context, PostReqBody) (PostReqBody, error) {
	panic("unimplemented")
}

// DeleteByID - implements NewsStorer.
func (n *Store) DeleteByID(context.Context, uuid.UUID) error {
	panic("unimplemented")
}

// FindAll - implements NewsStorer.
func (n *Store) FindAll() ([]PostReqBody, error) {
	panic("unimplemented")
}

// FindByID - implements NewsStorer.
func (n *Store) FindByID(context.Context, uuid.UUID) (PostReqBody, error) {
	panic("unimplemented")
}

// UpdateByID - implements NewsStorer.
func (n *Store) UpdateByID(context.Context, PostReqBody) error {
	panic("unimplemented")
}

var _ Storer = &Store{}

// NewNewsStore - store constructor.
func NewNewsStore() *Store {
	return &Store{}
}
