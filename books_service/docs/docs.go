// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/add": {
            "post": {
                "description": "create a book using the POST request",
                "produces": [
                    "application/json"
                ],
                "summary": "Creates a new book",
                "operationId": "add-book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "List of category iIDs",
                        "name": "category_uuid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Book author ID",
                        "name": "author_uuid",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Book"
                        }
                    }
                }
            }
        },
        "/add-author": {
            "post": {
                "description": "creates a new author",
                "produces": [
                    "application/json"
                ],
                "summary": "Create an author",
                "operationId": "create-author",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Author"
                        }
                    }
                }
            }
        },
        "/add-category": {
            "post": {
                "description": "creates a new category",
                "produces": [
                    "application/json"
                ],
                "summary": "Create a category",
                "operationId": "create-category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Parent id",
                        "name": "parent_uuid",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Category"
                        }
                    }
                }
            }
        },
        "/delete-author/{author_uuid}": {
            "delete": {
                "description": "delete an author using the DELETE request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Deletes an author",
                "operationId": "delete-author",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Author uuid",
                        "name": "author_uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/delete-category/{category_uuid}": {
            "delete": {
                "description": "delete a category using the DELETE request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Deletes a category",
                "operationId": "delete-category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category uuid",
                        "name": "category_uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/delete/{book_uuid}": {
            "delete": {
                "description": "delete a book using the DELETE request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Deletes a book",
                "operationId": "delete-book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book uuid",
                        "name": "book_uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/filter-by-author/{author_uuid}": {
            "get": {
                "description": "shows the basic data of books by author",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Shows all books by specified author",
                "operationId": "filter-by-author",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Author uuid",
                        "name": "author_uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Book"
                        }
                    }
                }
            }
        },
        "/filter-by-category/{category_uuid}": {
            "get": {
                "description": "shows the basic data of books by category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Shows all books by specified category",
                "operationId": "filter-by-category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category uuid",
                        "name": "category_uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Category"
                        }
                    }
                }
            }
        },
        "/paginate-authors/{page_number}": {
            "get": {
                "description": "shows authors by pages",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get authors",
                "operationId": "paginate-authors",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page_number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Author"
                        }
                    }
                }
            }
        },
        "/paginate-categories/{page_number}": {
            "get": {
                "description": "shows categories by pages",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get categories",
                "operationId": "paginate-categories",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page_number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Category"
                        }
                    }
                }
            }
        },
        "/paginate/{page_number}": {
            "get": {
                "description": "shows the books by pages",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Paginate books",
                "operationId": "paginate",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page_number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Book"
                        }
                    }
                }
            }
        },
        "/show-author/{author_uuid}": {
            "get": {
                "description": "show the basic author data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show author data",
                "operationId": "show-author",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Author uuid",
                        "name": "author_uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Author"
                        }
                    }
                }
            }
        },
        "/show-category/{category_uuid}": {
            "get": {
                "description": "show the basic category data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show category data",
                "operationId": "show-category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category uuid",
                        "name": "category_uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Category"
                        }
                    }
                }
            }
        },
        "/show/{book_uuid}": {
            "get": {
                "description": "shows the main data about a book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Shows a book",
                "operationId": "show-book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book uuid",
                        "name": "book_uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Book"
                        }
                    }
                }
            }
        },
        "/update": {
            "patch": {
                "description": "update a book using the PATCH request",
                "produces": [
                    "application/json"
                ],
                "summary": "Updates a book",
                "operationId": "update-book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book uuid",
                        "name": "book_uuid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Book name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "List of category iIDs",
                        "name": "category_uuid",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Book author ID",
                        "name": "author_uuid",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Book"
                        }
                    }
                }
            }
        },
        "/update-author": {
            "patch": {
                "description": "update an author using the PATCH request",
                "produces": [
                    "application/json"
                ],
                "summary": "Updates an author",
                "operationId": "update-author",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Author uuid",
                        "name": "author_uuid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Author name",
                        "name": "name",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Author"
                        }
                    }
                }
            }
        },
        "/update-category": {
            "patch": {
                "description": "update a category using the PUT request",
                "produces": [
                    "application/json"
                ],
                "summary": "Updates a category",
                "operationId": "update-category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category uuid",
                        "name": "category_uuid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Category name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Parent id",
                        "name": "parent_uuid",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Category"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Author": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "uuid": {
                    "type": "integer"
                }
            }
        },
        "main.Book": {
            "type": "object",
            "properties": {
                "author_uuid": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "uuid": {
                    "type": "integer"
                }
            }
        },
        "main.Category": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "parent_name": {
                    "type": "string"
                },
                "parent_uuid": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
