// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderService.proto

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

// Api Endpoints for OrderService service

func NewOrderServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderService service

type OrderService interface {
	//订单插入【service】
	OrderInsert(ctx context.Context, in *Order, opts ...client.CallOption) (*OrderResponse, error)
	//订单取消
	OrderCancel(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单删除
	OrderDelete(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单确认
	OrderConfirm(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//修改订单价格
	ModifyPrice(ctx context.Context, in *Order, opts ...client.CallOption) (*OrderResponse, error)
	//修改订单地址
	ModifyAddress(ctx context.Context, in *OrderAddress, opts ...client.CallOption) (*OrderResponse, error)
	//订单主信息列表(服务专用)
	OrderMainInfo(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单主信息(服务专用)
	OrderMainList(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单查询
	Search(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单详情
	Detail(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//支付同步处理
	Respond(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//支付异步处理
	Notify(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单获取【service】
	Get(ctx context.Context, in *Order, opts ...client.CallOption) (*OrderResponse, error)
	//订单列表【service】
	List(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//新订单通知数据(循环模式：将弃用)
	NewOrderNotice(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) OrderInsert(ctx context.Context, in *Order, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderInsert", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderCancel(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderCancel", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderDelete(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderDelete", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderConfirm(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderConfirm", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) ModifyPrice(ctx context.Context, in *Order, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.ModifyPrice", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) ModifyAddress(ctx context.Context, in *OrderAddress, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.ModifyAddress", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderMainInfo(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderMainInfo", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderMainList(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderMainList", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Search(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Search", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Detail(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Detail", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Respond(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Respond", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Notify(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Notify", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Get(ctx context.Context, in *Order, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.Get", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) List(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.List", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) NewOrderNotice(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.NewOrderNotice", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceHandler interface {
	//订单插入【service】
	OrderInsert(context.Context, *Order, *OrderResponse) error
	//订单取消
	OrderCancel(context.Context, *OrderRequest, *OrderResponse) error
	//订单删除
	OrderDelete(context.Context, *OrderRequest, *OrderResponse) error
	//订单确认
	OrderConfirm(context.Context, *OrderRequest, *OrderResponse) error
	//修改订单价格
	ModifyPrice(context.Context, *Order, *OrderResponse) error
	//修改订单地址
	ModifyAddress(context.Context, *OrderAddress, *OrderResponse) error
	//订单主信息列表(服务专用)
	OrderMainInfo(context.Context, *OrderRequest, *OrderResponse) error
	//订单主信息(服务专用)
	OrderMainList(context.Context, *OrderRequest, *OrderResponse) error
	//订单查询
	Search(context.Context, *OrderRequest, *OrderResponse) error
	//订单详情
	Detail(context.Context, *OrderRequest, *OrderResponse) error
	//支付同步处理
	Respond(context.Context, *OrderRequest, *OrderResponse) error
	//支付异步处理
	Notify(context.Context, *OrderRequest, *OrderResponse) error
	//订单获取【service】
	Get(context.Context, *Order, *OrderResponse) error
	//订单列表【service】
	List(context.Context, *OrderRequest, *OrderResponse) error
	//新订单通知数据(循环模式：将弃用)
	NewOrderNotice(context.Context, *OrderRequest, *OrderResponse) error
}

func RegisterOrderServiceHandler(s server.Server, hdlr OrderServiceHandler, opts ...server.HandlerOption) error {
	type orderService interface {
		OrderInsert(ctx context.Context, in *Order, out *OrderResponse) error
		OrderCancel(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		OrderDelete(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		OrderConfirm(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		ModifyPrice(ctx context.Context, in *Order, out *OrderResponse) error
		ModifyAddress(ctx context.Context, in *OrderAddress, out *OrderResponse) error
		OrderMainInfo(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		OrderMainList(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		Search(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		Detail(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		Respond(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		Notify(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		Get(ctx context.Context, in *Order, out *OrderResponse) error
		List(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		NewOrderNotice(ctx context.Context, in *OrderRequest, out *OrderResponse) error
	}
	type OrderService struct {
		orderService
	}
	h := &orderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderService{h}, opts...))
}

type orderServiceHandler struct {
	OrderServiceHandler
}

func (h *orderServiceHandler) OrderInsert(ctx context.Context, in *Order, out *OrderResponse) error {
	return h.OrderServiceHandler.OrderInsert(ctx, in, out)
}

func (h *orderServiceHandler) OrderCancel(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.OrderCancel(ctx, in, out)
}

func (h *orderServiceHandler) OrderDelete(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.OrderDelete(ctx, in, out)
}

func (h *orderServiceHandler) OrderConfirm(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.OrderConfirm(ctx, in, out)
}

func (h *orderServiceHandler) ModifyPrice(ctx context.Context, in *Order, out *OrderResponse) error {
	return h.OrderServiceHandler.ModifyPrice(ctx, in, out)
}

func (h *orderServiceHandler) ModifyAddress(ctx context.Context, in *OrderAddress, out *OrderResponse) error {
	return h.OrderServiceHandler.ModifyAddress(ctx, in, out)
}

func (h *orderServiceHandler) OrderMainInfo(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.OrderMainInfo(ctx, in, out)
}

func (h *orderServiceHandler) OrderMainList(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.OrderMainList(ctx, in, out)
}

func (h *orderServiceHandler) Search(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.Search(ctx, in, out)
}

func (h *orderServiceHandler) Detail(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.Detail(ctx, in, out)
}

func (h *orderServiceHandler) Respond(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.Respond(ctx, in, out)
}

func (h *orderServiceHandler) Notify(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.Notify(ctx, in, out)
}

func (h *orderServiceHandler) Get(ctx context.Context, in *Order, out *OrderResponse) error {
	return h.OrderServiceHandler.Get(ctx, in, out)
}

func (h *orderServiceHandler) List(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.List(ctx, in, out)
}

func (h *orderServiceHandler) NewOrderNotice(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.NewOrderNotice(ctx, in, out)
}
