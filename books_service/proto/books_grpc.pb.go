// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BooksServiceClient is the client API for BooksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BooksServiceClient interface {
	AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*Book, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error)
	DeleteBook(ctx context.Context, in *BookId, opts ...grpc.CallOption) (*Response, error)
	ShowBook(ctx context.Context, in *BookId, opts ...grpc.CallOption) (*BookData, error)
	FilterByAuthor(ctx context.Context, in *AuthorId, opts ...grpc.CallOption) (*BookData, error)
	FilterByCategory(ctx context.Context, in *CategoryId, opts ...grpc.CallOption) (*BookData, error)
	Paginate(ctx context.Context, in *PageNumber, opts ...grpc.CallOption) (*BookData, error)
	AddAuthor(ctx context.Context, in *AddAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*Category, error)
	ShowAuthor(ctx context.Context, in *AuthorId, opts ...grpc.CallOption) (*AuthorData, error)
	UpdateAuthor(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*Response, error)
	ShowCategory(ctx context.Context, in *CategoryId, opts ...grpc.CallOption) (*CategoryData, error)
	DeleteCategory(ctx context.Context, in *CategoryId, opts ...grpc.CallOption) (*Response, error)
	DeleteAuthor(ctx context.Context, in *AuthorId, opts ...grpc.CallOption) (*Response, error)
}

type booksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBooksServiceClient(cc grpc.ClientConnInterface) BooksServiceClient {
	return &booksServiceClient{cc}
}

func (c *booksServiceClient) AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/proto.BooksService/AddBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/proto.BooksService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) DeleteBook(ctx context.Context, in *BookId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.BooksService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) ShowBook(ctx context.Context, in *BookId, opts ...grpc.CallOption) (*BookData, error) {
	out := new(BookData)
	err := c.cc.Invoke(ctx, "/proto.BooksService/ShowBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) FilterByAuthor(ctx context.Context, in *AuthorId, opts ...grpc.CallOption) (*BookData, error) {
	out := new(BookData)
	err := c.cc.Invoke(ctx, "/proto.BooksService/FilterByAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) FilterByCategory(ctx context.Context, in *CategoryId, opts ...grpc.CallOption) (*BookData, error) {
	out := new(BookData)
	err := c.cc.Invoke(ctx, "/proto.BooksService/FilterByCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) Paginate(ctx context.Context, in *PageNumber, opts ...grpc.CallOption) (*BookData, error) {
	out := new(BookData)
	err := c.cc.Invoke(ctx, "/proto.BooksService/Paginate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) AddAuthor(ctx context.Context, in *AddAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/proto.BooksService/AddAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/proto.BooksService/AddCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) ShowAuthor(ctx context.Context, in *AuthorId, opts ...grpc.CallOption) (*AuthorData, error) {
	out := new(AuthorData)
	err := c.cc.Invoke(ctx, "/proto.BooksService/ShowAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) UpdateAuthor(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.BooksService/UpdateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.BooksService/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) ShowCategory(ctx context.Context, in *CategoryId, opts ...grpc.CallOption) (*CategoryData, error) {
	out := new(CategoryData)
	err := c.cc.Invoke(ctx, "/proto.BooksService/ShowCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) DeleteCategory(ctx context.Context, in *CategoryId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.BooksService/DeleteCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) DeleteAuthor(ctx context.Context, in *AuthorId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.BooksService/DeleteAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BooksServiceServer is the server API for BooksService service.
// All implementations must embed UnimplementedBooksServiceServer
// for forward compatibility
type BooksServiceServer interface {
	AddBook(context.Context, *AddBookRequest) (*Book, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
	DeleteBook(context.Context, *BookId) (*Response, error)
	ShowBook(context.Context, *BookId) (*BookData, error)
	FilterByAuthor(context.Context, *AuthorId) (*BookData, error)
	FilterByCategory(context.Context, *CategoryId) (*BookData, error)
	Paginate(context.Context, *PageNumber) (*BookData, error)
	AddAuthor(context.Context, *AddAuthorRequest) (*Author, error)
	AddCategory(context.Context, *AddCategoryRequest) (*Category, error)
	ShowAuthor(context.Context, *AuthorId) (*AuthorData, error)
	UpdateAuthor(context.Context, *UpdateAuthorRequest) (*Response, error)
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*Response, error)
	ShowCategory(context.Context, *CategoryId) (*CategoryData, error)
	DeleteCategory(context.Context, *CategoryId) (*Response, error)
	DeleteAuthor(context.Context, *AuthorId) (*Response, error)
	mustEmbedUnimplementedBooksServiceServer()
}

// UnimplementedBooksServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBooksServiceServer struct {
}

func (UnimplementedBooksServiceServer) AddBook(context.Context, *AddBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (UnimplementedBooksServiceServer) UpdateBook(context.Context, *UpdateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBooksServiceServer) DeleteBook(context.Context, *BookId) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBooksServiceServer) ShowBook(context.Context, *BookId) (*BookData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowBook not implemented")
}
func (UnimplementedBooksServiceServer) FilterByAuthor(context.Context, *AuthorId) (*BookData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterByAuthor not implemented")
}
func (UnimplementedBooksServiceServer) FilterByCategory(context.Context, *CategoryId) (*BookData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterByCategory not implemented")
}
func (UnimplementedBooksServiceServer) Paginate(context.Context, *PageNumber) (*BookData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Paginate not implemented")
}
func (UnimplementedBooksServiceServer) AddAuthor(context.Context, *AddAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAuthor not implemented")
}
func (UnimplementedBooksServiceServer) AddCategory(context.Context, *AddCategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedBooksServiceServer) ShowAuthor(context.Context, *AuthorId) (*AuthorData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAuthor not implemented")
}
func (UnimplementedBooksServiceServer) UpdateAuthor(context.Context, *UpdateAuthorRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthor not implemented")
}
func (UnimplementedBooksServiceServer) UpdateCategory(context.Context, *UpdateCategoryRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedBooksServiceServer) ShowCategory(context.Context, *CategoryId) (*CategoryData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowCategory not implemented")
}
func (UnimplementedBooksServiceServer) DeleteCategory(context.Context, *CategoryId) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedBooksServiceServer) DeleteAuthor(context.Context, *AuthorId) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthor not implemented")
}
func (UnimplementedBooksServiceServer) mustEmbedUnimplementedBooksServiceServer() {}

// UnsafeBooksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BooksServiceServer will
// result in compilation errors.
type UnsafeBooksServiceServer interface {
	mustEmbedUnimplementedBooksServiceServer()
}

func RegisterBooksServiceServer(s grpc.ServiceRegistrar, srv BooksServiceServer) {
	s.RegisterService(&BooksService_ServiceDesc, srv)
}

func _BooksService_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).AddBook(ctx, req.(*AddBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).DeleteBook(ctx, req.(*BookId))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_ShowBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).ShowBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/ShowBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).ShowBook(ctx, req.(*BookId))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_FilterByAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).FilterByAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/FilterByAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).FilterByAuthor(ctx, req.(*AuthorId))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_FilterByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).FilterByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/FilterByCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).FilterByCategory(ctx, req.(*CategoryId))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_Paginate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).Paginate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/Paginate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).Paginate(ctx, req.(*PageNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_AddAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).AddAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/AddAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).AddAuthor(ctx, req.(*AddAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/AddCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).AddCategory(ctx, req.(*AddCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_ShowAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).ShowAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/ShowAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).ShowAuthor(ctx, req.(*AuthorId))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_UpdateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).UpdateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/UpdateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).UpdateAuthor(ctx, req.(*UpdateAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).UpdateCategory(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_ShowCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).ShowCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/ShowCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).ShowCategory(ctx, req.(*CategoryId))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/DeleteCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).DeleteCategory(ctx, req.(*CategoryId))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_DeleteAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).DeleteAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BooksService/DeleteAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).DeleteAuthor(ctx, req.(*AuthorId))
	}
	return interceptor(ctx, in, info, handler)
}

// BooksService_ServiceDesc is the grpc.ServiceDesc for BooksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BooksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BooksService",
	HandlerType: (*BooksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBook",
			Handler:    _BooksService_AddBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BooksService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BooksService_DeleteBook_Handler,
		},
		{
			MethodName: "ShowBook",
			Handler:    _BooksService_ShowBook_Handler,
		},
		{
			MethodName: "FilterByAuthor",
			Handler:    _BooksService_FilterByAuthor_Handler,
		},
		{
			MethodName: "FilterByCategory",
			Handler:    _BooksService_FilterByCategory_Handler,
		},
		{
			MethodName: "Paginate",
			Handler:    _BooksService_Paginate_Handler,
		},
		{
			MethodName: "AddAuthor",
			Handler:    _BooksService_AddAuthor_Handler,
		},
		{
			MethodName: "AddCategory",
			Handler:    _BooksService_AddCategory_Handler,
		},
		{
			MethodName: "ShowAuthor",
			Handler:    _BooksService_ShowAuthor_Handler,
		},
		{
			MethodName: "UpdateAuthor",
			Handler:    _BooksService_UpdateAuthor_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _BooksService_UpdateCategory_Handler,
		},
		{
			MethodName: "ShowCategory",
			Handler:    _BooksService_ShowCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _BooksService_DeleteCategory_Handler,
		},
		{
			MethodName: "DeleteAuthor",
			Handler:    _BooksService_DeleteAuthor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/books.proto",
}
