package repository

import (
	"awesomeProject1/Internal/entity"
	"context"
)

type DB interface {
	CreateOrder(ctx context.Context, order *entity.Order) error
	GetOrderByID(ctx context.Context, id string) (*entity.Order, error)
	ProductExist(ctx context.Context, productID string) (bool, error)
	UpdateOrder(ctx context.Context, order *entity.Order) error
}
