package service

import (
	"awesomeProject1/Internal/entity"
	"awesomeProject1/Internal/repository"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type OrderService interface {
	CreatedOrder(ctx context.Context, userID string, products []string, price float64, DeliveryTyoe entity.DType, AddressID string) *entity.Order
	UpdateOrderStatus(ctx context.Context, userID, orderID string) error
	GetOrders(ctx context.Context, userID string, limit, offset int, asc bool) ([]entity.Order, error)
}

type service struct {
	repo repository.DB
}

func (s *service) CreateOrder(ctx context.Context, userID string, products []string, price float64, deliveryType entity.DType, addressID string) (*entity.Order, error) {
	for _, p := range products {
		ok, err := s.repo.ProductExist(ctx, p)
		if err != nil {
			return nil, err
		}

		if !ok {
			return nil, entity.ProductDoesNotExistError
		}
	}

	now := time.Now()

	order := entity.Order{
		ID:           uuid.New().String(),
		UserID:       userID,
		ProductIDs:   products,
		CreatedAt:    now,
		UpdatedAt:    now,
		Price:        price,
		DeliveryType: deliveryType,
		Address:      addressID,
		OrderStatus:  entity.Created,
	}

	err := s.repo.CreateOrder(ctx, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *service) UpdateOrderStatus(ctx context.Context, userID, orderID string) error {
	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	switch order.OrderStatus {
	case entity.Created:
		order.OrderStatus = entity.Paid
	case entity.Paid:
		order.OrderStatus = entity.Collect
	case entity.Collect:
		order.OrderStatus = entity.Collected
	case entity.Collected:
		order.OrderStatus = entity.Delivery
	case entity.Delivery:
		order.OrderStatus = entity.Done
	default:
		fmt.Println("error with status")
	}

	if order.OrderStatus == entity.Cancelled && (order.OrderStatus == entity.Delivery || order.OrderStatus == entity.Done) {
		return entity.ErrOrderCannotBeCancelled
	}

	err = s.repo.UpdateOrder(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetOrders(ctx context.Context, userID string, limit, offset int, asc bool) ([]entity.Order, error) {

}
