// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: tunnelPayService.proto

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

// Api Endpoints for TunnelPayService service

func NewTunnelPayServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TunnelPayService service

type TunnelPayService interface {
	Get(ctx context.Context, in *TunnelPay, opts ...client.CallOption) (*TunnelPayResponse, error)
	GetByName(ctx context.Context, in *TunnelPay, opts ...client.CallOption) (*TunnelPayResponse, error)
	Search(ctx context.Context, in *TunnelPayRequest, opts ...client.CallOption) (*TunnelPayResponse, error)
}

type tunnelPayService struct {
	c    client.Client
	name string
}

func NewTunnelPayService(name string, c client.Client) TunnelPayService {
	return &tunnelPayService{
		c:    c,
		name: name,
	}
}

func (c *tunnelPayService) Get(ctx context.Context, in *TunnelPay, opts ...client.CallOption) (*TunnelPayResponse, error) {
	req := c.c.NewRequest(c.name, "TunnelPayService.Get", in)
	out := new(TunnelPayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelPayService) GetByName(ctx context.Context, in *TunnelPay, opts ...client.CallOption) (*TunnelPayResponse, error) {
	req := c.c.NewRequest(c.name, "TunnelPayService.GetByName", in)
	out := new(TunnelPayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tunnelPayService) Search(ctx context.Context, in *TunnelPayRequest, opts ...client.CallOption) (*TunnelPayResponse, error) {
	req := c.c.NewRequest(c.name, "TunnelPayService.Search", in)
	out := new(TunnelPayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TunnelPayService service

type TunnelPayServiceHandler interface {
	Get(context.Context, *TunnelPay, *TunnelPayResponse) error
	GetByName(context.Context, *TunnelPay, *TunnelPayResponse) error
	Search(context.Context, *TunnelPayRequest, *TunnelPayResponse) error
}

func RegisterTunnelPayServiceHandler(s server.Server, hdlr TunnelPayServiceHandler, opts ...server.HandlerOption) error {
	type tunnelPayService interface {
		Get(ctx context.Context, in *TunnelPay, out *TunnelPayResponse) error
		GetByName(ctx context.Context, in *TunnelPay, out *TunnelPayResponse) error
		Search(ctx context.Context, in *TunnelPayRequest, out *TunnelPayResponse) error
	}
	type TunnelPayService struct {
		tunnelPayService
	}
	h := &tunnelPayServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TunnelPayService{h}, opts...))
}

type tunnelPayServiceHandler struct {
	TunnelPayServiceHandler
}

func (h *tunnelPayServiceHandler) Get(ctx context.Context, in *TunnelPay, out *TunnelPayResponse) error {
	return h.TunnelPayServiceHandler.Get(ctx, in, out)
}

func (h *tunnelPayServiceHandler) GetByName(ctx context.Context, in *TunnelPay, out *TunnelPayResponse) error {
	return h.TunnelPayServiceHandler.GetByName(ctx, in, out)
}

func (h *tunnelPayServiceHandler) Search(ctx context.Context, in *TunnelPayRequest, out *TunnelPayResponse) error {
	return h.TunnelPayServiceHandler.Search(ctx, in, out)
}
