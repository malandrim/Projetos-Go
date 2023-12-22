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

func (c *GetOrderUseCase) Execute(OrderId string) (OrderOutputDTO, error) {

	var order entity.Order

	order = c.GetOrder.GetOrderById(OrderId)

	c.GetOrder.SetPayload(order)
	c.EventDispatcher.Dispatch(c.GetOrder)

	return OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
