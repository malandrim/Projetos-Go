package service

import (
	"context"

	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/infra/grpc/pb"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/usecase"
)

type CreateOrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
}
type GetOrderService struct {
	pb.UnimplementedOrderServiceServer
	GetOrderUseCase usecase.GetOrderUseCase
}
type GetOrdersListService struct {
	pb.UnimplementedOrderServiceServer
	GetOrdersListUseCase usecase.GetOrdersListUseCase
}

func NewCreateOrderService(createOrderUseCase usecase.CreateOrderUseCase) *CreateOrderService {
	return &CreateOrderService{
		CreateOrderUseCase: createOrderUseCase,
	}
}
func NewGetOrderService(getOrderUseCase usecase.GetOrderUseCase) *GetOrderService {
	return &GetOrderService{
		GetOrderUseCase: getOrderUseCase,
	}
}
func NewGetOrdersListService(getOrdersListUseCase usecase.GetOrdersListUseCase) *GetOrdersListService {
	return &GetOrdersListService{
		GetOrdersListUseCase: getOrdersListUseCase,
	}
}

func (s *CreateOrderService) CreateOrder(ctx context.Context, in *pb.OrderRequest) (*pb.OrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *GetOrderService) GetOrder(ctx context.Context, in *pb.OrderGetRequest) (*pb.OrderResponse, error) {

	order, err := s.GetOrderUseCase.Execute(in.Id)
	if err != nil {
		return nil, err
	}

	orderResponse := &pb.OrderResponse{
		Id:         order.ID,
		Price:      float32(order.Price),
		Tax:        float32(order.Tax),
		FinalPrice: float32(order.FinalPrice),
	}

	return orderResponse, nil
}

func (s *GetOrdersListService) GetOrdersList(ctx context.Context, in *pb.Blank) (*pb.OrderList, error) {
	orders, err := s.GetOrdersListUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var ordersResponse []*pb.OrderResponse

	for _, order := range orders {
		orderResponse := &pb.OrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}

		ordersResponse = append(ordersResponse, orderResponse)
	}

	return &pb.OrderList{Orders: ordersResponse}, nil
}
