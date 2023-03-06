// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: aftersaleService.proto

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

// Api Endpoints for AftersaleService service

func NewAftersaleServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AftersaleService service

type AftersaleService interface {
	// 售后处理页
	Index(ctx context.Context, in *AftersaleRequest, opts ...client.CallOption) (*AftersaleIndexResponse, error)
	// 申请售后
	Create(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error)
	//审核处理
	AuditHandle(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error)
	//商家已发货（针对换货商品的发货）
	MerchantShipped(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error)
	//商家已收货（已收到回寄的商品）
	MerchantReceived(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error)
	//买家已撤销维权
	BuyerCanceled(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error)
	//买家已寄回商品
	BuyerSent(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error)
	//买家已签收（针对换货商品的签收）
	BuyerReceived(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error)
	//立即打款
	DoRefund(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error)
	// 获取维权信息（基本信息）
	Get(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error)
	//查询维权信息
	Search(ctx context.Context, in *AftersaleRequest, opts ...client.CallOption) (*AftersaleResponse, error)
}

type aftersaleService struct {
	c    client.Client
	name string
}

func NewAftersaleService(name string, c client.Client) AftersaleService {
	return &aftersaleService{
		c:    c,
		name: name,
	}
}

func (c *aftersaleService) Index(ctx context.Context, in *AftersaleRequest, opts ...client.CallOption) (*AftersaleIndexResponse, error) {
	req := c.c.NewRequest(c.name, "AftersaleService.Index", in)
	out := new(AftersaleIndexResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aftersaleService) Create(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error) {
	req := c.c.NewRequest(c.name, "AftersaleService.Create", in)
	out := new(AftersaleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aftersaleService) AuditHandle(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error) {
	req := c.c.NewRequest(c.name, "AftersaleService.AuditHandle", in)
	out := new(AftersaleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aftersaleService) MerchantShipped(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error) {
	req := c.c.NewRequest(c.name, "AftersaleService.MerchantShipped", in)
	out := new(AftersaleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aftersaleService) MerchantReceived(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error) {
	req := c.c.NewRequest(c.name, "AftersaleService.MerchantReceived", in)
	out := new(AftersaleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aftersaleService) BuyerCanceled(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error) {
	req := c.c.NewRequest(c.name, "AftersaleService.BuyerCanceled", in)
	out := new(AftersaleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aftersaleService) BuyerSent(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error) {
	req := c.c.NewRequest(c.name, "AftersaleService.BuyerSent", in)
	out := new(AftersaleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aftersaleService) BuyerReceived(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error) {
	req := c.c.NewRequest(c.name, "AftersaleService.BuyerReceived", in)
	out := new(AftersaleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aftersaleService) DoRefund(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error) {
	req := c.c.NewRequest(c.name, "AftersaleService.DoRefund", in)
	out := new(AftersaleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aftersaleService) Get(ctx context.Context, in *Aftersale, opts ...client.CallOption) (*AftersaleResponse, error) {
	req := c.c.NewRequest(c.name, "AftersaleService.Get", in)
	out := new(AftersaleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aftersaleService) Search(ctx context.Context, in *AftersaleRequest, opts ...client.CallOption) (*AftersaleResponse, error) {
	req := c.c.NewRequest(c.name, "AftersaleService.Search", in)
	out := new(AftersaleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AftersaleService service

type AftersaleServiceHandler interface {
	// 售后处理页
	Index(context.Context, *AftersaleRequest, *AftersaleIndexResponse) error
	// 申请售后
	Create(context.Context, *Aftersale, *AftersaleResponse) error
	//审核处理
	AuditHandle(context.Context, *Aftersale, *AftersaleResponse) error
	//商家已发货（针对换货商品的发货）
	MerchantShipped(context.Context, *Aftersale, *AftersaleResponse) error
	//商家已收货（已收到回寄的商品）
	MerchantReceived(context.Context, *Aftersale, *AftersaleResponse) error
	//买家已撤销维权
	BuyerCanceled(context.Context, *Aftersale, *AftersaleResponse) error
	//买家已寄回商品
	BuyerSent(context.Context, *Aftersale, *AftersaleResponse) error
	//买家已签收（针对换货商品的签收）
	BuyerReceived(context.Context, *Aftersale, *AftersaleResponse) error
	//立即打款
	DoRefund(context.Context, *Aftersale, *AftersaleResponse) error
	// 获取维权信息（基本信息）
	Get(context.Context, *Aftersale, *AftersaleResponse) error
	//查询维权信息
	Search(context.Context, *AftersaleRequest, *AftersaleResponse) error
}

func RegisterAftersaleServiceHandler(s server.Server, hdlr AftersaleServiceHandler, opts ...server.HandlerOption) error {
	type aftersaleService interface {
		Index(ctx context.Context, in *AftersaleRequest, out *AftersaleIndexResponse) error
		Create(ctx context.Context, in *Aftersale, out *AftersaleResponse) error
		AuditHandle(ctx context.Context, in *Aftersale, out *AftersaleResponse) error
		MerchantShipped(ctx context.Context, in *Aftersale, out *AftersaleResponse) error
		MerchantReceived(ctx context.Context, in *Aftersale, out *AftersaleResponse) error
		BuyerCanceled(ctx context.Context, in *Aftersale, out *AftersaleResponse) error
		BuyerSent(ctx context.Context, in *Aftersale, out *AftersaleResponse) error
		BuyerReceived(ctx context.Context, in *Aftersale, out *AftersaleResponse) error
		DoRefund(ctx context.Context, in *Aftersale, out *AftersaleResponse) error
		Get(ctx context.Context, in *Aftersale, out *AftersaleResponse) error
		Search(ctx context.Context, in *AftersaleRequest, out *AftersaleResponse) error
	}
	type AftersaleService struct {
		aftersaleService
	}
	h := &aftersaleServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AftersaleService{h}, opts...))
}

type aftersaleServiceHandler struct {
	AftersaleServiceHandler
}

func (h *aftersaleServiceHandler) Index(ctx context.Context, in *AftersaleRequest, out *AftersaleIndexResponse) error {
	return h.AftersaleServiceHandler.Index(ctx, in, out)
}

func (h *aftersaleServiceHandler) Create(ctx context.Context, in *Aftersale, out *AftersaleResponse) error {
	return h.AftersaleServiceHandler.Create(ctx, in, out)
}

func (h *aftersaleServiceHandler) AuditHandle(ctx context.Context, in *Aftersale, out *AftersaleResponse) error {
	return h.AftersaleServiceHandler.AuditHandle(ctx, in, out)
}

func (h *aftersaleServiceHandler) MerchantShipped(ctx context.Context, in *Aftersale, out *AftersaleResponse) error {
	return h.AftersaleServiceHandler.MerchantShipped(ctx, in, out)
}

func (h *aftersaleServiceHandler) MerchantReceived(ctx context.Context, in *Aftersale, out *AftersaleResponse) error {
	return h.AftersaleServiceHandler.MerchantReceived(ctx, in, out)
}

func (h *aftersaleServiceHandler) BuyerCanceled(ctx context.Context, in *Aftersale, out *AftersaleResponse) error {
	return h.AftersaleServiceHandler.BuyerCanceled(ctx, in, out)
}

func (h *aftersaleServiceHandler) BuyerSent(ctx context.Context, in *Aftersale, out *AftersaleResponse) error {
	return h.AftersaleServiceHandler.BuyerSent(ctx, in, out)
}

func (h *aftersaleServiceHandler) BuyerReceived(ctx context.Context, in *Aftersale, out *AftersaleResponse) error {
	return h.AftersaleServiceHandler.BuyerReceived(ctx, in, out)
}

func (h *aftersaleServiceHandler) DoRefund(ctx context.Context, in *Aftersale, out *AftersaleResponse) error {
	return h.AftersaleServiceHandler.DoRefund(ctx, in, out)
}

func (h *aftersaleServiceHandler) Get(ctx context.Context, in *Aftersale, out *AftersaleResponse) error {
	return h.AftersaleServiceHandler.Get(ctx, in, out)
}

func (h *aftersaleServiceHandler) Search(ctx context.Context, in *AftersaleRequest, out *AftersaleResponse) error {
	return h.AftersaleServiceHandler.Search(ctx, in, out)
}
