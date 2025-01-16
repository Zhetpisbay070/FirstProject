package service

import (
	"awesomeProject1/Internal/entity"
	"awesomeProject1/Internal/repository"
	"context"
	_ "errors"
	_ "fmt"
	"time"

	"github.com/google/uuid"
)

type OrderService interface {
	CreatedOrder(ctx context.Context, req *entity.CreateOrderRequest) (*entity.Order, error)
	UpdateOrderStatus(ctx context.Context, orderStatus entity.OrderStatus, orderID string, cancel bool) error
	GetOrders(ctx context.Context, req1 *entity.GetOrders) ([]entity.Order, error)
}

type service struct {
	repo repository.DB
}

func (s *service) CreateOrder(ctx context.Context, req *entity.CreateOrderRequest) (*entity.Order, error) {
	for _, p := range req.Products {
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
		UserID:       req.UserID,
		ProductIDs:   req.Products,
		CreatedAt:    now,
		UpdatedAt:    now,
		Price:        req.Price,
		DeliveryType: req.DeliveryType,
		Address:      req.AddressID,
		OrderStatus:  entity.Created,
	}

	err := s.repo.CreateOrder(ctx, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *service) UpdateOrderStatus(ctx context.Context, userID, orderID string, cancel bool) error {
	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {

		return err
	}

	if cancel {
		if order.OrderStatus == entity.Delivery || order.OrderStatus == entity.Done {
			return entity.OrderCannotBeCancelled
		}

		order.OrderStatus = entity.Cancelled
	}

	err = s.repo.UpdateOrder(ctx, order)
	if err != nil {
		return err
	}

	var statusMove = map[entity.OrderStatus]entity.OrderStatus{
		entity.Created:   entity.Paid,
		entity.Paid:      entity.Collect,
		entity.Collect:   entity.Collected,
		entity.Collected: entity.Delivery,
		entity.Delivery:  entity.Done,
	}

	nextStatus, exists := statusMove[order.OrderStatus]
	if !exists {
		return entity.InvalidStatus
	}

	order.OrderStatus = nextStatus

	err = s.repo.UpdateOrder(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
func (s *service) GetOrders(ctx context.Context, req *entity.GetOrders) ([]entity.Order, error) {
	orders, err := s.repo.GetOrders(ctx, req)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
