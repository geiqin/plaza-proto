// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: goodsService.proto

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

// Api Endpoints for MyGoodsService service

func NewMyGoodsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MyGoodsService service

type MyGoodsService interface {
	// 分页查询分销商品列表
	Search(ctx context.Context, in *GoodsWhere, opts ...client.CallOption) (*GoodsInfoResponse, error)
}

type myGoodsService struct {
	c    client.Client
	name string
}

func NewMyGoodsService(name string, c client.Client) MyGoodsService {
	return &myGoodsService{
		c:    c,
		name: name,
	}
}

func (c *myGoodsService) Search(ctx context.Context, in *GoodsWhere, opts ...client.CallOption) (*GoodsInfoResponse, error) {
	req := c.c.NewRequest(c.name, "MyGoodsService.Search", in)
	out := new(GoodsInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyGoodsService service

type MyGoodsServiceHandler interface {
	// 分页查询分销商品列表
	Search(context.Context, *GoodsWhere, *GoodsInfoResponse) error
}

func RegisterMyGoodsServiceHandler(s server.Server, hdlr MyGoodsServiceHandler, opts ...server.HandlerOption) error {
	type myGoodsService interface {
		Search(ctx context.Context, in *GoodsWhere, out *GoodsInfoResponse) error
	}
	type MyGoodsService struct {
		myGoodsService
	}
	h := &myGoodsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MyGoodsService{h}, opts...))
}

type myGoodsServiceHandler struct {
	MyGoodsServiceHandler
}

func (h *myGoodsServiceHandler) Search(ctx context.Context, in *GoodsWhere, out *GoodsInfoResponse) error {
	return h.MyGoodsServiceHandler.Search(ctx, in, out)
}

// Api Endpoints for GoodsService service

func NewGoodsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GoodsService service

type GoodsService interface {
	//分页查询分销商品列表
	Search(ctx context.Context, in *GoodsWhere, opts ...client.CallOption) (*GoodsInfoResponse, error)
	//商品取消/参与推广
	Disabled(ctx context.Context, in *GoodsWhere, opts ...client.CallOption) (*GoodsResponse, error)
	// 检查商品是否参与分销
	Check(ctx context.Context, in *Goods, opts ...client.CallOption) (*GoodsResponse, error)
	//获取分销商品信息
	Get(ctx context.Context, in *Goods, opts ...client.CallOption) (*GoodsResponse, error)
	//设置分销商品信息
	Set(ctx context.Context, in *Goods, opts ...client.CallOption) (*GoodsResponse, error)
}

type goodsService struct {
	c    client.Client
	name string
}

func NewGoodsService(name string, c client.Client) GoodsService {
	return &goodsService{
		c:    c,
		name: name,
	}
}

func (c *goodsService) Search(ctx context.Context, in *GoodsWhere, opts ...client.CallOption) (*GoodsInfoResponse, error) {
	req := c.c.NewRequest(c.name, "GoodsService.Search", in)
	out := new(GoodsInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsService) Disabled(ctx context.Context, in *GoodsWhere, opts ...client.CallOption) (*GoodsResponse, error) {
	req := c.c.NewRequest(c.name, "GoodsService.Disabled", in)
	out := new(GoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsService) Check(ctx context.Context, in *Goods, opts ...client.CallOption) (*GoodsResponse, error) {
	req := c.c.NewRequest(c.name, "GoodsService.Check", in)
	out := new(GoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsService) Get(ctx context.Context, in *Goods, opts ...client.CallOption) (*GoodsResponse, error) {
	req := c.c.NewRequest(c.name, "GoodsService.Get", in)
	out := new(GoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsService) Set(ctx context.Context, in *Goods, opts ...client.CallOption) (*GoodsResponse, error) {
	req := c.c.NewRequest(c.name, "GoodsService.Set", in)
	out := new(GoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoodsService service

type GoodsServiceHandler interface {
	//分页查询分销商品列表
	Search(context.Context, *GoodsWhere, *GoodsInfoResponse) error
	//商品取消/参与推广
	Disabled(context.Context, *GoodsWhere, *GoodsResponse) error
	// 检查商品是否参与分销
	Check(context.Context, *Goods, *GoodsResponse) error
	//获取分销商品信息
	Get(context.Context, *Goods, *GoodsResponse) error
	//设置分销商品信息
	Set(context.Context, *Goods, *GoodsResponse) error
}

func RegisterGoodsServiceHandler(s server.Server, hdlr GoodsServiceHandler, opts ...server.HandlerOption) error {
	type goodsService interface {
		Search(ctx context.Context, in *GoodsWhere, out *GoodsInfoResponse) error
		Disabled(ctx context.Context, in *GoodsWhere, out *GoodsResponse) error
		Check(ctx context.Context, in *Goods, out *GoodsResponse) error
		Get(ctx context.Context, in *Goods, out *GoodsResponse) error
		Set(ctx context.Context, in *Goods, out *GoodsResponse) error
	}
	type GoodsService struct {
		goodsService
	}
	h := &goodsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GoodsService{h}, opts...))
}

type goodsServiceHandler struct {
	GoodsServiceHandler
}

func (h *goodsServiceHandler) Search(ctx context.Context, in *GoodsWhere, out *GoodsInfoResponse) error {
	return h.GoodsServiceHandler.Search(ctx, in, out)
}

func (h *goodsServiceHandler) Disabled(ctx context.Context, in *GoodsWhere, out *GoodsResponse) error {
	return h.GoodsServiceHandler.Disabled(ctx, in, out)
}

func (h *goodsServiceHandler) Check(ctx context.Context, in *Goods, out *GoodsResponse) error {
	return h.GoodsServiceHandler.Check(ctx, in, out)
}

func (h *goodsServiceHandler) Get(ctx context.Context, in *Goods, out *GoodsResponse) error {
	return h.GoodsServiceHandler.Get(ctx, in, out)
}

func (h *goodsServiceHandler) Set(ctx context.Context, in *Goods, out *GoodsResponse) error {
	return h.GoodsServiceHandler.Set(ctx, in, out)
}
