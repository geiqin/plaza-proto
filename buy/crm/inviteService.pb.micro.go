// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: inviteService.proto

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

// Api Endpoints for InviteService service

func NewInviteServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for InviteService service

type InviteService interface {
	//微信小程序邀请码
	MiniQrcode(ctx context.Context, in *InviteRequest, opts ...client.CallOption) (*InviteResponse, error)
	//APP邀请二维码
	AppQrcode(ctx context.Context, in *InviteRequest, opts ...client.CallOption) (*InviteResponse, error)
	//邀请链接
	Link(ctx context.Context, in *InviteRequest, opts ...client.CallOption) (*InviteResponse, error)
}

type inviteService struct {
	c    client.Client
	name string
}

func NewInviteService(name string, c client.Client) InviteService {
	return &inviteService{
		c:    c,
		name: name,
	}
}

func (c *inviteService) MiniQrcode(ctx context.Context, in *InviteRequest, opts ...client.CallOption) (*InviteResponse, error) {
	req := c.c.NewRequest(c.name, "InviteService.MiniQrcode", in)
	out := new(InviteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteService) AppQrcode(ctx context.Context, in *InviteRequest, opts ...client.CallOption) (*InviteResponse, error) {
	req := c.c.NewRequest(c.name, "InviteService.AppQrcode", in)
	out := new(InviteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inviteService) Link(ctx context.Context, in *InviteRequest, opts ...client.CallOption) (*InviteResponse, error) {
	req := c.c.NewRequest(c.name, "InviteService.Link", in)
	out := new(InviteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InviteService service

type InviteServiceHandler interface {
	//微信小程序邀请码
	MiniQrcode(context.Context, *InviteRequest, *InviteResponse) error
	//APP邀请二维码
	AppQrcode(context.Context, *InviteRequest, *InviteResponse) error
	//邀请链接
	Link(context.Context, *InviteRequest, *InviteResponse) error
}

func RegisterInviteServiceHandler(s server.Server, hdlr InviteServiceHandler, opts ...server.HandlerOption) error {
	type inviteService interface {
		MiniQrcode(ctx context.Context, in *InviteRequest, out *InviteResponse) error
		AppQrcode(ctx context.Context, in *InviteRequest, out *InviteResponse) error
		Link(ctx context.Context, in *InviteRequest, out *InviteResponse) error
	}
	type InviteService struct {
		inviteService
	}
	h := &inviteServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&InviteService{h}, opts...))
}

type inviteServiceHandler struct {
	InviteServiceHandler
}

func (h *inviteServiceHandler) MiniQrcode(ctx context.Context, in *InviteRequest, out *InviteResponse) error {
	return h.InviteServiceHandler.MiniQrcode(ctx, in, out)
}

func (h *inviteServiceHandler) AppQrcode(ctx context.Context, in *InviteRequest, out *InviteResponse) error {
	return h.InviteServiceHandler.AppQrcode(ctx, in, out)
}

func (h *inviteServiceHandler) Link(ctx context.Context, in *InviteRequest, out *InviteResponse) error {
	return h.InviteServiceHandler.Link(ctx, in, out)
}
