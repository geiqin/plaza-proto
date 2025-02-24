// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: mentionPointService.proto

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

// Api Endpoints for MentionPointService service

func NewMentionPointServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MentionPointService service

type MentionPointService interface {
	Create(ctx context.Context, in *MentionPoint, opts ...client.CallOption) (*MentionPointResponse, error)
	Update(ctx context.Context, in *MentionPoint, opts ...client.CallOption) (*MentionPointResponse, error)
	Delete(ctx context.Context, in *MentionPointRequest, opts ...client.CallOption) (*MentionPointResponse, error)
	Get(ctx context.Context, in *MentionPoint, opts ...client.CallOption) (*MentionPointResponse, error)
	Search(ctx context.Context, in *MentionPointRequest, opts ...client.CallOption) (*MentionPointResponse, error)
	List(ctx context.Context, in *MentionPointRequest, opts ...client.CallOption) (*MentionPointResponse, error)
	// 获取距离用户最近的自提点
	NearestList(ctx context.Context, in *MentionPointRequest, opts ...client.CallOption) (*MentionPointResponse, error)
}

type mentionPointService struct {
	c    client.Client
	name string
}

func NewMentionPointService(name string, c client.Client) MentionPointService {
	return &mentionPointService{
		c:    c,
		name: name,
	}
}

func (c *mentionPointService) Create(ctx context.Context, in *MentionPoint, opts ...client.CallOption) (*MentionPointResponse, error) {
	req := c.c.NewRequest(c.name, "MentionPointService.Create", in)
	out := new(MentionPointResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mentionPointService) Update(ctx context.Context, in *MentionPoint, opts ...client.CallOption) (*MentionPointResponse, error) {
	req := c.c.NewRequest(c.name, "MentionPointService.Update", in)
	out := new(MentionPointResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mentionPointService) Delete(ctx context.Context, in *MentionPointRequest, opts ...client.CallOption) (*MentionPointResponse, error) {
	req := c.c.NewRequest(c.name, "MentionPointService.Delete", in)
	out := new(MentionPointResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mentionPointService) Get(ctx context.Context, in *MentionPoint, opts ...client.CallOption) (*MentionPointResponse, error) {
	req := c.c.NewRequest(c.name, "MentionPointService.Get", in)
	out := new(MentionPointResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mentionPointService) Search(ctx context.Context, in *MentionPointRequest, opts ...client.CallOption) (*MentionPointResponse, error) {
	req := c.c.NewRequest(c.name, "MentionPointService.Search", in)
	out := new(MentionPointResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mentionPointService) List(ctx context.Context, in *MentionPointRequest, opts ...client.CallOption) (*MentionPointResponse, error) {
	req := c.c.NewRequest(c.name, "MentionPointService.List", in)
	out := new(MentionPointResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mentionPointService) NearestList(ctx context.Context, in *MentionPointRequest, opts ...client.CallOption) (*MentionPointResponse, error) {
	req := c.c.NewRequest(c.name, "MentionPointService.NearestList", in)
	out := new(MentionPointResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MentionPointService service

type MentionPointServiceHandler interface {
	Create(context.Context, *MentionPoint, *MentionPointResponse) error
	Update(context.Context, *MentionPoint, *MentionPointResponse) error
	Delete(context.Context, *MentionPointRequest, *MentionPointResponse) error
	Get(context.Context, *MentionPoint, *MentionPointResponse) error
	Search(context.Context, *MentionPointRequest, *MentionPointResponse) error
	List(context.Context, *MentionPointRequest, *MentionPointResponse) error
	// 获取距离用户最近的自提点
	NearestList(context.Context, *MentionPointRequest, *MentionPointResponse) error
}

func RegisterMentionPointServiceHandler(s server.Server, hdlr MentionPointServiceHandler, opts ...server.HandlerOption) error {
	type mentionPointService interface {
		Create(ctx context.Context, in *MentionPoint, out *MentionPointResponse) error
		Update(ctx context.Context, in *MentionPoint, out *MentionPointResponse) error
		Delete(ctx context.Context, in *MentionPointRequest, out *MentionPointResponse) error
		Get(ctx context.Context, in *MentionPoint, out *MentionPointResponse) error
		Search(ctx context.Context, in *MentionPointRequest, out *MentionPointResponse) error
		List(ctx context.Context, in *MentionPointRequest, out *MentionPointResponse) error
		NearestList(ctx context.Context, in *MentionPointRequest, out *MentionPointResponse) error
	}
	type MentionPointService struct {
		mentionPointService
	}
	h := &mentionPointServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MentionPointService{h}, opts...))
}

type mentionPointServiceHandler struct {
	MentionPointServiceHandler
}

func (h *mentionPointServiceHandler) Create(ctx context.Context, in *MentionPoint, out *MentionPointResponse) error {
	return h.MentionPointServiceHandler.Create(ctx, in, out)
}

func (h *mentionPointServiceHandler) Update(ctx context.Context, in *MentionPoint, out *MentionPointResponse) error {
	return h.MentionPointServiceHandler.Update(ctx, in, out)
}

func (h *mentionPointServiceHandler) Delete(ctx context.Context, in *MentionPointRequest, out *MentionPointResponse) error {
	return h.MentionPointServiceHandler.Delete(ctx, in, out)
}

func (h *mentionPointServiceHandler) Get(ctx context.Context, in *MentionPoint, out *MentionPointResponse) error {
	return h.MentionPointServiceHandler.Get(ctx, in, out)
}

func (h *mentionPointServiceHandler) Search(ctx context.Context, in *MentionPointRequest, out *MentionPointResponse) error {
	return h.MentionPointServiceHandler.Search(ctx, in, out)
}

func (h *mentionPointServiceHandler) List(ctx context.Context, in *MentionPointRequest, out *MentionPointResponse) error {
	return h.MentionPointServiceHandler.List(ctx, in, out)
}

func (h *mentionPointServiceHandler) NearestList(ctx context.Context, in *MentionPointRequest, out *MentionPointResponse) error {
	return h.MentionPointServiceHandler.NearestList(ctx, in, out)
}
