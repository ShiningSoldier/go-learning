package main

import (
	proto "../proto"
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
// @Consume application/x-www-form-urlencoded
// @Produce  json
// @Param book_uuid formData int true "The uuid of the book you want to order"
// @Param description formData string true "The description of your order"
// @Success 200 {object} main.Order
// @Router /create-order [post]
func (s *server) CreateOrder(ctx context.Context, request *proto.CreateOrderRequest) (*proto.Order, error) {
	bookUuid, description := request.GetBookUuid(), request.GetDescription()
	insertQuery := `INSERT INTO orders(book_uuid, description) VALUES(?,?)`

	row, err := db.Exec(insertQuery, bookUuid, description)
	if err != nil {
		return &proto.Order{}, err
	}

	lastInsertedId, err := row.LastInsertId()
	if err != nil {
		return &proto.Order{}, err
	}

	return &proto.Order{
		OrderUuid:   lastInsertedId,
		BookUuid:    bookUuid,
		Description: description,
	}, nil
}

// Paginate godoc
// @Summary Show order by page number
// @Description allows to show all pages by page number
// @ID paginate
// @Accept  json
// @Produce  json
// @Param page_number path int true "Page number"
// @Success 200 {object} main.Order
// @Router /paginate/{page_number} [get]
func (s *server) Paginate(ctx context.Context, request *proto.PageNumber) (*proto.Orders, error) {
	pageNumber := request.GetPageNumber()
	offset := (pageNumber - 1) * 10
	orders := []Order{}
	response := []*proto.Order{}

	selectQuery := `SELECT uuid, book_uuid, description
    FROM orders
    WHERE deleted_at IS NULL LIMIT 10 OFFSET ?`

	err := db.Select(&orders, selectQuery, offset)
	if err != nil {
		return &proto.Orders{}, err
	}

	for _, item := range orders {
		ri := &proto.Order{
			OrderUuid:   item.Uuid,
			BookUuid:    item.Book_uuid,
			Description: item.Description,
		}

		response = append(response, ri)
	}

	return &proto.Orders{Order: response}, nil
}

// GetOrderData godoc
// @Summary Get the specific order data
// @Description shows the basic information about the specific order
// @ID get-order-data
// @Accept  json
// @Produce  json
// @Param order_uuid path int true "Order uuid"
// @Success 200 {object} main.Order
// @Router /show/{order_uuid} [get]
func (s *server) GetOrderData(ctx context.Context, request *proto.OrderId) (*proto.Order, error) {
	orderUuid := request.GetOrderUuid()
	order := Order{}

	selectQuery := `SELECT uuid, book_uuid, description FROM orders WHERE uuid = ?`

	err := db.Get(&order, selectQuery, orderUuid)

	if err != nil {
		return &proto.Order{}, err
	}

	return &proto.Order{
		OrderUuid:   order.Uuid,
		BookUuid:    order.Book_uuid,
		Description: order.Description,
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
