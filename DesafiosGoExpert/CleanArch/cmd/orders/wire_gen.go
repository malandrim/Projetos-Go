// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, orderCreated, eventDispatcher)
	return createOrderUseCase
}

func NewGetOrdersListUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.GetOrdersListUseCase {
	orderRepository := database.NewOrderRepository(db)
	getOrdersList := event.NewGetOrdersList()
	getOrdersListUseCase := usecase.NewGetOrdersListUseCase(orderRepository, getOrdersList, eventDispatcher)
	return getOrdersListUseCase
}

func NewGetOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.GetOrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	getOrder := event.NewGetOrder()
	getOrderUseCase := usecase.NewGetOrderUseCase(orderRepository, getOrder, eventDispatcher)
	return getOrderUseCase
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	webOrderHandler := web.NewWebOrderHandler(eventDispatcher, orderRepository, orderCreated)
	return webOrderHandler
}

func NewWebGetOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebGetOrderHandler {
	orderRepository := database.NewOrderRepository(db)
	getOrder := event.NewGetOrder()
	webGetOrderHandler := web.NewWebGetOrderHandler(eventDispatcher, orderRepository, getOrder)
	return webGetOrderHandler
}

func NewWebGetOrdersListHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebGetOrdersListHandler {
	orderRepository := database.NewOrderRepository(db)
	getOrdersList := event.NewGetOrdersList()
	webGetOrdersListHandler := web.NewWebGetOrdersListHandler(eventDispatcher, orderRepository, getOrdersList)
	return webGetOrdersListHandler
}

// wire.go:

var setOrderRepositoryDependency = wire.NewSet(database.NewOrderRepository, wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)))

var setGetOrderRepositoryDependency = wire.NewSet(database.NewOrderRepository, wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)))

var setGetOrdersListRepositoryDependency = wire.NewSet(database.NewOrderRepository, wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)))

var setEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)), wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setGetOrderEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)), wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setGetOrdersListEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)), wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setOrderCreatedEvent = wire.NewSet(event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)))

var setGetOrderEvent = wire.NewSet(event.NewGetOrder, wire.Bind(new(events.EventInterface), new(*event.GetOrder)))

var setGetOrdersListEvent = wire.NewSet(event.NewGetOrdersList, wire.Bind(new(events.EventInterface), new(*event.GetOrdersList)))
