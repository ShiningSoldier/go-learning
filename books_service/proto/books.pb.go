// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/books.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookName   string `protobuf:"bytes,1,opt,name=book_name,json=bookName,proto3" json:"book_name,omitempty"`
	CategoryId string `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	AuthorId   int64  `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *AddBookRequest) Reset() {
	*x = AddBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookRequest) ProtoMessage() {}

func (x *AddBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookRequest.ProtoReflect.Descriptor instead.
func (*AddBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{0}
}

func (x *AddBookRequest) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

func (x *AddBookRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *AddBookRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type AddCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParentUuid int64  `protobuf:"varint,2,opt,name=parent_uuid,json=parentUuid,proto3" json:"parent_uuid,omitempty"`
}

func (x *AddCategoryRequest) Reset() {
	*x = AddCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCategoryRequest) ProtoMessage() {}

func (x *AddCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCategoryRequest.ProtoReflect.Descriptor instead.
func (*AddCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{1}
}

func (x *AddCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddCategoryRequest) GetParentUuid() int64 {
	if x != nil {
		return x.ParentUuid
	}
	return 0
}

type AddAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddAuthorRequest) Reset() {
	*x = AddAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAuthorRequest) ProtoMessage() {}

func (x *AddAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAuthorRequest.ProtoReflect.Descriptor instead.
func (*AddAuthorRequest) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{2}
}

func (x *AddAuthorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookUuid   int64  `protobuf:"varint,1,opt,name=book_uuid,json=bookUuid,proto3" json:"book_uuid,omitempty"`
	BookName   string `protobuf:"bytes,2,opt,name=book_name,json=bookName,proto3" json:"book_name,omitempty"`
	CategoryId string `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	AuthorId   int64  `protobuf:"varint,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBookRequest) GetBookUuid() int64 {
	if x != nil {
		return x.BookUuid
	}
	return 0
}

func (x *UpdateBookRequest) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

func (x *UpdateBookRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *UpdateBookRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type UpdateAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorUuid int64  `protobuf:"varint,1,opt,name=author_uuid,json=authorUuid,proto3" json:"author_uuid,omitempty"`
	AuthorName string `protobuf:"bytes,2,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
}

func (x *UpdateAuthorRequest) Reset() {
	*x = UpdateAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthorRequest) ProtoMessage() {}

func (x *UpdateAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthorRequest.ProtoReflect.Descriptor instead.
func (*UpdateAuthorRequest) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAuthorRequest) GetAuthorUuid() int64 {
	if x != nil {
		return x.AuthorUuid
	}
	return 0
}

func (x *UpdateAuthorRequest) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

type UpdateCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryUuid int64  `protobuf:"varint,1,opt,name=category_uuid,json=categoryUuid,proto3" json:"category_uuid,omitempty"`
	CategoryName string `protobuf:"bytes,2,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	ParentUuid   int64  `protobuf:"varint,3,opt,name=parent_uuid,json=parentUuid,proto3" json:"parent_uuid,omitempty"`
}

func (x *UpdateCategoryRequest) Reset() {
	*x = UpdateCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCategoryRequest) ProtoMessage() {}

func (x *UpdateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCategoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCategoryRequest) GetCategoryUuid() int64 {
	if x != nil {
		return x.CategoryUuid
	}
	return 0
}

func (x *UpdateCategoryRequest) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *UpdateCategoryRequest) GetParentUuid() int64 {
	if x != nil {
		return x.ParentUuid
	}
	return 0
}

type BookId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookUuid int64 `protobuf:"varint,1,opt,name=book_uuid,json=bookUuid,proto3" json:"book_uuid,omitempty"`
}

func (x *BookId) Reset() {
	*x = BookId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookId) ProtoMessage() {}

func (x *BookId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookId.ProtoReflect.Descriptor instead.
func (*BookId) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{6}
}

func (x *BookId) GetBookUuid() int64 {
	if x != nil {
		return x.BookUuid
	}
	return 0
}

type BookData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *BookData) Reset() {
	*x = BookData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookData) ProtoMessage() {}

func (x *BookData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookData.ProtoReflect.Descriptor instead.
func (*BookData) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{7}
}

func (x *BookData) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type AuthorId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorUuid int64 `protobuf:"varint,1,opt,name=author_uuid,json=authorUuid,proto3" json:"author_uuid,omitempty"`
}

func (x *AuthorId) Reset() {
	*x = AuthorId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorId) ProtoMessage() {}

func (x *AuthorId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorId.ProtoReflect.Descriptor instead.
func (*AuthorId) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{8}
}

func (x *AuthorId) GetAuthorUuid() int64 {
	if x != nil {
		return x.AuthorUuid
	}
	return 0
}

type AuthorData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AuthorData) Reset() {
	*x = AuthorData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorData) ProtoMessage() {}

func (x *AuthorData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorData.ProtoReflect.Descriptor instead.
func (*AuthorData) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{9}
}

func (x *AuthorData) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type CategoryId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryUuid int64 `protobuf:"varint,1,opt,name=category_uuid,json=categoryUuid,proto3" json:"category_uuid,omitempty"`
}

func (x *CategoryId) Reset() {
	*x = CategoryId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryId) ProtoMessage() {}

func (x *CategoryId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryId.ProtoReflect.Descriptor instead.
func (*CategoryId) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{10}
}

func (x *CategoryId) GetCategoryUuid() int64 {
	if x != nil {
		return x.CategoryUuid
	}
	return 0
}

type CategoryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CategoryData) Reset() {
	*x = CategoryData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryData) ProtoMessage() {}

func (x *CategoryData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryData.ProtoReflect.Descriptor instead.
func (*CategoryData) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{11}
}

func (x *CategoryData) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type PageNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int64 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
}

func (x *PageNumber) Reset() {
	*x = PageNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageNumber) ProtoMessage() {}

func (x *PageNumber) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageNumber.ProtoReflect.Descriptor instead.
func (*PageNumber) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{12}
}

func (x *PageNumber) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{13}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_books_proto protoreflect.FileDescriptor

var file_proto_books_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x75,
	0x69, 0x64, 0x22, 0x26, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x75, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x82, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x75, 0x69, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x55, 0x75, 0x69, 0x64, 0x22, 0x22, 0x0a,
	0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x2b, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0x24,
	0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x31, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x55, 0x75, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x2d, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x24,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x32, 0xb2, 0x06, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x1a, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x0e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x1a, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x36, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x53, 0x68, 0x6f,
	0x77, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x53, 0x68, 0x6f,
	0x77, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x34, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_books_proto_rawDescOnce sync.Once
	file_proto_books_proto_rawDescData = file_proto_books_proto_rawDesc
)

func file_proto_books_proto_rawDescGZIP() []byte {
	file_proto_books_proto_rawDescOnce.Do(func() {
		file_proto_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_books_proto_rawDescData)
	})
	return file_proto_books_proto_rawDescData
}

var file_proto_books_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_books_proto_goTypes = []interface{}{
	(*AddBookRequest)(nil),        // 0: proto.AddBookRequest
	(*AddCategoryRequest)(nil),    // 1: proto.AddCategoryRequest
	(*AddAuthorRequest)(nil),      // 2: proto.AddAuthorRequest
	(*UpdateBookRequest)(nil),     // 3: proto.UpdateBookRequest
	(*UpdateAuthorRequest)(nil),   // 4: proto.UpdateAuthorRequest
	(*UpdateCategoryRequest)(nil), // 5: proto.UpdateCategoryRequest
	(*BookId)(nil),                // 6: proto.BookId
	(*BookData)(nil),              // 7: proto.BookData
	(*AuthorId)(nil),              // 8: proto.AuthorId
	(*AuthorData)(nil),            // 9: proto.AuthorData
	(*CategoryId)(nil),            // 10: proto.CategoryId
	(*CategoryData)(nil),          // 11: proto.CategoryData
	(*PageNumber)(nil),            // 12: proto.PageNumber
	(*Response)(nil),              // 13: proto.Response
}
var file_proto_books_proto_depIdxs = []int32{
	0,  // 0: proto.BooksService.AddBook:input_type -> proto.AddBookRequest
	3,  // 1: proto.BooksService.UpdateBook:input_type -> proto.UpdateBookRequest
	6,  // 2: proto.BooksService.DeleteBook:input_type -> proto.BookId
	6,  // 3: proto.BooksService.ShowBook:input_type -> proto.BookId
	8,  // 4: proto.BooksService.FilterByAuthor:input_type -> proto.AuthorId
	10, // 5: proto.BooksService.FilterByCategory:input_type -> proto.CategoryId
	12, // 6: proto.BooksService.Paginate:input_type -> proto.PageNumber
	2,  // 7: proto.BooksService.AddAuthor:input_type -> proto.AddAuthorRequest
	1,  // 8: proto.BooksService.AddCategory:input_type -> proto.AddCategoryRequest
	8,  // 9: proto.BooksService.ShowAuthor:input_type -> proto.AuthorId
	4,  // 10: proto.BooksService.UpdateAuthor:input_type -> proto.UpdateAuthorRequest
	5,  // 11: proto.BooksService.UpdateCategory:input_type -> proto.UpdateCategoryRequest
	10, // 12: proto.BooksService.ShowCategory:input_type -> proto.CategoryId
	10, // 13: proto.BooksService.DeleteCategory:input_type -> proto.CategoryId
	8,  // 14: proto.BooksService.DeleteAuthor:input_type -> proto.AuthorId
	13, // 15: proto.BooksService.AddBook:output_type -> proto.Response
	13, // 16: proto.BooksService.UpdateBook:output_type -> proto.Response
	13, // 17: proto.BooksService.DeleteBook:output_type -> proto.Response
	7,  // 18: proto.BooksService.ShowBook:output_type -> proto.BookData
	7,  // 19: proto.BooksService.FilterByAuthor:output_type -> proto.BookData
	7,  // 20: proto.BooksService.FilterByCategory:output_type -> proto.BookData
	7,  // 21: proto.BooksService.Paginate:output_type -> proto.BookData
	13, // 22: proto.BooksService.AddAuthor:output_type -> proto.Response
	13, // 23: proto.BooksService.AddCategory:output_type -> proto.Response
	9,  // 24: proto.BooksService.ShowAuthor:output_type -> proto.AuthorData
	13, // 25: proto.BooksService.UpdateAuthor:output_type -> proto.Response
	13, // 26: proto.BooksService.UpdateCategory:output_type -> proto.Response
	11, // 27: proto.BooksService.ShowCategory:output_type -> proto.CategoryData
	13, // 28: proto.BooksService.DeleteCategory:output_type -> proto.Response
	13, // 29: proto.BooksService.DeleteAuthor:output_type -> proto.Response
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_books_proto_init() }
func file_proto_books_proto_init() {
	if File_proto_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCategoryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAuthorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_books_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_books_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAuthorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_books_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCategoryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_books_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_books_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_books_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_books_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_books_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_books_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_books_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageNumber); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_books_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_books_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_books_proto_goTypes,
		DependencyIndexes: file_proto_books_proto_depIdxs,
		MessageInfos:      file_proto_books_proto_msgTypes,
	}.Build()
	File_proto_books_proto = out.File
	file_proto_books_proto_rawDesc = nil
	file_proto_books_proto_goTypes = nil
	file_proto_books_proto_depIdxs = nil
}
