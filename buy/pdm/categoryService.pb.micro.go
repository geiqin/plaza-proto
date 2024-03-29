// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: categoryService.proto

package services

import (
	fmt "fmt"
	_ "github.com/geiqin/micro-kit/protobuf/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CategoryService service

func NewCategoryServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CategoryService service

type CategoryService interface {
	Create(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error)
	Update(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error)
	Delete(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error)
	Get(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error)
	Tree(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryResponse, error)
	List(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryResponse, error)
	Search(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryResponse, error)
}

type categoryService struct {
	c    client.Client
	name string
}

func NewCategoryService(name string, c client.Client) CategoryService {
	return &categoryService{
		c:    c,
		name: name,
	}
}

func (c *categoryService) Create(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.Create", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) Update(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.Update", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) Delete(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.Delete", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) Get(ctx context.Context, in *Category, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.Get", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) Tree(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.Tree", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) List(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.List", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) Search(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.Search", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CategoryService service

type CategoryServiceHandler interface {
	Create(context.Context, *Category, *CategoryResponse) error
	Update(context.Context, *Category, *CategoryResponse) error
	Delete(context.Context, *Category, *CategoryResponse) error
	Get(context.Context, *Category, *CategoryResponse) error
	Tree(context.Context, *CategoryRequest, *CategoryResponse) error
	List(context.Context, *CategoryRequest, *CategoryResponse) error
	Search(context.Context, *CategoryRequest, *CategoryResponse) error
}

func RegisterCategoryServiceHandler(s server.Server, hdlr CategoryServiceHandler, opts ...server.HandlerOption) error {
	type categoryService interface {
		Create(ctx context.Context, in *Category, out *CategoryResponse) error
		Update(ctx context.Context, in *Category, out *CategoryResponse) error
		Delete(ctx context.Context, in *Category, out *CategoryResponse) error
		Get(ctx context.Context, in *Category, out *CategoryResponse) error
		Tree(ctx context.Context, in *CategoryRequest, out *CategoryResponse) error
		List(ctx context.Context, in *CategoryRequest, out *CategoryResponse) error
		Search(ctx context.Context, in *CategoryRequest, out *CategoryResponse) error
	}
	type CategoryService struct {
		categoryService
	}
	h := &categoryServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CategoryService{h}, opts...))
}

type categoryServiceHandler struct {
	CategoryServiceHandler
}

func (h *categoryServiceHandler) Create(ctx context.Context, in *Category, out *CategoryResponse) error {
	return h.CategoryServiceHandler.Create(ctx, in, out)
}

func (h *categoryServiceHandler) Update(ctx context.Context, in *Category, out *CategoryResponse) error {
	return h.CategoryServiceHandler.Update(ctx, in, out)
}

func (h *categoryServiceHandler) Delete(ctx context.Context, in *Category, out *CategoryResponse) error {
	return h.CategoryServiceHandler.Delete(ctx, in, out)
}

func (h *categoryServiceHandler) Get(ctx context.Context, in *Category, out *CategoryResponse) error {
	return h.CategoryServiceHandler.Get(ctx, in, out)
}

func (h *categoryServiceHandler) Tree(ctx context.Context, in *CategoryRequest, out *CategoryResponse) error {
	return h.CategoryServiceHandler.Tree(ctx, in, out)
}

func (h *categoryServiceHandler) List(ctx context.Context, in *CategoryRequest, out *CategoryResponse) error {
	return h.CategoryServiceHandler.List(ctx, in, out)
}

func (h *categoryServiceHandler) Search(ctx context.Context, in *CategoryRequest, out *CategoryResponse) error {
	return h.CategoryServiceHandler.Search(ctx, in, out)
}
