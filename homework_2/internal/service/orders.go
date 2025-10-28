package service

import (
	"context"

	"orders-microservice/internal/models"
	"orders-microservice/pkg/api/test"
)

type Repository interface {
	Create(ctx context.Context, order models.Order) (string, error)
	GetAll(ctx context.Context) ([]*test.Order, error)
	GetOne(ctx context.Context, id string) (models.Order, error)
	Update(ctx context.Context, id string, newOrder models.Order) error
	Delete(ctx context.Context, id string) error
}

type Service struct {
	test.UnimplementedOrderServiceServer

	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateOrder(ctx context.Context, req *test.CreateOrderRequest) (*test.CreateOrderResponse, error) {
	repoOrder := models.Order{
		Item:     req.GetItem(),
		Quantity: req.GetQuantity(),
	}

	id, err := s.repo.Create(ctx, repoOrder)
	if err != nil {
		return nil, err
	}

	return &test.CreateOrderResponse{Id: id}, nil
}

func (s *Service) GetOrder(ctx context.Context, req *test.GetOrderRequest) (*test.GetOrderResponse, error) {
	orderData, err := s.repo.GetOne(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &test.GetOrderResponse{Order: &test.Order{
		Id:       req.GetId(),
		Item:     orderData.Item,
		Quantity: orderData.Quantity,
	}}, nil
}

func (s *Service) UpdateOrder(ctx context.Context, req *test.UpdateOrderRequest) (*test.UpdateOrderResponse, error) {
	err := s.repo.Update(ctx, req.Id, models.Order{Item: req.GetItem(), Quantity: req.GetQuantity()})
	if err != nil {
		return nil, err
	}

	return &test.UpdateOrderResponse{
		Order: &test.Order{
			Id:       req.GetId(),
			Item:     req.GetItem(),
			Quantity: req.GetQuantity(),
		},
	}, nil
}

func (s *Service) DeleteOrder(ctx context.Context, req *test.DeleteOrderRequest) (*test.DeleteOrderResponse, error) {
	err := s.repo.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &test.DeleteOrderResponse{Success: true}, nil
}

func (s *Service) ListOrders(ctx context.Context, _ *test.ListOrdersRequest) (*test.ListOrdersResponse, error) {
	orders, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return &test.ListOrdersResponse{Orders: orders}, nil
}
