package usecase

import (
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/entity"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/pkg/events"
)

type GetOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	GetOrder        events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewGetOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	GetOrder events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *GetOrderUseCase {
	return &GetOrderUseCase{
		OrderRepository: OrderRepository,
		GetOrder:        GetOrder,
		EventDispatcher: EventDispatcher,
	}
}

func (g *GetOrderUseCase) Execute(id string) (*OrderOutputDTO, error) {
	var order *entity.Order

	order, err := g.OrderRepository.FindByID(id)
	if err != nil {
		return &OrderOutputDTO{}, err
	}

	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil

}
