// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: memberIdentityService.proto

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

// Api Endpoints for MemberIdentityService service

func NewMemberIdentityServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MemberIdentityService service

type MemberIdentityService interface {
	Create(ctx context.Context, in *MemberIdentity, opts ...client.CallOption) (*MemberIdentityResponse, error)
	Get(ctx context.Context, in *MemberIdentity, opts ...client.CallOption) (*MemberIdentityResponse, error)
	Search(ctx context.Context, in *MemberIdentityRequest, opts ...client.CallOption) (*MemberIdentityResponse, error)
}

type memberIdentityService struct {
	c    client.Client
	name string
}

func NewMemberIdentityService(name string, c client.Client) MemberIdentityService {
	return &memberIdentityService{
		c:    c,
		name: name,
	}
}

func (c *memberIdentityService) Create(ctx context.Context, in *MemberIdentity, opts ...client.CallOption) (*MemberIdentityResponse, error) {
	req := c.c.NewRequest(c.name, "MemberIdentityService.Create", in)
	out := new(MemberIdentityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberIdentityService) Get(ctx context.Context, in *MemberIdentity, opts ...client.CallOption) (*MemberIdentityResponse, error) {
	req := c.c.NewRequest(c.name, "MemberIdentityService.Get", in)
	out := new(MemberIdentityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberIdentityService) Search(ctx context.Context, in *MemberIdentityRequest, opts ...client.CallOption) (*MemberIdentityResponse, error) {
	req := c.c.NewRequest(c.name, "MemberIdentityService.Search", in)
	out := new(MemberIdentityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MemberIdentityService service

type MemberIdentityServiceHandler interface {
	Create(context.Context, *MemberIdentity, *MemberIdentityResponse) error
	Get(context.Context, *MemberIdentity, *MemberIdentityResponse) error
	Search(context.Context, *MemberIdentityRequest, *MemberIdentityResponse) error
}

func RegisterMemberIdentityServiceHandler(s server.Server, hdlr MemberIdentityServiceHandler, opts ...server.HandlerOption) error {
	type memberIdentityService interface {
		Create(ctx context.Context, in *MemberIdentity, out *MemberIdentityResponse) error
		Get(ctx context.Context, in *MemberIdentity, out *MemberIdentityResponse) error
		Search(ctx context.Context, in *MemberIdentityRequest, out *MemberIdentityResponse) error
	}
	type MemberIdentityService struct {
		memberIdentityService
	}
	h := &memberIdentityServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MemberIdentityService{h}, opts...))
}

type memberIdentityServiceHandler struct {
	MemberIdentityServiceHandler
}

func (h *memberIdentityServiceHandler) Create(ctx context.Context, in *MemberIdentity, out *MemberIdentityResponse) error {
	return h.MemberIdentityServiceHandler.Create(ctx, in, out)
}

func (h *memberIdentityServiceHandler) Get(ctx context.Context, in *MemberIdentity, out *MemberIdentityResponse) error {
	return h.MemberIdentityServiceHandler.Get(ctx, in, out)
}

func (h *memberIdentityServiceHandler) Search(ctx context.Context, in *MemberIdentityRequest, out *MemberIdentityResponse) error {
	return h.MemberIdentityServiceHandler.Search(ctx, in, out)
}
