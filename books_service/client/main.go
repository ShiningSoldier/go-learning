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

	client := proto.NewBooksServiceClient(conn)

	g := gin.Default()

	g.POST("/add", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		categoryId := ctx.PostForm("category_id")
		authorId, err := strconv.ParseUint(ctx.PostForm("author_id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param author_id"})
		}

		req := &proto.AddBookRequest{BookName: name, CategoryId: string(categoryId), AuthorId: int64(authorId)}
		if response, err := client.AddBook(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Success),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.POST("/add-category", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		parentId, err := strconv.ParseUint(ctx.PostForm("parent_id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param parent_id"})
		}

		req := &proto.AddCategoryRequest{Name: name, ParentUuid: int64(parentId)}
		if response, err := client.AddCategory(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Success),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.POST("/add-author", func(ctx *gin.Context) {
		name := ctx.PostForm("name")

		req := &proto.AddAuthorRequest{Name: name}
		if response, err := client.AddAuthor(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Success),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.POST("/update", func(ctx *gin.Context) {
		bookUuid, err := strconv.ParseInt(ctx.PostForm("book_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param book_uuid"})
		}
		name := ctx.PostForm("name")
		categoryId := ctx.PostForm("category_id")
		authorId, err := strconv.ParseUint(ctx.PostForm("author_id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param author_id"})
		}

		req := &proto.UpdateBookRequest{BookUuid: bookUuid, BookName: name, CategoryId: string(categoryId), AuthorId: int64(authorId)}
		if response, err := client.UpdateBook(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Success),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/delete/:book_uuid", func(ctx *gin.Context) {
		bookUuid, err := strconv.ParseUint(ctx.Param("book_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param book_uuid"})
		}

		req := &proto.BookId{BookUuid: int64(bookUuid)}

		if response, err := client.DeleteBook(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Success),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/show/:book_uuid", func(ctx *gin.Context) {
		bookUuid, err := strconv.ParseUint(ctx.Param("book_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param book_uuid"})
		}

		req := &proto.BookId{BookUuid: int64(bookUuid)}

		if response, err := client.ShowBook(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
