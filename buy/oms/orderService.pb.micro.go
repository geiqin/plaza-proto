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
	//订单支付【user】
	OrderPay(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderPayResponse, error)
	//订单线下支付【admin】
	OrderUnderLinePay(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单发货【admin】
	OrderDelivery(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单取消【admin/user】
	OrderCancel(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单删除【admin/user】
	OrderDelete(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单确认【admin】
	OrderConfirm(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单收货【user】
	OrderCollect(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//支付状态校验
	OrderPayCheck(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//修改订单价格【admin】
	ModifyOrderPrice(ctx context.Context, in *Order, opts ...client.CallOption) (*OrderResponse, error)
	//修改订单地址【admin】
	ModifyOrderAddress(ctx context.Context, in *OrderAddress, opts ...client.CallOption) (*OrderResponse, error)
	//订单每个环节状态总数
	OrderStatusStepTotal(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderStatusStepTotalResponse, error)
	//订单总数
	OrderTotal(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单查询【admin/user】
	OrderSearch(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单详情【admin/user】
	OrderDetail(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderDetailResponse, error)
	//支付同步处理
	Respond(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//支付异步处理
	Notify(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//订单获取【service】
	Get(ctx context.Context, in *Order, opts ...client.CallOption) (*OrderResponse, error)
	//订单列表【service】
	List(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	//扩展列表【service】
	OrderExtensionList(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderExtensionResponse, error)
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

func (c *orderService) OrderPay(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderPayResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderPay", in)
	out := new(OrderPayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderUnderLinePay(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderUnderLinePay", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderDelivery(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderDelivery", in)
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

func (c *orderService) OrderCollect(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderCollect", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderPayCheck(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderPayCheck", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) ModifyOrderPrice(ctx context.Context, in *Order, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.ModifyOrderPrice", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) ModifyOrderAddress(ctx context.Context, in *OrderAddress, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.ModifyOrderAddress", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderStatusStepTotal(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderStatusStepTotalResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderStatusStepTotal", in)
	out := new(OrderStatusStepTotalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderTotal(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderTotal", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderSearch(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderSearch", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderDetail(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderDetailResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderDetail", in)
	out := new(OrderDetailResponse)
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

func (c *orderService) OrderExtensionList(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderExtensionResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.OrderExtensionList", in)
	out := new(OrderExtensionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceHandler interface {
	//订单支付【user】
	OrderPay(context.Context, *OrderRequest, *OrderPayResponse) error
	//订单线下支付【admin】
	OrderUnderLinePay(context.Context, *OrderRequest, *OrderResponse) error
	//订单发货【admin】
	OrderDelivery(context.Context, *OrderRequest, *OrderResponse) error
	//订单取消【admin/user】
	OrderCancel(context.Context, *OrderRequest, *OrderResponse) error
	//订单删除【admin/user】
	OrderDelete(context.Context, *OrderRequest, *OrderResponse) error
	//订单确认【admin】
	OrderConfirm(context.Context, *OrderRequest, *OrderResponse) error
	//订单收货【user】
	OrderCollect(context.Context, *OrderRequest, *OrderResponse) error
	//支付状态校验
	OrderPayCheck(context.Context, *OrderRequest, *OrderResponse) error
	//修改订单价格【admin】
	ModifyOrderPrice(context.Context, *Order, *OrderResponse) error
	//修改订单地址【admin】
	ModifyOrderAddress(context.Context, *OrderAddress, *OrderResponse) error
	//订单每个环节状态总数
	OrderStatusStepTotal(context.Context, *OrderRequest, *OrderStatusStepTotalResponse) error
	//订单总数
	OrderTotal(context.Context, *OrderRequest, *OrderResponse) error
	//订单查询【admin/user】
	OrderSearch(context.Context, *OrderRequest, *OrderResponse) error
	//订单详情【admin/user】
	OrderDetail(context.Context, *OrderRequest, *OrderDetailResponse) error
	//支付同步处理
	Respond(context.Context, *OrderRequest, *OrderResponse) error
	//支付异步处理
	Notify(context.Context, *OrderRequest, *OrderResponse) error
	//订单获取【service】
	Get(context.Context, *Order, *OrderResponse) error
	//订单列表【service】
	List(context.Context, *OrderRequest, *OrderResponse) error
	//扩展列表【service】
	OrderExtensionList(context.Context, *OrderRequest, *OrderExtensionResponse) error
}

func RegisterOrderServiceHandler(s server.Server, hdlr OrderServiceHandler, opts ...server.HandlerOption) error {
	type orderService interface {
		OrderPay(ctx context.Context, in *OrderRequest, out *OrderPayResponse) error
		OrderUnderLinePay(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		OrderDelivery(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		OrderCancel(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		OrderDelete(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		OrderConfirm(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		OrderCollect(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		OrderPayCheck(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		ModifyOrderPrice(ctx context.Context, in *Order, out *OrderResponse) error
		ModifyOrderAddress(ctx context.Context, in *OrderAddress, out *OrderResponse) error
		OrderStatusStepTotal(ctx context.Context, in *OrderRequest, out *OrderStatusStepTotalResponse) error
		OrderTotal(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		OrderSearch(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		OrderDetail(ctx context.Context, in *OrderRequest, out *OrderDetailResponse) error
		Respond(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		Notify(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		Get(ctx context.Context, in *Order, out *OrderResponse) error
		List(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		OrderExtensionList(ctx context.Context, in *OrderRequest, out *OrderExtensionResponse) error
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

func (h *orderServiceHandler) OrderPay(ctx context.Context, in *OrderRequest, out *OrderPayResponse) error {
	return h.OrderServiceHandler.OrderPay(ctx, in, out)
}

func (h *orderServiceHandler) OrderUnderLinePay(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.OrderUnderLinePay(ctx, in, out)
}

func (h *orderServiceHandler) OrderDelivery(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.OrderDelivery(ctx, in, out)
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

func (h *orderServiceHandler) OrderCollect(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.OrderCollect(ctx, in, out)
}

func (h *orderServiceHandler) OrderPayCheck(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.OrderPayCheck(ctx, in, out)
}

func (h *orderServiceHandler) ModifyOrderPrice(ctx context.Context, in *Order, out *OrderResponse) error {
	return h.OrderServiceHandler.ModifyOrderPrice(ctx, in, out)
}

func (h *orderServiceHandler) ModifyOrderAddress(ctx context.Context, in *OrderAddress, out *OrderResponse) error {
	return h.OrderServiceHandler.ModifyOrderAddress(ctx, in, out)
}

func (h *orderServiceHandler) OrderStatusStepTotal(ctx context.Context, in *OrderRequest, out *OrderStatusStepTotalResponse) error {
	return h.OrderServiceHandler.OrderStatusStepTotal(ctx, in, out)
}

func (h *orderServiceHandler) OrderTotal(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.OrderTotal(ctx, in, out)
}

func (h *orderServiceHandler) OrderSearch(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.OrderSearch(ctx, in, out)
}

func (h *orderServiceHandler) OrderDetail(ctx context.Context, in *OrderRequest, out *OrderDetailResponse) error {
	return h.OrderServiceHandler.OrderDetail(ctx, in, out)
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

func (h *orderServiceHandler) OrderExtensionList(ctx context.Context, in *OrderRequest, out *OrderExtensionResponse) error {
	return h.OrderServiceHandler.OrderExtensionList(ctx, in, out)
}
