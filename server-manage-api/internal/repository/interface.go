package repository

import (
	"context"
)

type ServerRepository interface {
	AutoMigrate() error
	Create(ctx context.Context, rowServer *RowServer) (*RowServer, error)
	List(ctx context.Context) ([]*RowServer, error)
	GetByID(ctx context.Context, id string) (*RowServer, error)
	DeleteByID(ctx context.Context, id string) error
}

type RowServer struct {
	ID     string
	Name   string
	Type   string
	Status string
}

func (RowServer) TableName() string {
	return "servers"
}
