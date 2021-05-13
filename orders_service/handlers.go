package main

import (
	proto "./proto"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var db, err = sqlx.Connect("mysql", "root:@tcp(127.0.0.1:3306)/go_orders")

type server struct {
	proto.UnimplementedOrdersServiceServer
}

type Order struct {
	Uuid        int64  `json:"uuid"`
	Book_uuid   int64  `json:"book_uuid"`
	Description string `json:"description"`
	Created_at  string `json:"created_at"`
}

func main() {
	handleDatabase()
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":9877")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	proto.RegisterOrdersServiceServer(srv, &server{})
	reflection.Register(srv)

	defer db.Close()
	if e := srv.Serve(listener); e != nil {
		log.Fatal(e)
	}
}

// CreateOrder godoc
// @Summary Created an order
// @Description allows to buy a book and return the information about it
// @ID create-order
// @Accept  json
// @Produce  json
// @Param book_uuid body int true "Book uuid"
// @Param description body string true "Author name"
// @Success 200 {object} bool
// @Router /create-order [post]
func (s *server) CreateOrder(ctx context.Context, request *proto.CreateOrderRequest) (*proto.Response, error) {
	bookUuid, description := request.GetBookUuid(), request.GetDescription()

	insertQuery := `INSERT INTO orders(book_uuid, description) VALUES(?,?)`

	_, err := db.Exec(insertQuery, bookUuid, description)

	if err != nil {
		return &proto.Response{Success: false}, err
	}

	return &proto.Response{Success: true}, nil
}

// GetOrderData godoc
// @Summary Get the specific order data
// @Description shows the basic information about the specific order
// @ID get-order-data
// @Accept  json
// @Produce  json
// @Param order_uuid path int true "Order uuid"
// @Success 200 {object} object
// @Router /show/{order_uuid} [get]
func (s *server) GetOrderData(ctx context.Context, request *proto.OrderId) (*proto.OrderData, error) {
	orderUuid := request.GetOrderUuid()
	order := Order{}

	selectQuery := `SELECT uuid, book_uuid, description, created_at FROM orders WHERE uuid = ?`

	err := db.Get(&order, selectQuery, orderUuid)

	if err != nil {
		return &proto.OrderData{
			OrderUuid:   0,
			BookUuid:    0,
			Description: "",
			CreatedAt:   "",
		}, err
	}

	return &proto.OrderData{
		OrderUuid:   order.Uuid,
		BookUuid:    order.Book_uuid,
		Description: order.Description,
		CreatedAt:   order.Created_at,
	}, nil
}

func handleDatabase() {
	ordersQuery := `CREATE TABLE IF NOT EXISTS orders (
        uuid INT PRIMARY KEY AUTO_INCREMENT,
        book_uuid INT NOT NULL,
        description TEXT,
		created_at datetime default CURRENT_TIMESTAMP,
        updated_at datetime default CURRENT_TIMESTAMP,
        deleted_at datetime
    );`

	db.MustExec(ordersQuery)
}
