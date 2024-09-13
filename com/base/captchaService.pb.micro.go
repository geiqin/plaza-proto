// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: captchaService.proto

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

// Api Endpoints for CaptchaService service

func NewCaptchaServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CaptchaService service

type CaptchaService interface {
	// 验证
	Check(ctx context.Context, in *CaptchaRequest, opts ...client.CallOption) (*CaptchaResponse, error)
	// 创建
	Create(ctx context.Context, in *CaptchaRequest, opts ...client.CallOption) (*CaptchaResponse, error)
}

type captchaService struct {
	c    client.Client
	name string
}

func NewCaptchaService(name string, c client.Client) CaptchaService {
	return &captchaService{
		c:    c,
		name: name,
	}
}

func (c *captchaService) Check(ctx context.Context, in *CaptchaRequest, opts ...client.CallOption) (*CaptchaResponse, error) {
	req := c.c.NewRequest(c.name, "CaptchaService.Check", in)
	out := new(CaptchaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captchaService) Create(ctx context.Context, in *CaptchaRequest, opts ...client.CallOption) (*CaptchaResponse, error) {
	req := c.c.NewRequest(c.name, "CaptchaService.Create", in)
	out := new(CaptchaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CaptchaService service

type CaptchaServiceHandler interface {
	// 验证
	Check(context.Context, *CaptchaRequest, *CaptchaResponse) error
	// 创建
	Create(context.Context, *CaptchaRequest, *CaptchaResponse) error
}

func RegisterCaptchaServiceHandler(s server.Server, hdlr CaptchaServiceHandler, opts ...server.HandlerOption) error {
	type captchaService interface {
		Check(ctx context.Context, in *CaptchaRequest, out *CaptchaResponse) error
		Create(ctx context.Context, in *CaptchaRequest, out *CaptchaResponse) error
	}
	type CaptchaService struct {
		captchaService
	}
	h := &captchaServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CaptchaService{h}, opts...))
}

type captchaServiceHandler struct {
	CaptchaServiceHandler
}

func (h *captchaServiceHandler) Check(ctx context.Context, in *CaptchaRequest, out *CaptchaResponse) error {
	return h.CaptchaServiceHandler.Check(ctx, in, out)
}

func (h *captchaServiceHandler) Create(ctx context.Context, in *CaptchaRequest, out *CaptchaResponse) error {
	return h.CaptchaServiceHandler.Create(ctx, in, out)
}
