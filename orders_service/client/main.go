package main

import (
	proto "../proto"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

func main() {
	conn, err := grpc.Dial("localhost:9876", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewOrdersServiceClient(conn)
	g := gin.Default()

	g.POST("/create", func(ctx *gin.Context) {
		bookUuid, err := strconv.ParseUint(ctx.PostForm("book_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param author_uuid"})
		}
		description := ctx.PostForm("description")

		req := &proto.CreateOrderRequest{BookUuid: int64(bookUuid), Description: description}

		if response, err := client.CreateOrder(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
}
