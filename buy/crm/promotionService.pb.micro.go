// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: promotionService.proto

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

// Api Endpoints for PromotionService service

func NewPromotionServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PromotionService service

type PromotionService interface {
	//微信小程序邀请码
	MiniQrcode(ctx context.Context, in *PromotionRequest, opts ...client.CallOption) (*PromotionResponse, error)
	//APP邀请二维码
	AppQrcode(ctx context.Context, in *PromotionRequest, opts ...client.CallOption) (*PromotionResponse, error)
	//邀请链接
	Link(ctx context.Context, in *PromotionRequest, opts ...client.CallOption) (*PromotionResponse, error)
}

type promotionService struct {
	c    client.Client
	name string
}

func NewPromotionService(name string, c client.Client) PromotionService {
	return &promotionService{
		c:    c,
		name: name,
	}
}

func (c *promotionService) MiniQrcode(ctx context.Context, in *PromotionRequest, opts ...client.CallOption) (*PromotionResponse, error) {
	req := c.c.NewRequest(c.name, "PromotionService.MiniQrcode", in)
	out := new(PromotionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionService) AppQrcode(ctx context.Context, in *PromotionRequest, opts ...client.CallOption) (*PromotionResponse, error) {
	req := c.c.NewRequest(c.name, "PromotionService.AppQrcode", in)
	out := new(PromotionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionService) Link(ctx context.Context, in *PromotionRequest, opts ...client.CallOption) (*PromotionResponse, error) {
	req := c.c.NewRequest(c.name, "PromotionService.Link", in)
	out := new(PromotionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PromotionService service

type PromotionServiceHandler interface {
	//微信小程序邀请码
	MiniQrcode(context.Context, *PromotionRequest, *PromotionResponse) error
	//APP邀请二维码
	AppQrcode(context.Context, *PromotionRequest, *PromotionResponse) error
	//邀请链接
	Link(context.Context, *PromotionRequest, *PromotionResponse) error
}

func RegisterPromotionServiceHandler(s server.Server, hdlr PromotionServiceHandler, opts ...server.HandlerOption) error {
	type promotionService interface {
		MiniQrcode(ctx context.Context, in *PromotionRequest, out *PromotionResponse) error
		AppQrcode(ctx context.Context, in *PromotionRequest, out *PromotionResponse) error
		Link(ctx context.Context, in *PromotionRequest, out *PromotionResponse) error
	}
	type PromotionService struct {
		promotionService
	}
	h := &promotionServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PromotionService{h}, opts...))
}

type promotionServiceHandler struct {
	PromotionServiceHandler
}

func (h *promotionServiceHandler) MiniQrcode(ctx context.Context, in *PromotionRequest, out *PromotionResponse) error {
	return h.PromotionServiceHandler.MiniQrcode(ctx, in, out)
}

func (h *promotionServiceHandler) AppQrcode(ctx context.Context, in *PromotionRequest, out *PromotionResponse) error {
	return h.PromotionServiceHandler.AppQrcode(ctx, in, out)
}

func (h *promotionServiceHandler) Link(ctx context.Context, in *PromotionRequest, out *PromotionResponse) error {
	return h.PromotionServiceHandler.Link(ctx, in, out)
}
