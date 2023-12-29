//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/entity"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/event"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/infra/database"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/infra/web"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/usecase"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)
var setGetOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)
var setGetOrdersListRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)
var setGetOrderEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)
var setGetOrdersListEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)
var setGetOrderEvent = wire.NewSet(
	event.NewGetOrder,
	wire.Bind(new(events.EventInterface), new(*event.GetOrder)),
)
var setGetOrdersListEvent = wire.NewSet(
	event.NewGetOrdersList,
	wire.Bind(new(events.EventInterface), new(*event.GetOrdersList)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}
func NewGetOrdersListUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.GetOrdersListUseCase {
	wire.Build(
		setGetOrdersListRepositoryDependency,
		setGetOrdersListEvent,
		usecase.NewGetOrdersListUseCase,
	)
	return &usecase.GetOrdersListUseCase{}
}
func NewGetOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.GetOrderUseCase {
	wire.Build(
		setGetOrderRepositoryDependency,
		setGetOrderEvent,
		usecase.NewGetOrderUseCase,
	)
	return &usecase.GetOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
func NewWebGetOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebGetOrderHandler {
	wire.Build(
		setGetOrderRepositoryDependency,
		setGetOrderEvent,
		web.NewWebGetOrderHandler,
	)
	return &web.WebGetOrderHandler{}
}
func NewWebGetOrdersListHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebGetOrdersListHandler {
	wire.Build(
		setGetOrdersListRepositoryDependency,
		setGetOrdersListEvent,
		web.NewWebGetOrdersListHandler,
	)
	return &web.WebGetOrdersListHandler{}
}
