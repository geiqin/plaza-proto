// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: foodService.proto

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

// Api Endpoints for FoodService service

func NewFoodServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FoodService service

type FoodService interface {
	Create(ctx context.Context, in *Item, opts ...client.CallOption) (*FoodResponse, error)
	Update(ctx context.Context, in *Item, opts ...client.CallOption) (*FoodResponse, error)
	Delete(ctx context.Context, in *Item, opts ...client.CallOption) (*FoodResponse, error)
	Get(ctx context.Context, in *Item, opts ...client.CallOption) (*FoodResponse, error)
	Display(ctx context.Context, in *Item, opts ...client.CallOption) (*FoodResponse, error)
	Search(ctx context.Context, in *FoodRequest, opts ...client.CallOption) (*ItemResponse, error)
	FrontSearch(ctx context.Context, in *FoodRequest, opts ...client.CallOption) (*ItemResponse, error)
}

type foodService struct {
	c    client.Client
	name string
}

func NewFoodService(name string, c client.Client) FoodService {
	return &foodService{
		c:    c,
		name: name,
	}
}

func (c *foodService) Create(ctx context.Context, in *Item, opts ...client.CallOption) (*FoodResponse, error) {
	req := c.c.NewRequest(c.name, "FoodService.Create", in)
	out := new(FoodResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodService) Update(ctx context.Context, in *Item, opts ...client.CallOption) (*FoodResponse, error) {
	req := c.c.NewRequest(c.name, "FoodService.Update", in)
	out := new(FoodResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodService) Delete(ctx context.Context, in *Item, opts ...client.CallOption) (*FoodResponse, error) {
	req := c.c.NewRequest(c.name, "FoodService.Delete", in)
	out := new(FoodResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodService) Get(ctx context.Context, in *Item, opts ...client.CallOption) (*FoodResponse, error) {
	req := c.c.NewRequest(c.name, "FoodService.Get", in)
	out := new(FoodResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodService) Display(ctx context.Context, in *Item, opts ...client.CallOption) (*FoodResponse, error) {
	req := c.c.NewRequest(c.name, "FoodService.Display", in)
	out := new(FoodResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodService) Search(ctx context.Context, in *FoodRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "FoodService.Search", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodService) FrontSearch(ctx context.Context, in *FoodRequest, opts ...client.CallOption) (*ItemResponse, error) {
	req := c.c.NewRequest(c.name, "FoodService.FrontSearch", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FoodService service

type FoodServiceHandler interface {
	Create(context.Context, *Item, *FoodResponse) error
	Update(context.Context, *Item, *FoodResponse) error
	Delete(context.Context, *Item, *FoodResponse) error
	Get(context.Context, *Item, *FoodResponse) error
	Display(context.Context, *Item, *FoodResponse) error
	Search(context.Context, *FoodRequest, *ItemResponse) error
	FrontSearch(context.Context, *FoodRequest, *ItemResponse) error
}

func RegisterFoodServiceHandler(s server.Server, hdlr FoodServiceHandler, opts ...server.HandlerOption) error {
	type foodService interface {
		Create(ctx context.Context, in *Item, out *FoodResponse) error
		Update(ctx context.Context, in *Item, out *FoodResponse) error
		Delete(ctx context.Context, in *Item, out *FoodResponse) error
		Get(ctx context.Context, in *Item, out *FoodResponse) error
		Display(ctx context.Context, in *Item, out *FoodResponse) error
		Search(ctx context.Context, in *FoodRequest, out *ItemResponse) error
		FrontSearch(ctx context.Context, in *FoodRequest, out *ItemResponse) error
	}
	type FoodService struct {
		foodService
	}
	h := &foodServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FoodService{h}, opts...))
}

type foodServiceHandler struct {
	FoodServiceHandler
}

func (h *foodServiceHandler) Create(ctx context.Context, in *Item, out *FoodResponse) error {
	return h.FoodServiceHandler.Create(ctx, in, out)
}

func (h *foodServiceHandler) Update(ctx context.Context, in *Item, out *FoodResponse) error {
	return h.FoodServiceHandler.Update(ctx, in, out)
}

func (h *foodServiceHandler) Delete(ctx context.Context, in *Item, out *FoodResponse) error {
	return h.FoodServiceHandler.Delete(ctx, in, out)
}

func (h *foodServiceHandler) Get(ctx context.Context, in *Item, out *FoodResponse) error {
	return h.FoodServiceHandler.Get(ctx, in, out)
}

func (h *foodServiceHandler) Display(ctx context.Context, in *Item, out *FoodResponse) error {
	return h.FoodServiceHandler.Display(ctx, in, out)
}

func (h *foodServiceHandler) Search(ctx context.Context, in *FoodRequest, out *ItemResponse) error {
	return h.FoodServiceHandler.Search(ctx, in, out)
}

func (h *foodServiceHandler) FrontSearch(ctx context.Context, in *FoodRequest, out *ItemResponse) error {
	return h.FoodServiceHandler.FrontSearch(ctx, in, out)
}
