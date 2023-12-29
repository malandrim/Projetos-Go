package usecase

import (
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/entity"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/pkg/events"
)

type GetOrdersListUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	GetOrdersList   events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewGetOrdersListUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	GetOrdersList events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *GetOrdersListUseCase {
	return &GetOrdersListUseCase{
		OrderRepository: OrderRepository,
		GetOrdersList:   GetOrdersList,
		EventDispatcher: EventDispatcher,
	}
}

func (g *GetOrdersListUseCase) Execute() ([]OrderOutputDTO, error) {
	var output []OrderOutputDTO

	orders, err := g.OrderRepository.FindAll()
	if err != nil {
		return []OrderOutputDTO{}, err
	}

	//loop to append all returned orders >>>
	var o OrderOutputDTO
	for _, order := range orders {
		o = OrderOutputDTO{
			order.ID,
			order.Price,
			order.Tax,
			order.FinalPrice,
		}
		output = append(output, o)
	}
	//loop to append all returned orders <<<

	return output, nil
}
