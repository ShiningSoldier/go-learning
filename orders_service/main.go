package main

import (
	_ "./docs"
	proto "./proto"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

// @title Orders service api
// @version 1.0
// @description Allows to create orders

// @host localhost:8081
// @BasePath /

func main() {
	conn, err := grpc.Dial("localhost:9877", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	conn2, err := grpc.Dial("localhost:9876", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewOrdersServiceClient(conn)
	client2 := proto.NewBooksServiceClient(conn2)
	g := gin.Default()

	g.POST("/create", func(ctx *gin.Context) {
		bookUuid, err := strconv.ParseUint(ctx.PostForm("book_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param book_uuid"})
		}
		description := ctx.PostForm("description")

		req := &proto.CreateOrderRequest{BookUuid: int64(bookUuid), Description: description}

		_, err = client.CreateOrder(ctx, req)

		if err == nil {
			bookReq := &proto.BookId{BookUuid: int64(bookUuid)}
			bookResponse, err := client2.ShowBook(ctx, bookReq)
			if err == nil {
				ctx.JSON(http.StatusOK, gin.H{
					"result": fmt.Sprint(bookResponse.Result),
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/show/:order_uuid", func(ctx *gin.Context) {
		orderUuid, err := strconv.ParseUint(ctx.Param("order_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param order_uuid"})
		}

		req := &proto.OrderId{OrderUuid: int64(orderUuid)}

		if response, err := client.GetOrderData(ctx, req); err == nil {
			result := [4]string{
				strconv.FormatInt(response.OrderUuid, 10),
				strconv.FormatInt(response.BookUuid, 10),
				response.Description,
				response.CreatedAt,
			}

			ctx.JSON(http.StatusOK, gin.H{
				"result": result,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = g.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
