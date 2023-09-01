// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: warehouseGoodsService.proto

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

// Api Endpoints for WarehouseGoodsService service

func NewWarehouseGoodsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for WarehouseGoodsService service

type WarehouseGoodsService interface {
	Create(ctx context.Context, in *WarehouseGoods, opts ...client.CallOption) (*WarehouseGoodsResponse, error)
	UpdateInventory(ctx context.Context, in *WarehouseGoods, opts ...client.CallOption) (*WarehouseGoodsResponse, error)
	Delete(ctx context.Context, in *WarehouseGoods, opts ...client.CallOption) (*WarehouseGoodsResponse, error)
	Detail(ctx context.Context, in *WarehouseGoods, opts ...client.CallOption) (*WarehouseGoodsResponse, error)
	List(ctx context.Context, in *WarehouseGoodsRequest, opts ...client.CallOption) (*WarehouseGoodsResponse, error)
	Search(ctx context.Context, in *WarehouseGoodsRequest, opts ...client.CallOption) (*WarehouseGoodsResponse, error)
	SelectGoods(ctx context.Context, in *WarehouseGoodsRequest, opts ...client.CallOption) (*WarehouseGoodsSelectResponse, error)
}

type warehouseGoodsService struct {
	c    client.Client
	name string
}

func NewWarehouseGoodsService(name string, c client.Client) WarehouseGoodsService {
	return &warehouseGoodsService{
		c:    c,
		name: name,
	}
}

func (c *warehouseGoodsService) Create(ctx context.Context, in *WarehouseGoods, opts ...client.CallOption) (*WarehouseGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "WarehouseGoodsService.Create", in)
	out := new(WarehouseGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseGoodsService) UpdateInventory(ctx context.Context, in *WarehouseGoods, opts ...client.CallOption) (*WarehouseGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "WarehouseGoodsService.UpdateInventory", in)
	out := new(WarehouseGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseGoodsService) Delete(ctx context.Context, in *WarehouseGoods, opts ...client.CallOption) (*WarehouseGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "WarehouseGoodsService.Delete", in)
	out := new(WarehouseGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseGoodsService) Detail(ctx context.Context, in *WarehouseGoods, opts ...client.CallOption) (*WarehouseGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "WarehouseGoodsService.Detail", in)
	out := new(WarehouseGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseGoodsService) List(ctx context.Context, in *WarehouseGoodsRequest, opts ...client.CallOption) (*WarehouseGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "WarehouseGoodsService.List", in)
	out := new(WarehouseGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseGoodsService) Search(ctx context.Context, in *WarehouseGoodsRequest, opts ...client.CallOption) (*WarehouseGoodsResponse, error) {
	req := c.c.NewRequest(c.name, "WarehouseGoodsService.Search", in)
	out := new(WarehouseGoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseGoodsService) SelectGoods(ctx context.Context, in *WarehouseGoodsRequest, opts ...client.CallOption) (*WarehouseGoodsSelectResponse, error) {
	req := c.c.NewRequest(c.name, "WarehouseGoodsService.SelectGoods", in)
	out := new(WarehouseGoodsSelectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WarehouseGoodsService service

type WarehouseGoodsServiceHandler interface {
	Create(context.Context, *WarehouseGoods, *WarehouseGoodsResponse) error
	UpdateInventory(context.Context, *WarehouseGoods, *WarehouseGoodsResponse) error
	Delete(context.Context, *WarehouseGoods, *WarehouseGoodsResponse) error
	Detail(context.Context, *WarehouseGoods, *WarehouseGoodsResponse) error
	List(context.Context, *WarehouseGoodsRequest, *WarehouseGoodsResponse) error
	Search(context.Context, *WarehouseGoodsRequest, *WarehouseGoodsResponse) error
	SelectGoods(context.Context, *WarehouseGoodsRequest, *WarehouseGoodsSelectResponse) error
}

func RegisterWarehouseGoodsServiceHandler(s server.Server, hdlr WarehouseGoodsServiceHandler, opts ...server.HandlerOption) error {
	type warehouseGoodsService interface {
		Create(ctx context.Context, in *WarehouseGoods, out *WarehouseGoodsResponse) error
		UpdateInventory(ctx context.Context, in *WarehouseGoods, out *WarehouseGoodsResponse) error
		Delete(ctx context.Context, in *WarehouseGoods, out *WarehouseGoodsResponse) error
		Detail(ctx context.Context, in *WarehouseGoods, out *WarehouseGoodsResponse) error
		List(ctx context.Context, in *WarehouseGoodsRequest, out *WarehouseGoodsResponse) error
		Search(ctx context.Context, in *WarehouseGoodsRequest, out *WarehouseGoodsResponse) error
		SelectGoods(ctx context.Context, in *WarehouseGoodsRequest, out *WarehouseGoodsSelectResponse) error
	}
	type WarehouseGoodsService struct {
		warehouseGoodsService
	}
	h := &warehouseGoodsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WarehouseGoodsService{h}, opts...))
}

type warehouseGoodsServiceHandler struct {
	WarehouseGoodsServiceHandler
}

func (h *warehouseGoodsServiceHandler) Create(ctx context.Context, in *WarehouseGoods, out *WarehouseGoodsResponse) error {
	return h.WarehouseGoodsServiceHandler.Create(ctx, in, out)
}

func (h *warehouseGoodsServiceHandler) UpdateInventory(ctx context.Context, in *WarehouseGoods, out *WarehouseGoodsResponse) error {
	return h.WarehouseGoodsServiceHandler.UpdateInventory(ctx, in, out)
}

func (h *warehouseGoodsServiceHandler) Delete(ctx context.Context, in *WarehouseGoods, out *WarehouseGoodsResponse) error {
	return h.WarehouseGoodsServiceHandler.Delete(ctx, in, out)
}

func (h *warehouseGoodsServiceHandler) Detail(ctx context.Context, in *WarehouseGoods, out *WarehouseGoodsResponse) error {
	return h.WarehouseGoodsServiceHandler.Detail(ctx, in, out)
}

func (h *warehouseGoodsServiceHandler) List(ctx context.Context, in *WarehouseGoodsRequest, out *WarehouseGoodsResponse) error {
	return h.WarehouseGoodsServiceHandler.List(ctx, in, out)
}

func (h *warehouseGoodsServiceHandler) Search(ctx context.Context, in *WarehouseGoodsRequest, out *WarehouseGoodsResponse) error {
	return h.WarehouseGoodsServiceHandler.Search(ctx, in, out)
}

func (h *warehouseGoodsServiceHandler) SelectGoods(ctx context.Context, in *WarehouseGoodsRequest, out *WarehouseGoodsSelectResponse) error {
	return h.WarehouseGoodsServiceHandler.SelectGoods(ctx, in, out)
}
