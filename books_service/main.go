package main

import (
	proto "./proto"
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
		categoryId := ctx.PostForm("category_uuid")
		authorId, err := strconv.ParseUint(ctx.PostForm("author_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param author_uuid"})
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

	g.PUT("/update", func(ctx *gin.Context) {
		bookUuid, err := strconv.ParseInt(ctx.PostForm("book_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param book_uuid"})
		}
		name := ctx.PostForm("name")
		categoryId := ctx.PostForm("category_uuid")
		authorId, err := strconv.ParseUint(ctx.PostForm("author_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param author_id"})
		}

		req := &proto.UpdateBookRequest{BookUuid: bookUuid, BookName: name, CategoryId: categoryId, AuthorId: int64(authorId)}
		if response, err := client.UpdateBook(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Success),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.DELETE("/delete/:book_uuid", func(ctx *gin.Context) {
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

	g.GET("/show-author/:author_uuid", func(ctx *gin.Context) {
		authorUuid, err := strconv.ParseUint(ctx.Param("author_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param author_uuid"})
		}

		req := &proto.AuthorId{AuthorUuid: int64(authorUuid)}

		if response, err := client.ShowAuthor(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/show-category/:category_uuid", func(ctx *gin.Context) {
		bookUuid, err := strconv.ParseUint(ctx.Param("category_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param category_uuid"})
		}

		req := &proto.CategoryId{CategoryUuid: int64(bookUuid)}

		if response, err := client.ShowCategory(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/filter-by-author/:author_uuid", func(ctx *gin.Context) {
		authorUuid, err := strconv.ParseUint(ctx.Param("author_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param author_uuid"})
		}

		req := &proto.AuthorId{AuthorUuid: int64(authorUuid)}

		if response, err := client.FilterByAuthor(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/filter-by-category/:category_uuid", func(ctx *gin.Context) {
		categoryUuid, err := strconv.ParseUint(ctx.Param("category_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param author_uuid"})
		}

		req := &proto.CategoryId{CategoryUuid: int64(categoryUuid)}

		if response, err := client.FilterByCategory(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/paginate/:page_number", func(ctx *gin.Context) {
		pageNumber, err := strconv.ParseUint(ctx.Param("page_number"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param page_number"})
		}

		req := &proto.PageNumber{PageNumber: int64(pageNumber)}

		if response, err := client.Paginate(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.DELETE("/delete-author/:author_uuid", func(ctx *gin.Context) {
		authorUuid, err := strconv.ParseUint(ctx.Param("author_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param author_uuid"})
		}

		req := &proto.AuthorId{AuthorUuid: int64(authorUuid)}

		if response, err := client.DeleteAuthor(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Success),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.DELETE("/delete-category/:category_uuid", func(ctx *gin.Context) {
		categoryUuid, err := strconv.ParseUint(ctx.Param("category_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param category_uuid"})
		}

		req := &proto.CategoryId{CategoryUuid: int64(categoryUuid)}

		if response, err := client.DeleteCategory(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Success),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.PUT("/update-author", func(ctx *gin.Context) {
		authorUuid, err := strconv.ParseInt(ctx.PostForm("author_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param author_uuid"})
		}
		name := ctx.PostForm("name")

		req := &proto.UpdateAuthorRequest{AuthorUuid: authorUuid, AuthorName: name}
		if response, err := client.UpdateAuthor(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Success),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.PUT("/update-category", func(ctx *gin.Context) {
		categoryUuid, err := strconv.ParseInt(ctx.PostForm("category_uuid"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param author_uuid"})
		}
		name := ctx.PostForm("name")
		parentId, err := strconv.ParseUint(ctx.PostForm("parent_id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param parent_id"})
		}

		req := &proto.UpdateCategoryRequest{CategoryUuid: categoryUuid, CategoryName: name, ParentUuid: int64(parentId)}
		if response, err := client.UpdateCategory(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Success),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
