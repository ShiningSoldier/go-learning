package main

import (
	proto "../proto"
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
	Name string `json:"name"`
}

func main() {
	handleDatabase()
	checkErr(err)

	listener, err := net.Listen("tcp", ":9876")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	proto.RegisterBooksServiceServer(srv, &server{})
	reflection.Register(srv)

	defer db.Close()
	if e := srv.Serve(listener); e != nil {
		checkErr(err)
	}
}

func (s *server) AddBook(ctx context.Context, request *proto.AddBookRequest) (*proto.Response, error) {
	name, category, author := request.GetBookName(), request.GetCategoryId(), request.GetAuthorId()
	categoriesSlice := strings.Split(category, ",")
	insertQuery := `INSERT INTO books(name, author_id) VALUES(?,?)`

	row, err := db.Exec(insertQuery, name, author)
	checkErr(err)
	lastInsertedId, err := row.LastInsertId()
	checkErr(err)

	err = addCategories(lastInsertedId, categoriesSlice)
	checkErr(err)

	return &proto.Response{Success: true}, nil
}

func (s *server) UpdateBook(ctx context.Context, request *proto.UpdateBookRequest) (*proto.Response, error) {
	bookUuid, name, category, author, currentTimestamp := request.GetBookUuid(), request.GetBookName(), request.GetCategoryId(), request.GetAuthorId(), time.Now().Format("2006-01-02 15:04:05")
	categoriesSlice := strings.Split(category, ",")
	updateQuery := `UPDATE books SET name = ?, author_id = ?, updated_at = ? WHERE uuid = ?`

	_, err := db.Exec(updateQuery, name, author, currentTimestamp, bookUuid)
	checkErr(err)

	err = addCategories(bookUuid, categoriesSlice)
	checkErr(err)

	return &proto.Response{Success: true}, nil
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

func (s *server) DeleteBook(ctx context.Context, request *proto.BookId) (*proto.Response, error) {
	bookUuid := request.GetBookUuid()

	err := deleteEntity("books", bookUuid)

	checkErr(err)

	return &proto.Response{Success: true}, nil
}

func (s *server) ShowBook(ctx context.Context, request *proto.BookId) (*proto.BookData, error) {
	bookUuid := request.GetBookUuid()
	book := Book{}

	selectBookQuery := `SELECT books.uuid, books.name,
       authors.uuid "authors.uuid", authors.name "authors.name"
    FROM books
    INNER JOIN authors ON authors.uuid = books.author_id
    WHERE books.deleted_at IS NULL AND books.uuid = ?`

	err := db.Get(&book, selectBookQuery, bookUuid)
	checkErr(err)

	categories := getCategories(bookUuid)

	result := fmt.Sprintf("Book name: %s, author name: %s, categories: %s", book.Name, book.Author.Name, strings.TrimSpace(categories))

	return &proto.BookData{Result: result}, nil
}

func getCategories(bookUuid int64) string {
	booksCategories := []BooksCategories{}

	selectBooksCategoriesQuery := `SELECT category_uuid, categories.name "categories.name" FROM books_categories
    INNER JOIN categories ON categories.uuid = books_categories.category_uuid
    WHERE book_uuid = ?`
	err = db.Select(&booksCategories, selectBooksCategoriesQuery, bookUuid)
	checkErr(err)

	categories := ""

	for _, item := range booksCategories {
		categories = categories + item.Category.Name + "; "
	}

	return categories
}

func (s *server) AddCategory(ctx context.Context, request *proto.AddCategoryRequest) (*proto.Response, error) {
	name, parentId := request.GetName(), request.GetParentUuid()
	insertQuery := `INSERT INTO categories(name, parent_uuid) VALUES(?,?)`

	_, err := db.Exec(insertQuery, name, parentId)
	checkErr(err)

	return &proto.Response{Success: true}, nil
}

func (s *server) AddAuthor(ctx context.Context, request *proto.AddAuthorRequest) (*proto.Response, error) {
	name := request.GetName()
	insertQuery := `INSERT INTO authors(name) VALUES(?)`

	_, err := db.Exec(insertQuery, name)
	checkErr(err)

	return &proto.Response{Success: true}, nil
}

func (s *server) ShowAuthor(ctx context.Context, request *proto.AuthorId) (*proto.AuthorData, error) {
	authorUuid := request.GetAuthorUuid()

	selectQuery := `SELECT uuid, name, deleted_at FROM authors WHERE authors.uuid = ?`

	row, err := db.Query(selectQuery, authorUuid)
	checkErr(err)
	var (
		uuid      int
		name      string
		deletedAt sql.NullString
		result    string
	)

	for row.Next() {
		err := row.Scan(&uuid, &name, &deletedAt)
		checkErr(err)
		if deletedAt.Valid {
			result = fmt.Sprintf("This author was deleted at %s", deletedAt.String)
		} else {
			result = fmt.Sprintf("Author id: %d, name: %s", uuid, name)
		}
	}

	return &proto.AuthorData{Result: result}, nil
}

func (s *server) ShowCategory(ctx context.Context, request *proto.CategoryId) (*proto.CategoryData, error) {
	categoryUuid := request.GetCategoryUuid()

	selectQuery := `SELECT uuid, name, deleted_at, parent_uuid FROM categories WHERE categories.uuid = ?`

	row, err := db.Query(selectQuery, categoryUuid)
	checkErr(err)
	var (
		uuid       int
		name       string
		deletedAt  sql.NullString
		parentUuid sql.NullString
		result     string
	)

	for row.Next() {
		err := row.Scan(&uuid, &name, &deletedAt, &parentUuid)
		checkErr(err)
		parent := "no parent"
		if parentUuid.Valid && parentUuid.String != "0" {
			parent = parentUuid.String
		}
		if deletedAt.Valid {
			result = fmt.Sprintf("This category was deleted at %s", deletedAt.String)
		} else {
			result = fmt.Sprintf("Category id: %d, name: %s, parent: %s", uuid, name, parent)
		}
	}

	return &proto.CategoryData{Result: result}, nil
}

func (s *server) FilterByAuthor(ctx context.Context, request *proto.AuthorId) (*proto.BookData, error) {
	authorUuid := request.GetAuthorUuid()

	selectQuery := `SELECT books.uuid, books.name FROM books WHERE books.deleted_at IS NULL AND books.author_id = ?`

	row, err := db.Query(selectQuery, authorUuid)
	checkErr(err)
	var (
		uuid   int
		name   string
		result string
	)

	for row.Next() {
		err := row.Scan(&uuid, &name)
		checkErr(err)
		categories := getCategories(int64(uuid))
		result = result + fmt.Sprintf("Book uuid: %d, book name: %s, categories: %s", uuid, name, categories)
	}

	return &proto.BookData{Result: strings.TrimSpace(result)}, nil
}

func (s *server) FilterByCategory(ctx context.Context, request *proto.CategoryId) (*proto.BookData, error) {
	categoryUuid := request.GetCategoryUuid()

	selectQuery := `SELECT books.uuid, books.name FROM books
    INNER JOIN books_categories on books_categories.book_uuid = books.uuid
    INNER JOIN categories ON categories.uuid = books_categories.category_uuid
    WHERE books.deleted_at IS NULL AND categories.uuid = ?`

	row, err := db.Query(selectQuery, categoryUuid)
	checkErr(err)
	var (
		uuid   int
		name   string
		result string
	)

	for row.Next() {
		err := row.Scan(&uuid, &name)
		checkErr(err)
		categories := getCategories(int64(uuid))
		result = result + fmt.Sprintf("Book uuid: %d, book name: %s, categories: %s", uuid, name, categories)
	}

	return &proto.BookData{Result: strings.TrimSpace(result)}, nil
}

func (s *server) Paginate(ctx context.Context, request *proto.PageNumber) (*proto.BookData, error) {
	pageNumber := request.GetPageNumber()
	offset := (pageNumber - 1) * 10

	selectQuery := fmt.Sprintf("SELECT books.uuid, books.name, authors.name FROM books INNER JOIN authors ON authors.uuid = books.author_id WHERE books.deleted_at IS NULL LIMIT 10 OFFSET %d", offset)

	row, err := db.Query(selectQuery)
	checkErr(err)
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

func (s *server) DeleteAuthor(ctx context.Context, request *proto.AuthorId) (*proto.Response, error) {
	authorUuid := request.GetAuthorUuid()

	err := deleteEntity("authors", authorUuid)

	checkErr(err)

	return &proto.Response{Success: true}, nil
}

func (s *server) DeleteCategory(ctx context.Context, request *proto.CategoryId) (*proto.Response, error) {
	categoryUuid := request.GetCategoryUuid()

	err := deleteEntity("categories", categoryUuid)

	checkErr(err)

	return &proto.Response{Success: true}, nil
}

func deleteEntity(entity string, entityUuid int64) error {
	currentTimestamp := time.Now().Format("2006-01-02 15:04:05")
	deleteQuery := fmt.Sprintf("UPDATE %s ", entity)
	deleteQuery = deleteQuery + "SET updated_at = ?, deleted_at = ? WHERE uuid = ?"

	_, err := db.Exec(deleteQuery, currentTimestamp, currentTimestamp, entityUuid)

	return err
}

func (s *server) UpdateAuthor(ctx context.Context, request *proto.UpdateAuthorRequest) (*proto.Response, error) {
	authorUuid, name, currentTimestamp := request.GetAuthorUuid(), request.GetAuthorName(), time.Now().Format("2006-01-02 15:04:05")
	updateQuery := `UPDATE authors SET name = ?, updated_at = ? WHERE uuid = ?`

	_, err := db.Exec(updateQuery, name, currentTimestamp, authorUuid)
	checkErr(err)

	return &proto.Response{Success: true}, nil
}

func (s *server) UpdateCategory(ctx context.Context, request *proto.UpdateCategoryRequest) (*proto.Response, error) {
	categoryId, name, parentUuid, currentTimestamp := request.GetCategoryUuid(), request.GetCategoryName(), request.GetParentUuid(), time.Now().Format("2006-01-02 15:04:05")
	updateQuery := `UPDATE categories SET name = ?, parent_uuid = ?, updated_at = ? WHERE uuid = ?`

	_, err := db.Exec(updateQuery, name, parentUuid, currentTimestamp, categoryId)
	checkErr(err)

	return &proto.Response{Success: true}, nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
        author_id INT NOT NULL,
        created_at datetime default CURRENT_TIMESTAMP,
        updated_at datetime default CURRENT_TIMESTAMP,
        deleted_at datetime,
        FOREIGN KEY (author_id) REFERENCES authors(uuid) ON DELETE CASCADE
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
