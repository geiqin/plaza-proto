// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: attrParamService.proto

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

// Api Endpoints for AttrParamService service

func NewAttrParamServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AttrParamService service

type AttrParamService interface {
	Create(ctx context.Context, in *AttrParam, opts ...client.CallOption) (*AttrParamResponse, error)
	Update(ctx context.Context, in *AttrParam, opts ...client.CallOption) (*AttrParamResponse, error)
	Delete(ctx context.Context, in *AttrParam, opts ...client.CallOption) (*AttrParamResponse, error)
	Get(ctx context.Context, in *AttrParam, opts ...client.CallOption) (*AttrParamResponse, error)
	List(ctx context.Context, in *AttrParamRequest, opts ...client.CallOption) (*AttrParamResponse, error)
	Search(ctx context.Context, in *AttrParamRequest, opts ...client.CallOption) (*AttrParamResponse, error)
	AddValue(ctx context.Context, in *AttrParamRequest, opts ...client.CallOption) (*AttrParamResponse, error)
}

type attrParamService struct {
	c    client.Client
	name string
}

func NewAttrParamService(name string, c client.Client) AttrParamService {
	return &attrParamService{
		c:    c,
		name: name,
	}
}

func (c *attrParamService) Create(ctx context.Context, in *AttrParam, opts ...client.CallOption) (*AttrParamResponse, error) {
	req := c.c.NewRequest(c.name, "AttrParamService.Create", in)
	out := new(AttrParamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrParamService) Update(ctx context.Context, in *AttrParam, opts ...client.CallOption) (*AttrParamResponse, error) {
	req := c.c.NewRequest(c.name, "AttrParamService.Update", in)
	out := new(AttrParamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrParamService) Delete(ctx context.Context, in *AttrParam, opts ...client.CallOption) (*AttrParamResponse, error) {
	req := c.c.NewRequest(c.name, "AttrParamService.Delete", in)
	out := new(AttrParamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrParamService) Get(ctx context.Context, in *AttrParam, opts ...client.CallOption) (*AttrParamResponse, error) {
	req := c.c.NewRequest(c.name, "AttrParamService.Get", in)
	out := new(AttrParamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrParamService) List(ctx context.Context, in *AttrParamRequest, opts ...client.CallOption) (*AttrParamResponse, error) {
	req := c.c.NewRequest(c.name, "AttrParamService.List", in)
	out := new(AttrParamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrParamService) Search(ctx context.Context, in *AttrParamRequest, opts ...client.CallOption) (*AttrParamResponse, error) {
	req := c.c.NewRequest(c.name, "AttrParamService.Search", in)
	out := new(AttrParamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrParamService) AddValue(ctx context.Context, in *AttrParamRequest, opts ...client.CallOption) (*AttrParamResponse, error) {
	req := c.c.NewRequest(c.name, "AttrParamService.AddValue", in)
	out := new(AttrParamResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AttrParamService service

type AttrParamServiceHandler interface {
	Create(context.Context, *AttrParam, *AttrParamResponse) error
	Update(context.Context, *AttrParam, *AttrParamResponse) error
	Delete(context.Context, *AttrParam, *AttrParamResponse) error
	Get(context.Context, *AttrParam, *AttrParamResponse) error
	List(context.Context, *AttrParamRequest, *AttrParamResponse) error
	Search(context.Context, *AttrParamRequest, *AttrParamResponse) error
	AddValue(context.Context, *AttrParamRequest, *AttrParamResponse) error
}

func RegisterAttrParamServiceHandler(s server.Server, hdlr AttrParamServiceHandler, opts ...server.HandlerOption) error {
	type attrParamService interface {
		Create(ctx context.Context, in *AttrParam, out *AttrParamResponse) error
		Update(ctx context.Context, in *AttrParam, out *AttrParamResponse) error
		Delete(ctx context.Context, in *AttrParam, out *AttrParamResponse) error
		Get(ctx context.Context, in *AttrParam, out *AttrParamResponse) error
		List(ctx context.Context, in *AttrParamRequest, out *AttrParamResponse) error
		Search(ctx context.Context, in *AttrParamRequest, out *AttrParamResponse) error
		AddValue(ctx context.Context, in *AttrParamRequest, out *AttrParamResponse) error
	}
	type AttrParamService struct {
		attrParamService
	}
	h := &attrParamServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AttrParamService{h}, opts...))
}

type attrParamServiceHandler struct {
	AttrParamServiceHandler
}

func (h *attrParamServiceHandler) Create(ctx context.Context, in *AttrParam, out *AttrParamResponse) error {
	return h.AttrParamServiceHandler.Create(ctx, in, out)
}

func (h *attrParamServiceHandler) Update(ctx context.Context, in *AttrParam, out *AttrParamResponse) error {
	return h.AttrParamServiceHandler.Update(ctx, in, out)
}

func (h *attrParamServiceHandler) Delete(ctx context.Context, in *AttrParam, out *AttrParamResponse) error {
	return h.AttrParamServiceHandler.Delete(ctx, in, out)
}

func (h *attrParamServiceHandler) Get(ctx context.Context, in *AttrParam, out *AttrParamResponse) error {
	return h.AttrParamServiceHandler.Get(ctx, in, out)
}

func (h *attrParamServiceHandler) List(ctx context.Context, in *AttrParamRequest, out *AttrParamResponse) error {
	return h.AttrParamServiceHandler.List(ctx, in, out)
}

func (h *attrParamServiceHandler) Search(ctx context.Context, in *AttrParamRequest, out *AttrParamResponse) error {
	return h.AttrParamServiceHandler.Search(ctx, in, out)
}

func (h *attrParamServiceHandler) AddValue(ctx context.Context, in *AttrParamRequest, out *AttrParamResponse) error {
	return h.AttrParamServiceHandler.AddValue(ctx, in, out)
}
