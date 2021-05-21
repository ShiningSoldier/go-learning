package main

import (
	_ "../docs"
	proto "../proto"
	"context"
	_ "database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strings"
	"time"
)

var db, err = sqlx.Connect("mysql", "root:@tcp(127.0.0.1:3306)/go_books")

type server struct {
	proto.UnimplementedBooksServiceServer
}

type Book struct {
	Uuid        int    `json:"uuid"`
	Name        string `json:"name"`
	Author_uuid int    `json:"author_uuid"`
	Author      `db:"authors"`
}

type Author struct {
	Uuid int    `json:"uuid"`
	Name string `json:"name"`
}

type BooksCategories struct {
	Category_uuid int `json:"category_uuid"`
	Category      `db:"categories"`
}

type Category struct {
	Uuid        int    `json:"uuid"`
	Name        string `json:"name"`
	Parent_uuid int    `json:"parent_uuid"`
	Parent_name string `json:"parent_name"`
}

func main() {
	handleDatabase()
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":9876")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	proto.RegisterBooksServiceServer(srv, &server{})
	reflection.Register(srv)

	defer db.Close()
	if e := srv.Serve(listener); e != nil {
		log.Fatal(e)
	}
}

// AddBook godoc
// @Summary Creates a new book
// @Description create a book using the POST request
// @ID add-book
// @Consume application/x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "Book name"
// @Param category_uuid formData string true "List of category iIDs"
// @Param author_uuid formData int true "Book author ID"
// @Success 200 {object} main.Book
// @Router /add [post]
func (s *server) AddBook(ctx context.Context, request *proto.AddBookRequest) (*proto.Book, error) {
	name, category, author := request.GetBookName(), request.GetCategoryId(), request.GetAuthorId()
	categoriesSlice := strings.Split(category, ",")
	insertQuery := `INSERT INTO books(name, author_uuid) VALUES(?,?)`

	row, err := db.Exec(insertQuery, name, author)
	if err != nil {
		return &proto.Book{}, err
	}
	lastInsertedId, err := row.LastInsertId()
	if err != nil {
		return &proto.Book{}, err
	}

	err = addCategories(lastInsertedId, categoriesSlice)
	if err != nil {
		return &proto.Book{}, err
	}

	authorData, err := getAuthor(author)
	if err != nil {
		return &proto.Book{}, err
	}

	categories, err := getCategories(lastInsertedId)
	if err != nil {
		return &proto.Book{}, err
	}

	return &proto.Book{
		BookUuid:   lastInsertedId,
		Name:       name,
		Author:     authorData.Name,
		Categories: categories,
	}, nil
}

// UpdateBook godoc
// @Summary Updates a book
// @Description update a book using the PATCH request
// @ID update-book
// @Consume application/x-www-form-urlencoded
// @Produce  json
// @Param book_uuid formData int true "Book uuid"
// @Param name formData string false "Book name"
// @Param category_uuid formData string false "List of category iIDs"
// @Param author_uuid formData int false "Book author ID"
// @Success 200 {object} main.Book
// @Router /update [patch]
func (s *server) UpdateBook(ctx context.Context, request *proto.UpdateBookRequest) (*proto.Book, error) {
	bookUuid, name, category, author, currentTimestamp := request.GetBookUuid(), request.GetBookName(), request.GetCategoryId(), request.GetAuthorId(), time.Now().Format("2006-01-02 15:04:05")

	if len(name) > 0 || len(author) > 0 || len(category) > 0 {
		updateQuery := sq.Update("books").Where(sq.Eq{"uuid": bookUuid})

		if len(name) > 0 {
			updateQuery = updateQuery.Set("name", name)
		}

		if len(author) > 0 {
			updateQuery = updateQuery.Set("author_uuid", author)
		}

		if len(category) > 0 {
			categoriesSlice := strings.Split(category, ",")
			err = addCategories(bookUuid, categoriesSlice)
			if err != nil {
				return &proto.Book{}, err
			}
		}

		updateQuery = updateQuery.Set("updated_at", currentTimestamp)

		query, args, err := updateQuery.ToSql()
		if err != nil {
			return &proto.Book{}, err
		}

		_, err = db.Exec(query, args...)
		if err != nil {
			return &proto.Book{}, err
		}
	}

	bookName, authorName, categories, err := getBookData(bookUuid)
	if err != nil {
		return &proto.Book{}, err
	}

	return &proto.Book{
		BookUuid:   bookUuid,
		Name:       bookName,
		Author:     authorName,
		Categories: categories,
	}, nil
}

// UpdateAuthor godoc
// @Summary Updates an author
// @Description update an author using the PATCH request
// @ID update-author
// @Consume application/x-www-form-urlencoded
// @Produce  json
// @Param author_uuid formData int true "Author uuid"
// @Param name formData string false "Author name"
// @Success 200 {object} main.Author
// @Router /update-author [patch]
func (s *server) UpdateAuthor(ctx context.Context, request *proto.UpdateAuthorRequest) (*proto.Author, error) {
	authorUuid, name, currentTimestamp := request.GetAuthorUuid(), request.GetAuthorName(), time.Now().Format("2006-01-02 15:04:05")

	if len(name) > 0 {
		updateQuery := sq.Update("authors").Set("name", name).Set("updated_at", currentTimestamp).Where(sq.Eq{"uuid": authorUuid})
		query, args, err := updateQuery.ToSql()
		if err != nil {
			return &proto.Author{}, err
		}

		_, err = db.Exec(query, args...)
		if err != nil {
			return &proto.Author{}, err
		}
	}

	return &proto.Author{
		AuthorUuid: authorUuid,
		Name:       name,
	}, nil
}

// UpdateCategory godoc
// @Summary Updates a category
// @Description update a category using the PUT request
// @ID update-category
// @Consume application/x-www-form-urlencoded
// @Produce  json
// @Param category_uuid formData int true "Category uuid"
// @Param name formData string false "Category name"
// @Param parent_uuid formData string false "Parent id"
// @Success 200 {object} main.Category
// @Router /update-category [patch]
func (s *server) UpdateCategory(ctx context.Context, request *proto.UpdateCategoryRequest) (*proto.Category, error) {
	categoryUuid, name, parentUuid, currentTimestamp := request.GetCategoryUuid(), request.GetCategoryName(), request.GetParentUuid(), time.Now().Format("2006-01-02 15:04:05")

	if len(name) > 0 || len(parentUuid) > 0 {
		updateQuery := sq.Update("categories").Where(sq.Eq{"uuid": categoryUuid})
		if len(name) > 0 {
			updateQuery = updateQuery.Set("name", name)
		}

		if len(parentUuid) > 0 {
			updateQuery = updateQuery.Set("parent_uuid", parentUuid)
		}

		updateQuery = updateQuery.Set("updated_at", currentTimestamp)

		query, args, err := updateQuery.ToSql()
		if err != nil {
			return &proto.Category{}, err
		}

		_, err = db.Exec(query, args...)
		if err != nil {
			return &proto.Category{}, err
		}
	}

	category := Category{}
	selectQuery := `SELECT c.name, c.parent_uuid, IFNULL(c2.name, "") AS parent_name
    FROM categories c
    LEFT JOIN categories c2 ON c.parent_uuid = c2.uuid
    WHERE c.deleted_at IS NULL AND c.uuid = ?`
	err := db.Get(&category, selectQuery, categoryUuid)
	if err != nil {
		return &proto.Category{}, err
	}

	return &proto.Category{
		CategoryUuid: categoryUuid,
		ParentName:   category.Parent_name,
		Name:         name,
	}, nil
}

func addCategories(bookUuid int64, categoriesSlice []string) error {
	_, err = db.Exec("DELETE FROM books_categories WHERE book_uuid = ?", bookUuid)
	if err != nil {
		return err
	}

	for _, item := range categoriesSlice {
		_, err := db.Exec("INSERT INTO books_categories(book_uuid, category_uuid) VALUES(?,?)", bookUuid, item)
		if err != nil {
			return err
		}
	}

	return err
}

// DeleteBook godoc
// @Summary Deletes a book
// @Description delete a book using the DELETE request
// @ID delete-book
// @Accept  json
// @Produce  json
// @Param book_uuid path int true "Book uuid"
// @Success 200 {object} bool
// @Router /delete/{book_uuid} [delete]
func (s *server) DeleteBook(ctx context.Context, request *proto.BookId) (*proto.Response, error) {
	bookUuid := request.GetBookUuid()

	err := deleteEntity("books", bookUuid)

	if err != nil {
		return &proto.Response{Success: false}, err
	}

	err = deleteBookCategoriesLinking(bookUuid)
	if err != nil {
		return &proto.Response{Success: false}, err
	}

	return &proto.Response{Success: true}, nil
}

// ShowBook godoc
// @Summary Shows a book
// @Description shows the main data about a book
// @ID show-book
// @Accept  json
// @Produce  json
// @Param book_uuid path int true "Book uuid"
// @Success 200 {object} main.Book
// @Router /show/{book_uuid} [get]
func (s *server) ShowBook(ctx context.Context, request *proto.BookId) (*proto.Book, error) {
	bookUuid := request.GetBookUuid()

	bookName, authorName, categories, err := getBookData(bookUuid)

	if err != nil {
		return &proto.Book{}, err
	}

	return &proto.Book{
		BookUuid:   bookUuid,
		Name:       bookName,
		Author:     authorName,
		Categories: categories,
	}, nil
}

func getBookData(bookUuid int64) (string, string, string, error) {
	book := Book{}

	selectBookQuery := `SELECT books.name, authors.name "authors.name"
    FROM books
    INNER JOIN authors ON authors.uuid = books.author_uuid
    WHERE books.deleted_at IS NULL AND books.uuid = ?`

	err := db.Get(&book, selectBookQuery, bookUuid)
	if err != nil {
		return "", "", "", err
	}

	categories, err := getCategories(bookUuid)
	if err != nil {
		return "", "", "", err
	}

	return book.Name, book.Author.Name, categories, nil
}

func getCategories(bookUuid int64) (string, error) {
	booksCategories := []BooksCategories{}

	selectBooksCategoriesQuery := `SELECT category_uuid, categories.name "categories.name",
       categories.parent_uuid "categories.parent_uuid"
    FROM books_categories
    INNER JOIN categories ON categories.uuid = books_categories.category_uuid
    WHERE book_uuid = ?`
	err = db.Select(&booksCategories, selectBooksCategoriesQuery, bookUuid)
	if err != nil {
		return "", err

	}

	categories := ""

	for _, item := range booksCategories {
		categories = categories + item.Category.Name + "; "
	}

	return strings.TrimSpace(categories), nil
}

// AddCategory godoc
// @Summary Create a category
// @Description creates a new category
// @ID create-category
// @Consume application/x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "Category name"
// @Param parent_uuid formData string false "Parent id"
// @Success 200 {object} main.Category
// @Router /add-category [post]
func (s *server) AddCategory(ctx context.Context, request *proto.AddCategoryRequest) (*proto.Category, error) {
	name, parentUuid := request.GetName(), request.GetParentUuid()
	parentCategory := Category{}
	if len(parentUuid) == 0 {
		parentUuid = "0"
	}
	insertQuery := `INSERT INTO categories(name, parent_uuid) VALUES(?,?)`

	row, err := db.Exec(insertQuery, name, parentUuid)
	if err != nil {
		return &proto.Category{}, err
	}
	lastInsertedId, err := row.LastInsertId()
	if err != nil {
		return &proto.Category{}, err
	}

	if parentUuid != "0" {
		err := db.Get(&parentCategory, `SELECT name FROM categories WHERE uuid = ?`, parentUuid)
		if err != nil {
			return &proto.Category{}, err
		}
	}

	return &proto.Category{
		CategoryUuid: lastInsertedId,
		ParentName:   parentCategory.Name,
		Name:         name,
	}, nil
}

// AddAuthor godoc
// @Summary Create an author
// @Description creates a new author
// @ID create-author
// @Consume application/x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "Author name"
// @Success 200 {object} main.Author
// @Router /add-author [post]
func (s *server) AddAuthor(ctx context.Context, request *proto.AddAuthorRequest) (*proto.Author, error) {
	name := request.GetName()
	if len(name) == 0 {
		err := errors.New("name field is empty")
		return &proto.Author{}, err
	}
	insertQuery := `INSERT INTO authors(name) VALUES(?)`

	row, err := db.Exec(insertQuery, name)
	if err != nil {
		return &proto.Author{}, err
	}

	lastInsertedId, err := row.LastInsertId()
	if err != nil {
		return &proto.Author{}, err
	}

	return &proto.Author{
		AuthorUuid: lastInsertedId,
		Name:       name,
	}, nil
}

// ShowAuthor godoc
// @Summary Show author data
// @Description show the basic author data
// @ID show-author
// @Accept  json
// @Produce  json
// @Param author_uuid path int true "Author uuid"
// @Success 200 {object} main.Author
// @Router /show-author/{author_uuid} [get]
func (s *server) ShowAuthor(ctx context.Context, request *proto.AuthorId) (*proto.Author, error) {
	authorUuid := request.GetAuthorUuid()
	author, err := getAuthor(authorUuid)
	if err != nil {
		return &proto.Author{}, err
	}

	return &proto.Author{
		AuthorUuid: authorUuid,
		Name:       author.Name,
	}, nil
}

func getAuthor(authorUuid int64) (Author, error) {
	author := Author{}
	selectQuery := `SELECT name FROM authors WHERE deleted_at IS NULL AND uuid = ?`
	err := db.Get(&author, selectQuery, authorUuid)

	return author, err
}

// ShowCategory godoc
// @Summary Show category data
// @Description show the basic category data
// @ID show-category
// @Accept  json
// @Produce  json
// @Param category_uuid path int true "Category uuid"
// @Success 200 {object} main.Category
// @Router /show-category/{category_uuid} [get]
func (s *server) ShowCategory(ctx context.Context, request *proto.CategoryId) (*proto.Category, error) {
	categoryUuid := request.GetCategoryUuid()
	category := Category{}

	selectQuery := `SELECT c.name, c.parent_uuid, IFNULL(c2.name, "") AS parent_name
    FROM categories c
    LEFT JOIN categories c2 ON c.parent_uuid = c2.uuid
    WHERE c.deleted_at IS NULL AND c.uuid = ?`

	err := db.Get(&category, selectQuery, categoryUuid)
	if err != nil {
		return &proto.Category{}, err
	}

	return &proto.Category{
		CategoryUuid: categoryUuid,
		ParentName:   category.Parent_name,
		Name:         category.Name,
	}, nil
}

// FilterByAuthor godoc
// @Summary Shows all books by specified author
// @Description shows the basic data of books by author
// @ID filter-by-author
// @Accept  json
// @Produce  json
// @Param author_uuid path int true "Author uuid"
// @Success 200 {object} main.Book
// @Router /filter-by-author/{author_uuid} [get]
func (s *server) FilterByAuthor(ctx context.Context, request *proto.AuthorId) (*proto.Books, error) {
	authorUuid := request.GetAuthorUuid()
	books := []Book{}
	response := []*proto.Book{}

	selectQuery := `SELECT books.uuid, books.name, authors.name "authors.name"
    FROM books
    INNER JOIN authors ON authors.uuid = books.author_uuid
    WHERE books.deleted_at IS NULL AND books.author_uuid = ?`

	err := db.Select(&books, selectQuery, authorUuid)
	if err != nil {
		return &proto.Books{}, err
	}

	for _, item := range books {
		uuid := int64(item.Uuid)
		categories, err := getCategories(uuid)
		if err != nil {
			return &proto.Books{}, err
		}
		ri := &proto.Book{
			BookUuid:   uuid,
			Name:       item.Name,
			Author:     item.Author.Name,
			Categories: categories,
		}

		response = append(response, ri)
	}

	return &proto.Books{Book: response}, nil
}

// FilterByCategory godoc
// @Summary Shows all books by specified category
// @Description shows the basic data of books by category
// @ID filter-by-category
// @Accept  json
// @Produce  json
// @Param category_uuid path int true "Category uuid"
// @Success 200 {object} main.Category
// @Router /filter-by-category/{category_uuid} [get]
func (s *server) FilterByCategory(ctx context.Context, request *proto.CategoryId) (*proto.Books, error) {
	categoryUuid := request.GetCategoryUuid()
	books := []Book{}
	response := []*proto.Book{}

	selectQuery := `SELECT books.uuid, books.name, authors.name "authors.name" FROM books
    INNER JOIN authors ON authors.uuid = books.author_uuid
    INNER JOIN books_categories ON books_categories.book_uuid = books.uuid
    WHERE books.deleted_at IS NULL AND books_categories.category_uuid = ?`

	err := db.Select(&books, selectQuery, categoryUuid)
	if err != nil {
		return &proto.Books{}, err
	}

	for _, item := range books {
		uuid := int64(item.Uuid)
		categories, err := getCategories(uuid)
		if err != nil {
			return &proto.Books{}, err
		}
		ri := &proto.Book{
			BookUuid:   uuid,
			Name:       item.Name,
			Author:     item.Author.Name,
			Categories: categories,
		}

		response = append(response, ri)
	}

	return &proto.Books{Book: response}, nil
}

// Paginate godoc
// @Summary Paginate books
// @Description shows the books by pages
// @ID paginate
// @Accept  json
// @Produce  json
// @Param page_number path int true "Page number"
// @Success 200 {object} main.Book
// @Router /paginate/{page_number} [get]
func (s *server) Paginate(ctx context.Context, request *proto.PageNumber) (*proto.Books, error) {
	pageNumber := request.GetPageNumber()
	offset := (pageNumber - 1) * 10
	books := []Book{}
	response := []*proto.Book{}

	selectQuery := `SELECT books.uuid, books.name, authors.name "authors.name"
    FROM books
	INNER JOIN authors ON authors.uuid = books.author_uuid
    WHERE books.deleted_at IS NULL LIMIT 10 OFFSET ?`

	err := db.Select(&books, selectQuery, offset)
	if err != nil {
		return &proto.Books{}, err
	}

	for _, item := range books {
		uuid := int64(item.Uuid)
		categories, err := getCategories(uuid)
		if err != nil {
			return &proto.Books{}, err
		}
		ri := &proto.Book{
			BookUuid:   uuid,
			Name:       item.Name,
			Author:     item.Author.Name,
			Categories: categories,
		}

		response = append(response, ri)
	}

	return &proto.Books{Book: response}, nil
}

// PaginateAuthors godoc
// @Summary Get authors
// @Description shows authors by pages
// @ID paginate-authors
// @Accept  json
// @Produce  json
// @Param page_number path int true "Page number"
// @Success 200 {object} main.Author
// @Router /paginate-authors/{page_number} [get]
func (s *server) PaginateAuthors(ctx context.Context, request *proto.PageNumber) (*proto.Authors, error) {
	pageNumber := request.GetPageNumber()
	offset := (pageNumber - 1) * 10
	authors := []Author{}
	response := []*proto.Author{}

	selectQuery := `SELECT uuid, name
    FROM authors
    WHERE deleted_at IS NULL LIMIT 10 OFFSET ?`

	err := db.Select(&authors, selectQuery, offset)
	if err != nil {
		return &proto.Authors{}, err
	}

	for _, item := range authors {
		uuid := int64(item.Uuid)
		ri := &proto.Author{
			AuthorUuid: uuid,
			Name:       item.Name,
		}

		response = append(response, ri)
	}

	return &proto.Authors{Author: response}, nil
}

// PaginateCategories godoc
// @Summary Get categories
// @Description shows categories by pages
// @ID paginate-categories
// @Accept  json
// @Produce  json
// @Param page_number path int true "Page number"
// @Success 200 {object} main.Category
// @Router /paginate-categories/{page_number} [get]
func (s *server) PaginateCategories(ctx context.Context, request *proto.PageNumber) (*proto.Categories, error) {
	pageNumber := request.GetPageNumber()
	offset := (pageNumber - 1) * 10
	categories := []Category{}
	response := []*proto.Category{}

	selectQuery := `SELECT c.name, c.parent_uuid, IFNULL(c2.name, "") AS parent_name
    FROM categories c
    LEFT JOIN categories c2 ON c.parent_uuid = c2.uuid
    WHERE c.deleted_at IS NULL LIMIT 10 OFFSET ?`

	err := db.Select(&categories, selectQuery, offset)
	if err != nil {
		return &proto.Categories{}, err
	}

	for _, item := range categories {
		uuid := int64(item.Uuid)
		ri := &proto.Category{
			CategoryUuid: uuid,
			Name:         item.Name,
			ParentName:   item.Parent_name,
		}

		response = append(response, ri)
	}

	return &proto.Categories{Category: response}, nil
}

// DeleteAuthor godoc
// @Summary Deletes an author
// @Description delete an author using the DELETE request
// @ID delete-author
// @Accept  json
// @Produce  json
// @Param author_uuid path int true "Author uuid"
// @Success 200 {object} bool
// @Router /delete-author/{author_uuid} [delete]
func (s *server) DeleteAuthor(ctx context.Context, request *proto.AuthorId) (*proto.Response, error) {
	authorUuid := request.GetAuthorUuid()
	books := []Book{}

	err := deleteEntity("authors", authorUuid)
	if err != nil {
		return &proto.Response{Success: false}, err
	}
	selectBooksQuery := `SELECT uuid FROM books WHERE author_uuid = ?`
	err = db.Select(&books, selectBooksQuery, authorUuid)
	if err != nil {
		return &proto.Response{Success: false}, err
	}

	for _, item := range books {
		bookUuid := int64(item.Uuid)
		err := deleteEntity("books", bookUuid)
		if err != nil {
			return &proto.Response{Success: false}, err
		}

		err = deleteBookCategoriesLinking(bookUuid)
		if err != nil {
			return &proto.Response{Success: false}, err
		}
	}

	return &proto.Response{Success: true}, nil
}

func deleteBookCategoriesLinking(bookUuid int64) error {
	deleteQuery := `DELETE FROM books_categories WHERE book_uuid = ?`
	_, err = db.Exec(deleteQuery, bookUuid)

	return err
}

// DeleteCategory godoc
// @Summary Deletes a category
// @Description delete a category using the DELETE request
// @ID delete-category
// @Accept  json
// @Produce  json
// @Param category_uuid path int true "Category uuid"
// @Success 200 {object} bool
// @Router /delete-category/{category_uuid} [delete]
func (s *server) DeleteCategory(ctx context.Context, request *proto.CategoryId) (*proto.Response, error) {
	categoryUuid := request.GetCategoryUuid()

	err := deleteEntity("categories", categoryUuid)

	if err != nil {
		return &proto.Response{Success: false}, err
	}

	deleteQuery := `DELETE FROM books_categories WHERE category_uuid = ?`
	_, err = db.Exec(deleteQuery, categoryUuid)
	if err != nil {
		return &proto.Response{Success: false}, err
	}

	return &proto.Response{Success: true}, nil
}

func deleteEntity(entity string, entityUuid int64) error {
	currentTimestamp := time.Now().Format("2006-01-02 15:04:05")
	deleteQuery := fmt.Sprintf("UPDATE %s ", entity)
	deleteQuery = deleteQuery + "SET updated_at = ?, deleted_at = ? WHERE uuid = ?"

	_, err := db.Exec(deleteQuery, currentTimestamp, currentTimestamp, entityUuid)

	return err
}

func handleDatabase() {
	authorsQuery := `CREATE TABLE IF NOT EXISTS authors (
        uuid INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(60) NOT NULL,
        created_at datetime default CURRENT_TIMESTAMP,
        updated_at datetime default CURRENT_TIMESTAMP,
        deleted_at datetime
    )`

	db.MustExec(authorsQuery)

	booksQuery := `CREATE TABLE IF NOT EXISTS books (
        uuid INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(60) NOT NULL,
        author_uuid INT NOT NULL,
        created_at datetime default CURRENT_TIMESTAMP,
        updated_at datetime default CURRENT_TIMESTAMP,
        deleted_at datetime,
        FOREIGN KEY (author_uuid) REFERENCES authors(uuid) ON DELETE CASCADE
    )`

	db.MustExec(booksQuery)

	categoriesQuery := `CREATE TABLE IF NOT EXISTS categories (
        uuid INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(60) NOT NULL,
        parent_uuid INT,
        created_at datetime default CURRENT_TIMESTAMP,
        updated_at datetime default CURRENT_TIMESTAMP,
        deleted_at datetime
    )`

	db.MustExec(categoriesQuery)

	booksCategoriesQuery := `CREATE TABLE IF NOT EXISTS books_categories (
        uuid INT PRIMARY KEY AUTO_INCREMENT,
        book_uuid INT NOT NULL,
        category_uuid INT NOT NULL,
        FOREIGN KEY (category_uuid) REFERENCES categories(uuid) ON DELETE CASCADE
    )`

	db.MustExec(booksCategoriesQuery)
}
