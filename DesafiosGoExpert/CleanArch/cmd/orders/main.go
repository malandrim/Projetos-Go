package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/graph"

	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/configs"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/event/handler"
	service "github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/infra/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/infra/grpc/pb"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/infra/web/webserver"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/pkg/events"

	"github.com/streadway/amqp"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel(configs.RMQUser, configs.RMQPassword, configs.RMQHost, configs.RMQServerPort)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)

	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webGetOrderHandler := NewWebGetOrderHandler(db, eventDispatcher)
	webGetOrdersListHandler := NewWebGetOrdersListHandler(db, eventDispatcher)

	webserver.AddHandlerPost("/order", webOrderHandler.Create)
	webserver.AddHandlerGet("/order", webGetOrderHandler.GetOrder)
	webserver.AddHandlerGet("/orders", webGetOrdersListHandler.GetOrdersList)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewCreateOrderService(*createOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)

}

func getRabbitMQChannel(user, pass, host, port string) *amqp.Channel {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port))
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
