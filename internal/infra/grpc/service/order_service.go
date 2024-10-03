package service

import (
	"Clean_Architecture/internal/infra/grpc/pb"
	"Clean_Architecture/internal/usecase"
	"context"
	"strconv"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase  usecase.CreateOrderUseCase
	FindAllOrderUseCase usecase.FindAllOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, findAllOrderUseCase usecase.FindAllOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase:  createOrderUseCase,
		FindAllOrderUseCase: findAllOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}

	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.FindAllOrderRequest) (*pb.FindAllOrdersResponse, error) {
	page, err := strconv.Atoi(in.Page)
	if err != nil {
		return nil, err
	}

	// Mapeia a entrada para o DTO de uso
	dto := usecase.FindAllOrderInputDTO{
		Page:  strconv.Itoa(page),
		Limit: in.Limit,
		Sort:  in.Sort,
	}

	// Chama o caso de uso
	orders, err := s.FindAllOrderUseCase.FindAll(dto)
	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, nil // Certifique-se de que está retornando nil, e não um objeto vazio
	}

	var response pb.FindAllOrdersResponse // Isso deve ser FindAllOrdersResponse
	for _, order := range orders {
		response.Orders = append(response.Orders, &pb.FindAllOrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return &response, nil
}
