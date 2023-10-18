// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: currentMemberService.proto

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

// Api Endpoints for CurrentMemberService service

func NewCurrentMemberServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CurrentMemberService service

type CurrentMemberService interface {
	//个人信息中心
	Index(ctx context.Context, in *CurrentMemberRequest, opts ...client.CallOption) (*CurrentMemberIndexResponse, error)
	//客户基本信息
	BaseInfo(ctx context.Context, in *CurrentMemberRequest, opts ...client.CallOption) (*CurrentMemberResponse, error)
}

type currentMemberService struct {
	c    client.Client
	name string
}

func NewCurrentMemberService(name string, c client.Client) CurrentMemberService {
	return &currentMemberService{
		c:    c,
		name: name,
	}
}

func (c *currentMemberService) Index(ctx context.Context, in *CurrentMemberRequest, opts ...client.CallOption) (*CurrentMemberIndexResponse, error) {
	req := c.c.NewRequest(c.name, "CurrentMemberService.Index", in)
	out := new(CurrentMemberIndexResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentMemberService) BaseInfo(ctx context.Context, in *CurrentMemberRequest, opts ...client.CallOption) (*CurrentMemberResponse, error) {
	req := c.c.NewRequest(c.name, "CurrentMemberService.BaseInfo", in)
	out := new(CurrentMemberResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CurrentMemberService service

type CurrentMemberServiceHandler interface {
	//个人信息中心
	Index(context.Context, *CurrentMemberRequest, *CurrentMemberIndexResponse) error
	//客户基本信息
	BaseInfo(context.Context, *CurrentMemberRequest, *CurrentMemberResponse) error
}

func RegisterCurrentMemberServiceHandler(s server.Server, hdlr CurrentMemberServiceHandler, opts ...server.HandlerOption) error {
	type currentMemberService interface {
		Index(ctx context.Context, in *CurrentMemberRequest, out *CurrentMemberIndexResponse) error
		BaseInfo(ctx context.Context, in *CurrentMemberRequest, out *CurrentMemberResponse) error
	}
	type CurrentMemberService struct {
		currentMemberService
	}
	h := &currentMemberServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CurrentMemberService{h}, opts...))
}

type currentMemberServiceHandler struct {
	CurrentMemberServiceHandler
}

func (h *currentMemberServiceHandler) Index(ctx context.Context, in *CurrentMemberRequest, out *CurrentMemberIndexResponse) error {
	return h.CurrentMemberServiceHandler.Index(ctx, in, out)
}

func (h *currentMemberServiceHandler) BaseInfo(ctx context.Context, in *CurrentMemberRequest, out *CurrentMemberResponse) error {
	return h.CurrentMemberServiceHandler.BaseInfo(ctx, in, out)
}
