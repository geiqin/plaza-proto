// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: complaintProgressService.proto

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

// Api Endpoints for ComplaintProgressService service

func NewComplaintProgressServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ComplaintProgressService service

type ComplaintProgressService interface {
	// 售后处理进度获取
	Get(ctx context.Context, in *ComplaintProgress, opts ...client.CallOption) (*ComplaintProgressResponse, error)
	// 售后处理进度查询
	Search(ctx context.Context, in *ComplaintProgressRequest, opts ...client.CallOption) (*ComplaintProgressResponse, error)
	// 售后处理进度列表
	List(ctx context.Context, in *ComplaintProgressRequest, opts ...client.CallOption) (*ComplaintProgressResponse, error)
}

type complaintProgressService struct {
	c    client.Client
	name string
}

func NewComplaintProgressService(name string, c client.Client) ComplaintProgressService {
	return &complaintProgressService{
		c:    c,
		name: name,
	}
}

func (c *complaintProgressService) Get(ctx context.Context, in *ComplaintProgress, opts ...client.CallOption) (*ComplaintProgressResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintProgressService.Get", in)
	out := new(ComplaintProgressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintProgressService) Search(ctx context.Context, in *ComplaintProgressRequest, opts ...client.CallOption) (*ComplaintProgressResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintProgressService.Search", in)
	out := new(ComplaintProgressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintProgressService) List(ctx context.Context, in *ComplaintProgressRequest, opts ...client.CallOption) (*ComplaintProgressResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintProgressService.List", in)
	out := new(ComplaintProgressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ComplaintProgressService service

type ComplaintProgressServiceHandler interface {
	// 售后处理进度获取
	Get(context.Context, *ComplaintProgress, *ComplaintProgressResponse) error
	// 售后处理进度查询
	Search(context.Context, *ComplaintProgressRequest, *ComplaintProgressResponse) error
	// 售后处理进度列表
	List(context.Context, *ComplaintProgressRequest, *ComplaintProgressResponse) error
}

func RegisterComplaintProgressServiceHandler(s server.Server, hdlr ComplaintProgressServiceHandler, opts ...server.HandlerOption) error {
	type complaintProgressService interface {
		Get(ctx context.Context, in *ComplaintProgress, out *ComplaintProgressResponse) error
		Search(ctx context.Context, in *ComplaintProgressRequest, out *ComplaintProgressResponse) error
		List(ctx context.Context, in *ComplaintProgressRequest, out *ComplaintProgressResponse) error
	}
	type ComplaintProgressService struct {
		complaintProgressService
	}
	h := &complaintProgressServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ComplaintProgressService{h}, opts...))
}

type complaintProgressServiceHandler struct {
	ComplaintProgressServiceHandler
}

func (h *complaintProgressServiceHandler) Get(ctx context.Context, in *ComplaintProgress, out *ComplaintProgressResponse) error {
	return h.ComplaintProgressServiceHandler.Get(ctx, in, out)
}

func (h *complaintProgressServiceHandler) Search(ctx context.Context, in *ComplaintProgressRequest, out *ComplaintProgressResponse) error {
	return h.ComplaintProgressServiceHandler.Search(ctx, in, out)
}

func (h *complaintProgressServiceHandler) List(ctx context.Context, in *ComplaintProgressRequest, out *ComplaintProgressResponse) error {
	return h.ComplaintProgressServiceHandler.List(ctx, in, out)
}