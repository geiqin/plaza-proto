// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: skuService.proto

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

// Api Endpoints for SpuService service

func NewSpuServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SpuService service

type SpuService interface {
	Create(ctx context.Context, in *Spu, opts ...client.CallOption) (*SkuResponse, error)
	Update(ctx context.Context, in *Spu, opts ...client.CallOption) (*SkuResponse, error)
	Delete(ctx context.Context, in *Spu, opts ...client.CallOption) (*SkuResponse, error)
	Get(ctx context.Context, in *Spu, opts ...client.CallOption) (*SkuResponse, error)
	//商品操作
	Action(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error)
	//商品上下架
	Sale(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error)
	//商品排序
	Sort(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error)
}

type spuService struct {
	c    client.Client
	name string
}

func NewSpuService(name string, c client.Client) SpuService {
	return &spuService{
		c:    c,
		name: name,
	}
}

func (c *spuService) Create(ctx context.Context, in *Spu, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.Create", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) Update(ctx context.Context, in *Spu, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.Update", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) Delete(ctx context.Context, in *Spu, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.Delete", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) Get(ctx context.Context, in *Spu, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.Get", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) Action(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.Action", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) Sale(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.Sale", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spuService) Sort(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SpuService.Sort", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SpuService service

type SpuServiceHandler interface {
	Create(context.Context, *Spu, *SkuResponse) error
	Update(context.Context, *Spu, *SkuResponse) error
	Delete(context.Context, *Spu, *SkuResponse) error
	Get(context.Context, *Spu, *SkuResponse) error
	//商品操作
	Action(context.Context, *SkuRequest, *SkuResponse) error
	//商品上下架
	Sale(context.Context, *SkuRequest, *SkuResponse) error
	//商品排序
	Sort(context.Context, *SkuRequest, *SkuResponse) error
}

func RegisterSpuServiceHandler(s server.Server, hdlr SpuServiceHandler, opts ...server.HandlerOption) error {
	type spuService interface {
		Create(ctx context.Context, in *Spu, out *SkuResponse) error
		Update(ctx context.Context, in *Spu, out *SkuResponse) error
		Delete(ctx context.Context, in *Spu, out *SkuResponse) error
		Get(ctx context.Context, in *Spu, out *SkuResponse) error
		Action(ctx context.Context, in *SkuRequest, out *SkuResponse) error
		Sale(ctx context.Context, in *SkuRequest, out *SkuResponse) error
		Sort(ctx context.Context, in *SkuRequest, out *SkuResponse) error
	}
	type SpuService struct {
		spuService
	}
	h := &spuServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SpuService{h}, opts...))
}

type spuServiceHandler struct {
	SpuServiceHandler
}

func (h *spuServiceHandler) Create(ctx context.Context, in *Spu, out *SkuResponse) error {
	return h.SpuServiceHandler.Create(ctx, in, out)
}

func (h *spuServiceHandler) Update(ctx context.Context, in *Spu, out *SkuResponse) error {
	return h.SpuServiceHandler.Update(ctx, in, out)
}

func (h *spuServiceHandler) Delete(ctx context.Context, in *Spu, out *SkuResponse) error {
	return h.SpuServiceHandler.Delete(ctx, in, out)
}

func (h *spuServiceHandler) Get(ctx context.Context, in *Spu, out *SkuResponse) error {
	return h.SpuServiceHandler.Get(ctx, in, out)
}

func (h *spuServiceHandler) Action(ctx context.Context, in *SkuRequest, out *SkuResponse) error {
	return h.SpuServiceHandler.Action(ctx, in, out)
}

func (h *spuServiceHandler) Sale(ctx context.Context, in *SkuRequest, out *SkuResponse) error {
	return h.SpuServiceHandler.Sale(ctx, in, out)
}

func (h *spuServiceHandler) Sort(ctx context.Context, in *SkuRequest, out *SkuResponse) error {
	return h.SpuServiceHandler.Sort(ctx, in, out)
}

// Api Endpoints for SkuService service

func NewSkuServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SkuService service

type SkuService interface {
	Create(ctx context.Context, in *Sku, opts ...client.CallOption) (*SkuResponse, error)
	Update(ctx context.Context, in *Sku, opts ...client.CallOption) (*SkuResponse, error)
	Delete(ctx context.Context, in *Sku, opts ...client.CallOption) (*SkuResponse, error)
	Get(ctx context.Context, in *Sku, opts ...client.CallOption) (*SkuResponse, error)
	List(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error)
	//货品查询
	Search(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error)
	//商品查询
	GoodsSearch(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error)
	//商品详情显示
	Display(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error)
}

type skuService struct {
	c    client.Client
	name string
}

func NewSkuService(name string, c client.Client) SkuService {
	return &skuService{
		c:    c,
		name: name,
	}
}

func (c *skuService) Create(ctx context.Context, in *Sku, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SkuService.Create", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skuService) Update(ctx context.Context, in *Sku, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SkuService.Update", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skuService) Delete(ctx context.Context, in *Sku, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SkuService.Delete", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skuService) Get(ctx context.Context, in *Sku, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SkuService.Get", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skuService) List(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SkuService.List", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skuService) Search(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SkuService.Search", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skuService) GoodsSearch(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SkuService.GoodsSearch", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skuService) Display(ctx context.Context, in *SkuRequest, opts ...client.CallOption) (*SkuResponse, error) {
	req := c.c.NewRequest(c.name, "SkuService.Display", in)
	out := new(SkuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SkuService service

type SkuServiceHandler interface {
	Create(context.Context, *Sku, *SkuResponse) error
	Update(context.Context, *Sku, *SkuResponse) error
	Delete(context.Context, *Sku, *SkuResponse) error
	Get(context.Context, *Sku, *SkuResponse) error
	List(context.Context, *SkuRequest, *SkuResponse) error
	//货品查询
	Search(context.Context, *SkuRequest, *SkuResponse) error
	//商品查询
	GoodsSearch(context.Context, *SkuRequest, *SkuResponse) error
	//商品详情显示
	Display(context.Context, *SkuRequest, *SkuResponse) error
}

func RegisterSkuServiceHandler(s server.Server, hdlr SkuServiceHandler, opts ...server.HandlerOption) error {
	type skuService interface {
		Create(ctx context.Context, in *Sku, out *SkuResponse) error
		Update(ctx context.Context, in *Sku, out *SkuResponse) error
		Delete(ctx context.Context, in *Sku, out *SkuResponse) error
		Get(ctx context.Context, in *Sku, out *SkuResponse) error
		List(ctx context.Context, in *SkuRequest, out *SkuResponse) error
		Search(ctx context.Context, in *SkuRequest, out *SkuResponse) error
		GoodsSearch(ctx context.Context, in *SkuRequest, out *SkuResponse) error
		Display(ctx context.Context, in *SkuRequest, out *SkuResponse) error
	}
	type SkuService struct {
		skuService
	}
	h := &skuServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SkuService{h}, opts...))
}

type skuServiceHandler struct {
	SkuServiceHandler
}

func (h *skuServiceHandler) Create(ctx context.Context, in *Sku, out *SkuResponse) error {
	return h.SkuServiceHandler.Create(ctx, in, out)
}

func (h *skuServiceHandler) Update(ctx context.Context, in *Sku, out *SkuResponse) error {
	return h.SkuServiceHandler.Update(ctx, in, out)
}

func (h *skuServiceHandler) Delete(ctx context.Context, in *Sku, out *SkuResponse) error {
	return h.SkuServiceHandler.Delete(ctx, in, out)
}

func (h *skuServiceHandler) Get(ctx context.Context, in *Sku, out *SkuResponse) error {
	return h.SkuServiceHandler.Get(ctx, in, out)
}

func (h *skuServiceHandler) List(ctx context.Context, in *SkuRequest, out *SkuResponse) error {
	return h.SkuServiceHandler.List(ctx, in, out)
}

func (h *skuServiceHandler) Search(ctx context.Context, in *SkuRequest, out *SkuResponse) error {
	return h.SkuServiceHandler.Search(ctx, in, out)
}

func (h *skuServiceHandler) GoodsSearch(ctx context.Context, in *SkuRequest, out *SkuResponse) error {
	return h.SkuServiceHandler.GoodsSearch(ctx, in, out)
}

func (h *skuServiceHandler) Display(ctx context.Context, in *SkuRequest, out *SkuResponse) error {
	return h.SkuServiceHandler.Display(ctx, in, out)
}
