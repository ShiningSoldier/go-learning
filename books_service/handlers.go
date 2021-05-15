package main

import (
	proto "./proto"
	"context"
	"database/sql"
	"fmt"
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
	Uuid   int    `json:"uuid"`
	Name   string `json:"name"`
	Author `db:"authors"`
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
	Name        string         `json:"name"`
	Parent_uuid int            `json:"parent_uuid"`
	Parent_name sql.NullString `json:"parent_name"`
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
// @Accept  json
// @Produce  json
// @Param name body string true "Book name"
// @Param category_uuid body string true "List of category iIDs"
// @Param author_uuid body int true "Book author ID"
// @Success 200 {object} main.Book
// @Router /add [post]
func (s *server) AddBook(ctx context.Context, request *proto.AddBookRequest) (*proto.Book, error) {
	name, category, author := request.GetBookName(), request.GetCategoryId(), request.GetAuthorId()
	categoriesSlice := strings.Split(category, ",")
	insertQuery := `INSERT INTO books(name, author_uuid) VALUES(?,?)`

	row, err := db.Exec(insertQuery, name, author)
	if err != nil {
		return &proto.Book{
			BookUuid:   0,
			Name:       "",
			Author:     "",
			Categories: "",
		}, err
	}
	lastInsertedId, err := row.LastInsertId()
	if err != nil {
		return &proto.Book{
			BookUuid:   0,
			Name:       "",
			Author:     "",
			Categories: "",
		}, err
	}

	err = addCategories(lastInsertedId, categoriesSlice)
	if err != nil {
		return &proto.Book{
			BookUuid:   0,
			Name:       "",
			Author:     "",
			Categories: "",
		}, err
	}

	authorData, err := getAuthor(author)
	if err != nil {
		return &proto.Book{
			BookUuid:   0,
			Name:       "",
			Author:     "",
			Categories: "",
		}, err
	}

	categories := getCategories(lastInsertedId)

	return &proto.Book{
		BookUuid:   lastInsertedId,
		Name:       name,
		Author:     authorData.Name,
		Categories: categories,
	}, nil
}

// UpdateBook godoc
// @Summary Updates a book
// @Description update a book using the PUT request
// @ID update-book
// @Accept  json
// @Produce  json
// @Param book_uuid body int true "Book uuid"
// @Param name body string true "Book name"
// @Param category_uuid body string true "List of category iIDs"
// @Param author_uuid body int true "Book author ID"
// @Success 200 {object} main.Book
// @Router /update [put]
func (s *server) UpdateBook(ctx context.Context, request *proto.UpdateBookRequest) (*proto.Book, error) {
	bookUuid, name, category, author, currentTimestamp := request.GetBookUuid(), request.GetBookName(), request.GetCategoryId(), request.GetAuthorId(), time.Now().Format("2006-01-02 15:04:05")
	categoriesSlice := strings.Split(category, ",")
	updateQuery := `UPDATE books SET name = ?, author_uuid = ?, updated_at = ? WHERE uuid = ?`

	_, err := db.Exec(updateQuery, name, author, currentTimestamp, bookUuid)
	if err != nil {
		return &proto.Book{
			BookUuid:   0,
			Name:       "",
			Author:     "",
			Categories: "",
		}, err
	}

	err = addCategories(bookUuid, categoriesSlice)
	if err != nil {
		return &proto.Book{
			BookUuid:   0,
			Name:       "",
			Author:     "",
			Categories: "",
		}, err
	}

	authorData, err := getAuthor(author)
	if err != nil {
		return &proto.Book{
			BookUuid:   0,
			Name:       "",
			Author:     "",
			Categories: "",
		}, err
	}

	categories := getCategories(bookUuid)

	return &proto.Book{
		BookUuid:   bookUuid,
		Name:       name,
		Author:     authorData.Name,
		Categories: categories,
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

	return &proto.Response{Success: true}, nil
}

// ShowBook godoc
// @Summary Shows a book
// @Description shows the main data about a book
// @ID show-book
// @Accept  json
// @Produce  json
// @Param book_uuid path int true "Book uuid"
// @Success 200 {object} bool
// @Router /show/{book_uuid} [get]
func (s *server) ShowBook(ctx context.Context, request *proto.BookId) (*proto.BookData, error) {
	bookUuid := request.GetBookUuid()

	result, err := getBookData(bookUuid)

	if err != nil {
		return &proto.BookData{Result: ""}, err
	}

	return &proto.BookData{Result: result}, nil
}

func getBookData(bookUuid int64) (string, error) {
	book := Book{}

	selectBookQuery := `SELECT books.uuid, books.name,
       authors.uuid "authors.uuid", authors.name "authors.name"
    FROM books
    INNER JOIN authors ON authors.uuid = books.author_id
    WHERE books.deleted_at IS NULL AND books.uuid = ?`

	err := db.Get(&book, selectBookQuery, bookUuid)
	if err != nil {
		return "", err
	}

	categories := getCategories(bookUuid)
	result := fmt.Sprintf("Book uuid: %d, Book name: %s, author name: %s, categories: %s", book.Uuid, book.Name, book.Author.Name, strings.TrimSpace(categories))

	return result, nil
}

func getCategories(bookUuid int64) string {
	booksCategories := []BooksCategories{}

	selectBooksCategoriesQuery := `SELECT category_uuid, categories.name "categories.name",
       categories.parent_uuid "categories.parent_uuid"
    FROM books_categories
    INNER JOIN categories ON categories.uuid = books_categories.category_uuid
    WHERE book_uuid = ?`
	err = db.Select(&booksCategories, selectBooksCategoriesQuery, bookUuid)
	checkErr(err)

	categories := ""

	for _, item := range booksCategories {
		categories = categories + item.Category.Name + "; "
	}

	return strings.TrimSpace(categories)
}

// AddCategory godoc
// @Summary Create a category
// @Description creates a new category
// @ID create-category
// @Accept  json
// @Produce  json
// @Param name body string true "Category name"
// @Param parent_uuid body string false "Parent id"
// @Success 200 {object} main.Category
// @Router /add-category [post]
func (s *server) AddCategory(ctx context.Context, request *proto.AddCategoryRequest) (*proto.Category, error) {
	name, parentUuid := request.GetName(), request.GetParentUuid()
	if len(parentUuid) == 0 {
		parentUuid = "0"
	}
	insertQuery := `INSERT INTO categories(name, parent_uuid) VALUES(?,?)`

	row, err := db.Exec(insertQuery, name, parentUuid)
	if err != nil {
		return &proto.Category{
			CategoryUuid: 0,
			ParentUuid:   "",
			Name:         "",
		}, err
	}
	lastInsertedId, err := row.LastInsertId()
	if err != nil {
		return &proto.Category{
			CategoryUuid: 0,
			ParentUuid:   "",
			Name:         "",
		}, err
	}

	return &proto.Category{
		CategoryUuid: lastInsertedId,
		ParentUuid:   parentUuid,
		Name:         name,
	}, nil
}

// AddAuthor godoc
// @Summary Create an author
// @Description creates a new author
// @ID create-author
// @Accept  json
// @Produce  json
// @Param name body string true "Author name"
// @Success 200 {object} main.Author
// @Router /add-author [post]
func (s *server) AddAuthor(ctx context.Context, request *proto.AddAuthorRequest) (*proto.Author, error) {
	name := request.GetName()
	insertQuery := `INSERT INTO authors(name) VALUES(?)`

	row, err := db.Exec(insertQuery, name)
	if err != nil {
		return &proto.Author{
			AuthorUuid: 0,
			Name:       "",
		}, err
	}

	lastInsertedId, err := row.LastInsertId()
	if err != nil {
		return &proto.Author{
			AuthorUuid: 0,
			Name:       "",
		}, err
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
// @Success 200 {object} string
// @Router /show-author/{author_uuid} [get]
func (s *server) ShowAuthor(ctx context.Context, request *proto.AuthorId) (*proto.AuthorData, error) {
	authorUuid := request.GetAuthorUuid()
	author, err := getAuthor(authorUuid)
	if err != nil {
		return &proto.AuthorData{Result: ""}, err
	}

	return &proto.AuthorData{Result: fmt.Sprintf("Author name: %s", author.Name)}, nil
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
// @Success 200 {object} string
// @Router /show-category/{category_uuid} [get]
func (s *server) ShowCategory(ctx context.Context, request *proto.CategoryId) (*proto.CategoryData, error) {
	categoryUuid := request.GetCategoryUuid()
	category := Category{}

	selectQuery := `SELECT c.name, c.parent_uuid, c2.name AS parent_name
    FROM categories c
    LEFT JOIN categories c2 ON c.parent_uuid = c2.uuid
    WHERE c.deleted_at IS NULL AND c.uuid = ?`

	err := db.Get(&category, selectQuery, categoryUuid)
	if err != nil {
		return &proto.CategoryData{Result: ""}, err
	}

	return &proto.CategoryData{Result: fmt.Sprintf("Category name: %s, parent category: %s", category.Name, category.Parent_name.String)}, nil
}

// FilterByAuthor godoc
// @Summary Shows all books by specified author
// @Description shows the basic data of books by author
// @ID filter-by-author
// @Accept  json
// @Produce  json
// @Param author_uuid path int true "Author uuid"
// @Success 200 {object} string
// @Router /filter-by-author/{author_uuid} [get]
func (s *server) FilterByAuthor(ctx context.Context, request *proto.AuthorId) (*proto.BookData, error) {
	authorUuid := request.GetAuthorUuid()
	books := []Book{}

	selectQuery := `SELECT books.uuid, books.name, authors.name "authors.name"
    FROM books
    INNER JOIN authors ON authors.uuid = books.author_id
    WHERE books.deleted_at IS NULL AND books.author_id = ?`

	err := db.Select(&books, selectQuery, authorUuid)
	if err != nil {
		return &proto.BookData{Result: ""}, err
	}

	result := ""

	for _, item := range books {
		categories := getCategories(int64(item.Uuid))
		result = result + fmt.Sprintf("Book name: %s, author name: %s, categories: %s", item.Name, item.Author.Name, categories)
	}

	return &proto.BookData{Result: strings.TrimSpace(result)}, nil
}

// FilterByCategory godoc
// @Summary Shows all books by specified category
// @Description shows the basic data of books by category
// @ID filter-by-category
// @Accept  json
// @Produce  json
// @Param category_uuid path int true "Category uuid"
// @Success 200 {object} string
// @Router /filter-by-category/{category_uuid} [get]
func (s *server) FilterByCategory(ctx context.Context, request *proto.CategoryId) (*proto.BookData, error) {
	categoryUuid := request.GetCategoryUuid()
	books := []Book{}

	selectQuery := `SELECT books.uuid, books.name, authors.name "authors.name" FROM books
    INNER JOIN authors ON authors.uuid = books.author_id
    INNER JOIN books_categories ON books_categories.book_uuid = books.uuid
    WHERE books.deleted_at IS NULL AND books_categories.category_uuid = ?`

	err := db.Select(&books, selectQuery, categoryUuid)
	if err != nil {
		return &proto.BookData{Result: ""}, err
	}

	result := ""

	for _, item := range books {
		categories := getCategories(int64(item.Uuid))
		result = result + fmt.Sprintf("Book name: %s, author name: %s, categories: %s", item.Name, item.Author.Name, categories)
	}

	return &proto.BookData{Result: strings.TrimSpace(result)}, nil
}

// Paginate godoc
// @Summary Paginate books
// @Description shows the books by pages
// @ID paginate
// @Accept  json
// @Produce  json
// @Param page_number path int true "Page number"
// @Success 200 {object} string
// @Router /paginate/{page_number} [get]
func (s *server) Paginate(ctx context.Context, request *proto.PageNumber) (*proto.BookData, error) {
	pageNumber := request.GetPageNumber()
	offset := (pageNumber - 1) * 10

	selectQuery := fmt.Sprintf("SELECT books.uuid, books.name, authors.name FROM books INNER JOIN authors ON authors.uuid = books.author_id WHERE books.deleted_at IS NULL LIMIT 10 OFFSET %d", offset)

	row, err := db.Query(selectQuery)
	if err != nil {
		return &proto.BookData{Result: ""}, err
	}

	var (
		uuid       int
		name       string
		authorName string
		result     string
	)

	for row.Next() {
		err := row.Scan(&uuid, &name, &authorName)
		checkErr(err)
		categories := getCategories(int64(uuid))
		result = result + fmt.Sprintf("Book uuid: %d, book name: %s, categories: %s", uuid, name, categories)
	}

	return &proto.BookData{Result: strings.TrimSpace(result)}, nil
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
		err := deleteEntity("books", int64(item.Uuid))
		if err != nil {
			return &proto.Response{Success: false}, err
		}
	}

	return &proto.Response{Success: true}, nil
}

// DeleteCategory godoc
// @Summary Deletes a category
// @Description delete a category using the DELETE request
// @ID delete-category
// @Accept  json
// @Produce  json
// @Param category_uuid path int true "Category uuid"
// @Success 200 {object} bool
// @Router /delete-category/{delete_category} [delete]
func (s *server) DeleteCategory(ctx context.Context, request *proto.CategoryId) (*proto.Response, error) {
	categoryUuid := request.GetCategoryUuid()

	err := deleteEntity("categories", categoryUuid)

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

// UpdateAuthor godoc
// @Summary Updates an author
// @Description update an author using the PUT request
// @ID update-author
// @Accept  json
// @Produce  json
// @Param author_uuid body int true "Author uuid"
// @Param name body string true "Author name"
// @Success 200 {object} bool
// @Router /update-author [put]
func (s *server) UpdateAuthor(ctx context.Context, request *proto.UpdateAuthorRequest) (*proto.Response, error) {
	authorUuid, name, currentTimestamp := request.GetAuthorUuid(), request.GetAuthorName(), time.Now().Format("2006-01-02 15:04:05")
	updateQuery := `UPDATE authors SET name = ?, updated_at = ? WHERE uuid = ?`

	_, err := db.Exec(updateQuery, name, currentTimestamp, authorUuid)
	if err != nil {
		return &proto.Response{Success: false}, err
	}

	return &proto.Response{Success: true}, nil
}

// UpdateCategory godoc
// @Summary Updates a category
// @Description update a category using the PUT request
// @ID update-category
// @Accept  json
// @Produce  json
// @Param category_uuid body int true "Category uuid"
// @Param name path string true "Category name"
// @Param parent_uuid path int true "Parent id"
// @Success 200 {object} bool
// @Router /update-category [put]
func (s *server) UpdateCategory(ctx context.Context, request *proto.UpdateCategoryRequest) (*proto.Response, error) {
	categoryId, name, parentUuid, currentTimestamp := request.GetCategoryUuid(), request.GetCategoryName(), request.GetParentUuid(), time.Now().Format("2006-01-02 15:04:05")
	updateQuery := `UPDATE categories SET name = ?, parent_uuid = ?, updated_at = ? WHERE uuid = ?`

	_, err := db.Exec(updateQuery, name, parentUuid, currentTimestamp, categoryId)
	if err != nil {
		return &proto.Response{Success: false}, err
	}

	return &proto.Response{Success: true}, nil
}

func (s *server) GetBookData(ctx context.Context, request *proto.BookId) (*proto.BookData, error) {
	bookUuid := request.GetBookUuid()

	result, err := getBookData(bookUuid)

	if err != nil {
		return &proto.BookData{Result: ""}, err
	}

	return &proto.BookData{Result: result}, nil
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

func checkErr(err error) {
	if err != nil {
		log.Printf(err.Error())
	}
}
