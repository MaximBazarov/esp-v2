// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tests/endpoints/bookstore_grpc/proto/bookstore.proto

package endpoints_examples_bookstore

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A shelf resource.
type Shelf struct {
	// A unique shelf id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// A theme of the shelf (fiction, poetry, etc).
	Theme                string   `protobuf:"bytes,2,opt,name=theme,proto3" json:"theme,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Shelf) Reset()         { *m = Shelf{} }
func (m *Shelf) String() string { return proto.CompactTextString(m) }
func (*Shelf) ProtoMessage()    {}
func (*Shelf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2c1a6fc8c4d089, []int{0}
}

func (m *Shelf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Shelf.Unmarshal(m, b)
}
func (m *Shelf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Shelf.Marshal(b, m, deterministic)
}
func (m *Shelf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shelf.Merge(m, src)
}
func (m *Shelf) XXX_Size() int {
	return xxx_messageInfo_Shelf.Size(m)
}
func (m *Shelf) XXX_DiscardUnknown() {
	xxx_messageInfo_Shelf.DiscardUnknown(m)
}

var xxx_messageInfo_Shelf proto.InternalMessageInfo

func (m *Shelf) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Shelf) GetTheme() string {
	if m != nil {
		return m.Theme
	}
	return ""
}

// A book resource.
type Book struct {
	// A unique book id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// An author of the book.
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// A book title.
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2c1a6fc8c4d089, []int{1}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Book) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

// Response to ListShelves call.
type ListShelvesResponse struct {
	// Shelves in the bookstore.
	Shelves              []*Shelf `protobuf:"bytes,1,rep,name=shelves,proto3" json:"shelves,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListShelvesResponse) Reset()         { *m = ListShelvesResponse{} }
func (m *ListShelvesResponse) String() string { return proto.CompactTextString(m) }
func (*ListShelvesResponse) ProtoMessage()    {}
func (*ListShelvesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2c1a6fc8c4d089, []int{2}
}

func (m *ListShelvesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListShelvesResponse.Unmarshal(m, b)
}
func (m *ListShelvesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListShelvesResponse.Marshal(b, m, deterministic)
}
func (m *ListShelvesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListShelvesResponse.Merge(m, src)
}
func (m *ListShelvesResponse) XXX_Size() int {
	return xxx_messageInfo_ListShelvesResponse.Size(m)
}
func (m *ListShelvesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListShelvesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListShelvesResponse proto.InternalMessageInfo

func (m *ListShelvesResponse) GetShelves() []*Shelf {
	if m != nil {
		return m.Shelves
	}
	return nil
}

// Request message for CreateShelf method.
type CreateShelfRequest struct {
	// The shelf resource to create.
	Shelf                *Shelf   `protobuf:"bytes,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateShelfRequest) Reset()         { *m = CreateShelfRequest{} }
func (m *CreateShelfRequest) String() string { return proto.CompactTextString(m) }
func (*CreateShelfRequest) ProtoMessage()    {}
func (*CreateShelfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2c1a6fc8c4d089, []int{3}
}

func (m *CreateShelfRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShelfRequest.Unmarshal(m, b)
}
func (m *CreateShelfRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShelfRequest.Marshal(b, m, deterministic)
}
func (m *CreateShelfRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShelfRequest.Merge(m, src)
}
func (m *CreateShelfRequest) XXX_Size() int {
	return xxx_messageInfo_CreateShelfRequest.Size(m)
}
func (m *CreateShelfRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShelfRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShelfRequest proto.InternalMessageInfo

func (m *CreateShelfRequest) GetShelf() *Shelf {
	if m != nil {
		return m.Shelf
	}
	return nil
}

// Request message for GetShelf method.
type GetShelfRequest struct {
	// The ID of the shelf resource to retrieve.
	Shelf                int64    `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShelfRequest) Reset()         { *m = GetShelfRequest{} }
func (m *GetShelfRequest) String() string { return proto.CompactTextString(m) }
func (*GetShelfRequest) ProtoMessage()    {}
func (*GetShelfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2c1a6fc8c4d089, []int{4}
}

func (m *GetShelfRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShelfRequest.Unmarshal(m, b)
}
func (m *GetShelfRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShelfRequest.Marshal(b, m, deterministic)
}
func (m *GetShelfRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShelfRequest.Merge(m, src)
}
func (m *GetShelfRequest) XXX_Size() int {
	return xxx_messageInfo_GetShelfRequest.Size(m)
}
func (m *GetShelfRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShelfRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShelfRequest proto.InternalMessageInfo

func (m *GetShelfRequest) GetShelf() int64 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

// Request message for DeleteShelf method.
type DeleteShelfRequest struct {
	// The ID of the shelf to delete.
	Shelf                int64    `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteShelfRequest) Reset()         { *m = DeleteShelfRequest{} }
func (m *DeleteShelfRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteShelfRequest) ProtoMessage()    {}
func (*DeleteShelfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2c1a6fc8c4d089, []int{5}
}

func (m *DeleteShelfRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteShelfRequest.Unmarshal(m, b)
}
func (m *DeleteShelfRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteShelfRequest.Marshal(b, m, deterministic)
}
func (m *DeleteShelfRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteShelfRequest.Merge(m, src)
}
func (m *DeleteShelfRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteShelfRequest.Size(m)
}
func (m *DeleteShelfRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteShelfRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteShelfRequest proto.InternalMessageInfo

func (m *DeleteShelfRequest) GetShelf() int64 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

// Request message for ListBooks method.
type ListBooksRequest struct {
	// ID of the shelf which books to list.
	Shelf                int64    `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBooksRequest) Reset()         { *m = ListBooksRequest{} }
func (m *ListBooksRequest) String() string { return proto.CompactTextString(m) }
func (*ListBooksRequest) ProtoMessage()    {}
func (*ListBooksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2c1a6fc8c4d089, []int{6}
}

func (m *ListBooksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBooksRequest.Unmarshal(m, b)
}
func (m *ListBooksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBooksRequest.Marshal(b, m, deterministic)
}
func (m *ListBooksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBooksRequest.Merge(m, src)
}
func (m *ListBooksRequest) XXX_Size() int {
	return xxx_messageInfo_ListBooksRequest.Size(m)
}
func (m *ListBooksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBooksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBooksRequest proto.InternalMessageInfo

func (m *ListBooksRequest) GetShelf() int64 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

// Response message to ListBooks method.
type ListBooksResponse struct {
	// The books on the shelf.
	Books                []*Book  `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBooksResponse) Reset()         { *m = ListBooksResponse{} }
func (m *ListBooksResponse) String() string { return proto.CompactTextString(m) }
func (*ListBooksResponse) ProtoMessage()    {}
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2c1a6fc8c4d089, []int{7}
}

func (m *ListBooksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBooksResponse.Unmarshal(m, b)
}
func (m *ListBooksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBooksResponse.Marshal(b, m, deterministic)
}
func (m *ListBooksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBooksResponse.Merge(m, src)
}
func (m *ListBooksResponse) XXX_Size() int {
	return xxx_messageInfo_ListBooksResponse.Size(m)
}
func (m *ListBooksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBooksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBooksResponse proto.InternalMessageInfo

func (m *ListBooksResponse) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

// Request message for CreateBook method.
type CreateBookRequest struct {
	// The ID of the shelf on which to create a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// A book resource to create on the shelf.
	Book                 *Book    `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBookRequest) Reset()         { *m = CreateBookRequest{} }
func (m *CreateBookRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBookRequest) ProtoMessage()    {}
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2c1a6fc8c4d089, []int{8}
}

func (m *CreateBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBookRequest.Unmarshal(m, b)
}
func (m *CreateBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBookRequest.Marshal(b, m, deterministic)
}
func (m *CreateBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBookRequest.Merge(m, src)
}
func (m *CreateBookRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBookRequest.Size(m)
}
func (m *CreateBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBookRequest proto.InternalMessageInfo

func (m *CreateBookRequest) GetShelf() int64 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

func (m *CreateBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

// Request message for GetBook method.
type GetBookRequest struct {
	// The ID of the shelf from which to retrieve a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// The ID of the book to retrieve.
	Book                 int64    `protobuf:"varint,2,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookRequest) Reset()         { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()    {}
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2c1a6fc8c4d089, []int{9}
}

func (m *GetBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookRequest.Unmarshal(m, b)
}
func (m *GetBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookRequest.Marshal(b, m, deterministic)
}
func (m *GetBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookRequest.Merge(m, src)
}
func (m *GetBookRequest) XXX_Size() int {
	return xxx_messageInfo_GetBookRequest.Size(m)
}
func (m *GetBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookRequest proto.InternalMessageInfo

func (m *GetBookRequest) GetShelf() int64 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

func (m *GetBookRequest) GetBook() int64 {
	if m != nil {
		return m.Book
	}
	return 0
}

// Request message for DeleteBook method.
type DeleteBookRequest struct {
	// The ID of the shelf from which to delete a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// The ID of the book to delete.
	Book                 int64    `protobuf:"varint,2,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBookRequest) Reset()         { *m = DeleteBookRequest{} }
func (m *DeleteBookRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBookRequest) ProtoMessage()    {}
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2c1a6fc8c4d089, []int{10}
}

func (m *DeleteBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBookRequest.Unmarshal(m, b)
}
func (m *DeleteBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBookRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBookRequest.Merge(m, src)
}
func (m *DeleteBookRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBookRequest.Size(m)
}
func (m *DeleteBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBookRequest proto.InternalMessageInfo

func (m *DeleteBookRequest) GetShelf() int64 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

func (m *DeleteBookRequest) GetBook() int64 {
	if m != nil {
		return m.Book
	}
	return 0
}

func init() {
	proto.RegisterType((*Shelf)(nil), "endpoints.examples.bookstore.Shelf")
	proto.RegisterType((*Book)(nil), "endpoints.examples.bookstore.Book")
	proto.RegisterType((*ListShelvesResponse)(nil), "endpoints.examples.bookstore.ListShelvesResponse")
	proto.RegisterType((*CreateShelfRequest)(nil), "endpoints.examples.bookstore.CreateShelfRequest")
	proto.RegisterType((*GetShelfRequest)(nil), "endpoints.examples.bookstore.GetShelfRequest")
	proto.RegisterType((*DeleteShelfRequest)(nil), "endpoints.examples.bookstore.DeleteShelfRequest")
	proto.RegisterType((*ListBooksRequest)(nil), "endpoints.examples.bookstore.ListBooksRequest")
	proto.RegisterType((*ListBooksResponse)(nil), "endpoints.examples.bookstore.ListBooksResponse")
	proto.RegisterType((*CreateBookRequest)(nil), "endpoints.examples.bookstore.CreateBookRequest")
	proto.RegisterType((*GetBookRequest)(nil), "endpoints.examples.bookstore.GetBookRequest")
	proto.RegisterType((*DeleteBookRequest)(nil), "endpoints.examples.bookstore.DeleteBookRequest")
}

func init() {
	proto.RegisterFile("tests/endpoints/bookstore_grpc/proto/bookstore.proto", fileDescriptor_2f2c1a6fc8c4d089)
}

var fileDescriptor_2f2c1a6fc8c4d089 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x95, 0x93, 0xa6, 0x7f, 0x26, 0xfa, 0xf5, 0x47, 0xa7, 0xa5, 0x0a, 0x6e, 0x0f, 0x61, 0x39,
	0x34, 0x8a, 0x88, 0x97, 0x14, 0x84, 0x20, 0xa8, 0x42, 0x0a, 0x45, 0xbd, 0x80, 0xa8, 0x02, 0x27,
	0x6e, 0x4e, 0xb3, 0x49, 0xac, 0x26, 0x5e, 0x93, 0xdd, 0x54, 0xa0, 0x28, 0x02, 0x21, 0x71, 0xe6,
	0xc0, 0xc7, 0xe2, 0xc8, 0x89, 0x3b, 0x1f, 0x04, 0x79, 0x77, 0x9d, 0x98, 0x26, 0xb1, 0x0d, 0xa7,
	0x64, 0xd7, 0x6f, 0xe6, 0xbd, 0x99, 0x7d, 0x33, 0xf0, 0x40, 0x32, 0x21, 0x05, 0x65, 0x7e, 0x27,
	0xe0, 0x9e, 0x2f, 0x05, 0x6d, 0x73, 0x7e, 0x29, 0x24, 0x1f, 0xb1, 0x5a, 0x6f, 0x14, 0x5c, 0xd0,
	0x60, 0xc4, 0x25, 0x9f, 0x5f, 0x3a, 0xea, 0x8c, 0x87, 0x33, 0xbc, 0xc3, 0xde, 0xbb, 0xc3, 0x60,
	0xc0, 0x84, 0x33, 0xc3, 0xd8, 0x87, 0x3d, 0xce, 0x7b, 0x03, 0x46, 0xdd, 0xc0, 0xa3, 0xae, 0xef,
	0x73, 0xe9, 0x4a, 0x8f, 0xfb, 0x42, 0xc7, 0xda, 0x07, 0xe6, 0xab, 0x3a, 0xb5, 0xc7, 0x5d, 0xca,
	0x86, 0x81, 0xfc, 0xa0, 0x3f, 0x92, 0x1a, 0x14, 0x5e, 0xf7, 0xd9, 0xa0, 0x8b, 0xdb, 0x90, 0xf3,
	0x3a, 0x25, 0xab, 0x6c, 0x55, 0xf2, 0xad, 0x9c, 0xd7, 0xc1, 0x3d, 0x28, 0xc8, 0x3e, 0x1b, 0xb2,
	0x52, 0xae, 0x6c, 0x55, 0xb6, 0x5a, 0xfa, 0x40, 0x4e, 0x61, 0xad, 0xc9, 0xf9, 0xe5, 0x02, 0x7a,
	0x1f, 0xd6, 0xdd, 0xb1, 0xec, 0xf3, 0x91, 0x81, 0x9b, 0x93, 0xca, 0xe2, 0xc9, 0x01, 0x2b, 0xe5,
	0x4d, 0x96, 0xf0, 0x40, 0xde, 0xc0, 0xee, 0x0b, 0x4f, 0xc8, 0x90, 0xf8, 0x8a, 0x89, 0x16, 0x13,
	0x01, 0xf7, 0x05, 0xc3, 0x13, 0xd8, 0x10, 0xfa, 0xaa, 0x64, 0x95, 0xf3, 0x95, 0xe2, 0xf1, 0x1d,
	0x27, 0xa9, 0x6c, 0x47, 0x09, 0x6f, 0x45, 0x31, 0xe4, 0x15, 0xe0, 0xb3, 0x11, 0x73, 0x25, 0xd3,
	0xf7, 0xec, 0xdd, 0x98, 0x09, 0x89, 0x8f, 0xa1, 0x10, 0x02, 0xba, 0x4a, 0x6c, 0xc6, 0x94, 0x3a,
	0x82, 0x1c, 0xc1, 0xff, 0x67, 0x4c, 0xfe, 0x91, 0x6d, 0x2f, 0x9e, 0x2d, 0x1f, 0x01, 0xab, 0x80,
	0xa7, 0x6c, 0xc0, 0xae, 0x31, 0x2f, 0xc7, 0x56, 0xe0, 0x46, 0x58, 0x7b, 0xd8, 0x45, 0x91, 0x8c,
	0x7c, 0x09, 0x3b, 0x31, 0xa4, 0xe9, 0xd1, 0x23, 0x28, 0x28, 0xb5, 0xa6, 0x43, 0x24, 0xb9, 0x9c,
	0x30, 0xb6, 0xa5, 0x03, 0x88, 0x0b, 0x3b, 0xba, 0x3d, 0xea, 0x32, 0x89, 0x19, 0x1f, 0xc2, 0x5a,
	0x18, 0xa3, 0xde, 0x32, 0x1b, 0x87, 0xc2, 0x93, 0x06, 0x6c, 0x9f, 0x31, 0x99, 0x9e, 0x1f, 0x63,
	0xf9, 0xf3, 0x26, 0xf6, 0x04, 0x76, 0x74, 0x0f, 0xff, 0x29, 0xfc, 0xf8, 0xe7, 0x26, 0x6c, 0x35,
	0x23, 0x4d, 0x38, 0x84, 0x62, 0xcc, 0x60, 0xb8, 0xef, 0xe8, 0x11, 0x70, 0xa2, 0x11, 0x70, 0x9e,
	0x87, 0x23, 0x60, 0xd7, 0x93, 0x2b, 0x5b, 0xe2, 0x51, 0xb2, 0xfb, 0xf9, 0xc7, 0xaf, 0x6f, 0xb9,
	0xff, 0xb0, 0x48, 0xaf, 0xea, 0xd4, 0x38, 0x0f, 0x3f, 0x59, 0x50, 0x8c, 0x59, 0x0f, 0xef, 0x25,
	0xe7, 0x5d, 0x74, 0xa9, 0x9d, 0xc5, 0x96, 0xc4, 0x56, 0xdc, 0x7b, 0x24, 0xce, 0xdd, 0x30, 0x3d,
	0x99, 0xc0, 0x66, 0xe4, 0x55, 0xac, 0x25, 0x27, 0xbb, 0xe6, 0xe9, 0x6c, 0xdc, 0x07, 0x8a, 0xfb,
	0x26, 0xee, 0xc6, 0xb8, 0xe9, 0x44, 0x71, 0x4f, 0x51, 0x42, 0x31, 0xe6, 0xff, 0xb4, 0xf2, 0x17,
	0x47, 0xc5, 0x5e, 0xf1, 0x40, 0x11, 0x6b, 0x75, 0x29, 0xeb, 0x57, 0x0b, 0xb6, 0x66, 0x03, 0x82,
	0x4e, 0xfa, 0x5b, 0xc6, 0x67, 0xce, 0xa6, 0x99, 0xf1, 0xe6, 0xe5, 0x6f, 0x2b, 0x2d, 0x07, 0x78,
	0x6b, 0x89, 0x16, 0xbd, 0xaf, 0xf1, 0xbb, 0x05, 0x30, 0x9f, 0x31, 0xa4, 0x59, 0x6c, 0x10, 0xb3,
	0xbb, 0x9d, 0x61, 0xd2, 0x48, 0x5f, 0xc9, 0x68, 0x93, 0xd5, 0x32, 0x1a, 0x6a, 0x22, 0xde, 0x3e,
	0x25, 0xf5, 0x95, 0x10, 0x3a, 0x09, 0x7f, 0x1c, 0xaf, 0x33, 0x35, 0xff, 0xf4, 0xba, 0x9e, 0x36,
	0x40, 0x9d, 0xd4, 0x96, 0xc6, 0x2f, 0x16, 0x6c, 0x98, 0x71, 0xc6, 0xbb, 0xa9, 0x96, 0xfa, 0xdb,
	0x3a, 0x2a, 0xaa, 0x0e, 0x82, 0xe5, 0x14, 0x91, 0x53, 0xfc, 0x08, 0x30, 0xdf, 0x0c, 0x69, 0x4d,
	0x5d, 0xd8, 0x21, 0x2b, 0xbd, 0x65, 0x04, 0x54, 0x53, 0x05, 0x34, 0x9f, 0xc0, 0xd1, 0x05, 0x1f,
	0x46, 0x69, 0x92, 0xe8, 0x9b, 0xdb, 0xb3, 0x1d, 0x74, 0x1e, 0xb2, 0x9d, 0x5b, 0xed, 0x75, 0x45,
	0x7b, 0xff, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x57, 0x1c, 0xb0, 0xf8, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookstoreClient is the client API for Bookstore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookstoreClient interface {
	// Returns a list of all shelves in the bookstore.
	ListShelves(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListShelvesResponse, error)
	// Creates a new shelf in the bookstore.
	CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*Shelf, error)
	// Returns a specific bookstore shelf.
	GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error)
	// Deletes a shelf, including all books that are stored on the shelf.
	DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Returns a list of books on a shelf.
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	// Creates a new book.
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Returns a specific book.
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Deletes a book from a shelf.
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type bookstoreClient struct {
	cc *grpc.ClientConn
}

func NewBookstoreClient(cc *grpc.ClientConn) BookstoreClient {
	return &bookstoreClient{cc}
}

func (c *bookstoreClient) ListShelves(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListShelvesResponse, error) {
	out := new(ListShelvesResponse)
	err := c.cc.Invoke(ctx, "/endpoints.examples.bookstore.Bookstore/ListShelves", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := c.cc.Invoke(ctx, "/endpoints.examples.bookstore.Bookstore/CreateShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := c.cc.Invoke(ctx, "/endpoints.examples.bookstore.Bookstore/GetShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/endpoints.examples.bookstore.Bookstore/DeleteShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := c.cc.Invoke(ctx, "/endpoints.examples.bookstore.Bookstore/ListBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/endpoints.examples.bookstore.Bookstore/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/endpoints.examples.bookstore.Bookstore/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/endpoints.examples.bookstore.Bookstore/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookstoreServer is the server API for Bookstore service.
type BookstoreServer interface {
	// Returns a list of all shelves in the bookstore.
	ListShelves(context.Context, *empty.Empty) (*ListShelvesResponse, error)
	// Creates a new shelf in the bookstore.
	CreateShelf(context.Context, *CreateShelfRequest) (*Shelf, error)
	// Returns a specific bookstore shelf.
	GetShelf(context.Context, *GetShelfRequest) (*Shelf, error)
	// Deletes a shelf, including all books that are stored on the shelf.
	DeleteShelf(context.Context, *DeleteShelfRequest) (*empty.Empty, error)
	// Returns a list of books on a shelf.
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	// Creates a new book.
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	// Returns a specific book.
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	// Deletes a book from a shelf.
	DeleteBook(context.Context, *DeleteBookRequest) (*empty.Empty, error)
}

// UnimplementedBookstoreServer can be embedded to have forward compatible implementations.
type UnimplementedBookstoreServer struct {
}

func (*UnimplementedBookstoreServer) ListShelves(ctx context.Context, req *empty.Empty) (*ListShelvesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShelves not implemented")
}
func (*UnimplementedBookstoreServer) CreateShelf(ctx context.Context, req *CreateShelfRequest) (*Shelf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShelf not implemented")
}
func (*UnimplementedBookstoreServer) GetShelf(ctx context.Context, req *GetShelfRequest) (*Shelf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShelf not implemented")
}
func (*UnimplementedBookstoreServer) DeleteShelf(ctx context.Context, req *DeleteShelfRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShelf not implemented")
}
func (*UnimplementedBookstoreServer) ListBooks(ctx context.Context, req *ListBooksRequest) (*ListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (*UnimplementedBookstoreServer) CreateBook(ctx context.Context, req *CreateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (*UnimplementedBookstoreServer) GetBook(ctx context.Context, req *GetBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (*UnimplementedBookstoreServer) DeleteBook(ctx context.Context, req *DeleteBookRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}

func RegisterBookstoreServer(s *grpc.Server, srv BookstoreServer) {
	s.RegisterService(&_Bookstore_serviceDesc, srv)
}

func _Bookstore_ListShelves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).ListShelves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.examples.bookstore.Bookstore/ListShelves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).ListShelves(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_CreateShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreateShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.examples.bookstore.Bookstore/CreateShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreateShelf(ctx, req.(*CreateShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.examples.bookstore.Bookstore/GetShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetShelf(ctx, req.(*GetShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_DeleteShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).DeleteShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.examples.bookstore.Bookstore/DeleteShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).DeleteShelf(ctx, req.(*DeleteShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.examples.bookstore.Bookstore/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.examples.bookstore.Bookstore/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.examples.bookstore.Bookstore/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.examples.bookstore.Bookstore/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bookstore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "endpoints.examples.bookstore.Bookstore",
	HandlerType: (*BookstoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListShelves",
			Handler:    _Bookstore_ListShelves_Handler,
		},
		{
			MethodName: "CreateShelf",
			Handler:    _Bookstore_CreateShelf_Handler,
		},
		{
			MethodName: "GetShelf",
			Handler:    _Bookstore_GetShelf_Handler,
		},
		{
			MethodName: "DeleteShelf",
			Handler:    _Bookstore_DeleteShelf_Handler,
		},
		{
			MethodName: "ListBooks",
			Handler:    _Bookstore_ListBooks_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _Bookstore_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Bookstore_GetBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _Bookstore_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tests/endpoints/bookstore_grpc/proto/bookstore.proto",
}
