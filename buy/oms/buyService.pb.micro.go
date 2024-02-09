// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: buyService.proto

package services

import (
	fmt "fmt"
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

// Api Endpoints for BuyService service

func NewBuyServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for BuyService service

type BuyService interface {
	//确认信息
	Confirm(ctx context.Context, in *BuyRequest, opts ...client.CallOption) (*BuyResponse, error)
	//提交订单
	Submit(ctx context.Context, in *BuyRequest, opts ...client.CallOption) (*BuyResponse, error)
	//订单拆分处理【service】
	OrderSplitHandle(ctx context.Context, in *BuyRequest, opts ...client.CallOption) (*BuyResponse, error)
	//购买商品服务验证【service】
	BuyGoodsCheck(ctx context.Context, in *BuyRequest, opts ...client.CallOption) (*BuyResponse, error)
}

type buyService struct {
	c    client.Client
	name string
}

func NewBuyService(name string, c client.Client) BuyService {
	return &buyService{
		c:    c,
		name: name,
	}
}

func (c *buyService) Confirm(ctx context.Context, in *BuyRequest, opts ...client.CallOption) (*BuyResponse, error) {
	req := c.c.NewRequest(c.name, "BuyService.Confirm", in)
	out := new(BuyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buyService) Submit(ctx context.Context, in *BuyRequest, opts ...client.CallOption) (*BuyResponse, error) {
	req := c.c.NewRequest(c.name, "BuyService.Submit", in)
	out := new(BuyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buyService) OrderSplitHandle(ctx context.Context, in *BuyRequest, opts ...client.CallOption) (*BuyResponse, error) {
	req := c.c.NewRequest(c.name, "BuyService.OrderSplitHandle", in)
	out := new(BuyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buyService) BuyGoodsCheck(ctx context.Context, in *BuyRequest, opts ...client.CallOption) (*BuyResponse, error) {
	req := c.c.NewRequest(c.name, "BuyService.BuyGoodsCheck", in)
	out := new(BuyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BuyService service

type BuyServiceHandler interface {
	//确认信息
	Confirm(context.Context, *BuyRequest, *BuyResponse) error
	//提交订单
	Submit(context.Context, *BuyRequest, *BuyResponse) error
	//订单拆分处理【service】
	OrderSplitHandle(context.Context, *BuyRequest, *BuyResponse) error
	//购买商品服务验证【service】
	BuyGoodsCheck(context.Context, *BuyRequest, *BuyResponse) error
}

func RegisterBuyServiceHandler(s server.Server, hdlr BuyServiceHandler, opts ...server.HandlerOption) error {
	type buyService interface {
		Confirm(ctx context.Context, in *BuyRequest, out *BuyResponse) error
		Submit(ctx context.Context, in *BuyRequest, out *BuyResponse) error
		OrderSplitHandle(ctx context.Context, in *BuyRequest, out *BuyResponse) error
		BuyGoodsCheck(ctx context.Context, in *BuyRequest, out *BuyResponse) error
	}
	type BuyService struct {
		buyService
	}
	h := &buyServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&BuyService{h}, opts...))
}

type buyServiceHandler struct {
	BuyServiceHandler
}

func (h *buyServiceHandler) Confirm(ctx context.Context, in *BuyRequest, out *BuyResponse) error {
	return h.BuyServiceHandler.Confirm(ctx, in, out)
}

func (h *buyServiceHandler) Submit(ctx context.Context, in *BuyRequest, out *BuyResponse) error {
	return h.BuyServiceHandler.Submit(ctx, in, out)
}

func (h *buyServiceHandler) OrderSplitHandle(ctx context.Context, in *BuyRequest, out *BuyResponse) error {
	return h.BuyServiceHandler.OrderSplitHandle(ctx, in, out)
}

func (h *buyServiceHandler) BuyGoodsCheck(ctx context.Context, in *BuyRequest, out *BuyResponse) error {
	return h.BuyServiceHandler.BuyGoodsCheck(ctx, in, out)
}
