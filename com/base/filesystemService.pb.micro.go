// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: filesystemService.proto

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

// Api Endpoints for FilesystemService service

func NewFilesystemServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FilesystemService service

type FilesystemService interface {
	Create(ctx context.Context, in *Filesystem, opts ...client.CallOption) (*FilesystemResponse, error)
	Update(ctx context.Context, in *Filesystem, opts ...client.CallOption) (*FilesystemResponse, error)
	Delete(ctx context.Context, in *Filesystem, opts ...client.CallOption) (*FilesystemResponse, error)
	Get(ctx context.Context, in *Filesystem, opts ...client.CallOption) (*FilesystemResponse, error)
	Search(ctx context.Context, in *FilesystemRequest, opts ...client.CallOption) (*FilesystemResponse, error)
	List(ctx context.Context, in *FilesystemRequest, opts ...client.CallOption) (*FilesystemResponse, error)
}

type filesystemService struct {
	c    client.Client
	name string
}

func NewFilesystemService(name string, c client.Client) FilesystemService {
	return &filesystemService{
		c:    c,
		name: name,
	}
}

func (c *filesystemService) Create(ctx context.Context, in *Filesystem, opts ...client.CallOption) (*FilesystemResponse, error) {
	req := c.c.NewRequest(c.name, "FilesystemService.Create", in)
	out := new(FilesystemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemService) Update(ctx context.Context, in *Filesystem, opts ...client.CallOption) (*FilesystemResponse, error) {
	req := c.c.NewRequest(c.name, "FilesystemService.Update", in)
	out := new(FilesystemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemService) Delete(ctx context.Context, in *Filesystem, opts ...client.CallOption) (*FilesystemResponse, error) {
	req := c.c.NewRequest(c.name, "FilesystemService.Delete", in)
	out := new(FilesystemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemService) Get(ctx context.Context, in *Filesystem, opts ...client.CallOption) (*FilesystemResponse, error) {
	req := c.c.NewRequest(c.name, "FilesystemService.Get", in)
	out := new(FilesystemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemService) Search(ctx context.Context, in *FilesystemRequest, opts ...client.CallOption) (*FilesystemResponse, error) {
	req := c.c.NewRequest(c.name, "FilesystemService.Search", in)
	out := new(FilesystemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemService) List(ctx context.Context, in *FilesystemRequest, opts ...client.CallOption) (*FilesystemResponse, error) {
	req := c.c.NewRequest(c.name, "FilesystemService.List", in)
	out := new(FilesystemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FilesystemService service

type FilesystemServiceHandler interface {
	Create(context.Context, *Filesystem, *FilesystemResponse) error
	Update(context.Context, *Filesystem, *FilesystemResponse) error
	Delete(context.Context, *Filesystem, *FilesystemResponse) error
	Get(context.Context, *Filesystem, *FilesystemResponse) error
	Search(context.Context, *FilesystemRequest, *FilesystemResponse) error
	List(context.Context, *FilesystemRequest, *FilesystemResponse) error
}

func RegisterFilesystemServiceHandler(s server.Server, hdlr FilesystemServiceHandler, opts ...server.HandlerOption) error {
	type filesystemService interface {
		Create(ctx context.Context, in *Filesystem, out *FilesystemResponse) error
		Update(ctx context.Context, in *Filesystem, out *FilesystemResponse) error
		Delete(ctx context.Context, in *Filesystem, out *FilesystemResponse) error
		Get(ctx context.Context, in *Filesystem, out *FilesystemResponse) error
		Search(ctx context.Context, in *FilesystemRequest, out *FilesystemResponse) error
		List(ctx context.Context, in *FilesystemRequest, out *FilesystemResponse) error
	}
	type FilesystemService struct {
		filesystemService
	}
	h := &filesystemServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FilesystemService{h}, opts...))
}

type filesystemServiceHandler struct {
	FilesystemServiceHandler
}

func (h *filesystemServiceHandler) Create(ctx context.Context, in *Filesystem, out *FilesystemResponse) error {
	return h.FilesystemServiceHandler.Create(ctx, in, out)
}

func (h *filesystemServiceHandler) Update(ctx context.Context, in *Filesystem, out *FilesystemResponse) error {
	return h.FilesystemServiceHandler.Update(ctx, in, out)
}

func (h *filesystemServiceHandler) Delete(ctx context.Context, in *Filesystem, out *FilesystemResponse) error {
	return h.FilesystemServiceHandler.Delete(ctx, in, out)
}

func (h *filesystemServiceHandler) Get(ctx context.Context, in *Filesystem, out *FilesystemResponse) error {
	return h.FilesystemServiceHandler.Get(ctx, in, out)
}

func (h *filesystemServiceHandler) Search(ctx context.Context, in *FilesystemRequest, out *FilesystemResponse) error {
	return h.FilesystemServiceHandler.Search(ctx, in, out)
}

func (h *filesystemServiceHandler) List(ctx context.Context, in *FilesystemRequest, out *FilesystemResponse) error {
	return h.FilesystemServiceHandler.List(ctx, in, out)
}