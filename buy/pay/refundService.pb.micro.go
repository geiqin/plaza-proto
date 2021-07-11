// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: refundService.proto

package services

import (
	fmt "fmt"
	common "github.com/geiqin/micro-kit/protobuf/common"
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

// Api Endpoints for RefundService service

func NewRefundServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RefundService service

type RefundService interface {
	//创建退款
	Create(ctx context.Context, in *RefundRequest, opts ...client.CallOption) (*RefundResponse, error)
	//退款对账（向第三方发起对账）
	Reconciliation(ctx context.Context, in *Refund, opts ...client.CallOption) (*RefundResponse, error)
	//删除退款
	Delete(ctx context.Context, in *Refund, opts ...client.CallOption) (*RefundResponse, error)
	//获得退款信息
	Get(ctx context.Context, in *Refund, opts ...client.CallOption) (*RefundResponse, error)
	//查询退款
	Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*RefundResponse, error)
}

type refundService struct {
	c    client.Client
	name string
}

func NewRefundService(name string, c client.Client) RefundService {
	return &refundService{
		c:    c,
		name: name,
	}
}

func (c *refundService) Create(ctx context.Context, in *RefundRequest, opts ...client.CallOption) (*RefundResponse, error) {
	req := c.c.NewRequest(c.name, "RefundService.Create", in)
	out := new(RefundResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refundService) Reconciliation(ctx context.Context, in *Refund, opts ...client.CallOption) (*RefundResponse, error) {
	req := c.c.NewRequest(c.name, "RefundService.Reconciliation", in)
	out := new(RefundResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refundService) Delete(ctx context.Context, in *Refund, opts ...client.CallOption) (*RefundResponse, error) {
	req := c.c.NewRequest(c.name, "RefundService.Delete", in)
	out := new(RefundResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refundService) Get(ctx context.Context, in *Refund, opts ...client.CallOption) (*RefundResponse, error) {
	req := c.c.NewRequest(c.name, "RefundService.Get", in)
	out := new(RefundResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refundService) Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*RefundResponse, error) {
	req := c.c.NewRequest(c.name, "RefundService.Search", in)
	out := new(RefundResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RefundService service

type RefundServiceHandler interface {
	//创建退款
	Create(context.Context, *RefundRequest, *RefundResponse) error
	//退款对账（向第三方发起对账）
	Reconciliation(context.Context, *Refund, *RefundResponse) error
	//删除退款
	Delete(context.Context, *Refund, *RefundResponse) error
	//获得退款信息
	Get(context.Context, *Refund, *RefundResponse) error
	//查询退款
	Search(context.Context, *common.BaseWhere, *RefundResponse) error
}

func RegisterRefundServiceHandler(s server.Server, hdlr RefundServiceHandler, opts ...server.HandlerOption) error {
	type refundService interface {
		Create(ctx context.Context, in *RefundRequest, out *RefundResponse) error
		Reconciliation(ctx context.Context, in *Refund, out *RefundResponse) error
		Delete(ctx context.Context, in *Refund, out *RefundResponse) error
		Get(ctx context.Context, in *Refund, out *RefundResponse) error
		Search(ctx context.Context, in *common.BaseWhere, out *RefundResponse) error
	}
	type RefundService struct {
		refundService
	}
	h := &refundServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RefundService{h}, opts...))
}

type refundServiceHandler struct {
	RefundServiceHandler
}

func (h *refundServiceHandler) Create(ctx context.Context, in *RefundRequest, out *RefundResponse) error {
	return h.RefundServiceHandler.Create(ctx, in, out)
}

func (h *refundServiceHandler) Reconciliation(ctx context.Context, in *Refund, out *RefundResponse) error {
	return h.RefundServiceHandler.Reconciliation(ctx, in, out)
}

func (h *refundServiceHandler) Delete(ctx context.Context, in *Refund, out *RefundResponse) error {
	return h.RefundServiceHandler.Delete(ctx, in, out)
}

func (h *refundServiceHandler) Get(ctx context.Context, in *Refund, out *RefundResponse) error {
	return h.RefundServiceHandler.Get(ctx, in, out)
}

func (h *refundServiceHandler) Search(ctx context.Context, in *common.BaseWhere, out *RefundResponse) error {
	return h.RefundServiceHandler.Search(ctx, in, out)
}
