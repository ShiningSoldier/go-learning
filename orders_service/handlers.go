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
	Book_id     int    `json:"book_id"`
	Description string `json:"description"`
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
// @Summary Create a new order
// @Description Create an order by book ID. Also the description can be provided
// @ID create-order
// @Accept json
// @Produce json
// @Param book_uuid body int64 true
// @Param description body string false

func (s *server) CreateOrder(ctx context.Context, request *proto.CreateOrderRequest) (*proto.BookData, error) {
	bookUuid, description := request.GetBookUuid(), request.GetDescription()

	insertQuery := `INSERT INTO orders(book_uuid, description) VALUES(?,?)`

	_, err := db.Exec(insertQuery, bookUuid, description)

	if err != nil {
		return &proto.BookData{Result: ""}, err
	}

	return &proto.BookData{Result: "Test"}, nil

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
