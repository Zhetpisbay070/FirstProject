package server

import (
	"awesomeProject1/Internal/entity"
	"awesomeProject1/Internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	service service.OrderService
	router  gin.Engine
}

func (s *Server) Run() {
	s.router.POST("/create", s.CreateOrder)
}

func (s *Server) CreateOrder(ctx *gin.Context) {
	var req CreateOrderRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	order, err := s.service.CreatedOrder(ctx, &entity.CreateOrderRequest{
		UserID:       req.UserID,
		Products:     req.Products,
		Price:        req.Price,
		DeliveryType: entity.DType(req.DeliveryType),
		AddressID:    req.AddressID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	orderDTO := Order{
		ID:               order.ID,
		UserID:           order.UserID,
		ProductIDs:       order.ProductIDs,
		CreatedAt:        order.CreatedAt,
		UpdatedAt:        order.UpdatedAt,
		DeliveryDeadLine: order.DeliveryDeadLine,
		Price:            order.Price,
		DeliveryType:     string(order.DeliveryType),
		Address:          order.Address,
		OrderStatus:      string(order.OrderStatus),
	}

	ctx.JSON(http.StatusOK, orderDTO)
}
