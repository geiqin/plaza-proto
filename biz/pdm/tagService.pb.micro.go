// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: tagService.proto

package services

import (
	fmt "fmt"
	common "github.com/geiqin/micro-kit/protobuf/common"
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

// Api Endpoints for TagService service

func NewTagServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TagService service

type TagService interface {
	Create(ctx context.Context, in *Tag, opts ...client.CallOption) (*TagResponse, error)
	Update(ctx context.Context, in *Tag, opts ...client.CallOption) (*TagResponse, error)
	Delete(ctx context.Context, in *Tag, opts ...client.CallOption) (*TagResponse, error)
	Get(ctx context.Context, in *Tag, opts ...client.CallOption) (*TagResponse, error)
	List(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*TagResponse, error)
	Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*TagResponse, error)
}

type tagService struct {
	c    client.Client
	name string
}

func NewTagService(name string, c client.Client) TagService {
	return &tagService{
		c:    c,
		name: name,
	}
}

func (c *tagService) Create(ctx context.Context, in *Tag, opts ...client.CallOption) (*TagResponse, error) {
	req := c.c.NewRequest(c.name, "TagService.Create", in)
	out := new(TagResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagService) Update(ctx context.Context, in *Tag, opts ...client.CallOption) (*TagResponse, error) {
	req := c.c.NewRequest(c.name, "TagService.Update", in)
	out := new(TagResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagService) Delete(ctx context.Context, in *Tag, opts ...client.CallOption) (*TagResponse, error) {
	req := c.c.NewRequest(c.name, "TagService.Delete", in)
	out := new(TagResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagService) Get(ctx context.Context, in *Tag, opts ...client.CallOption) (*TagResponse, error) {
	req := c.c.NewRequest(c.name, "TagService.Get", in)
	out := new(TagResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagService) List(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*TagResponse, error) {
	req := c.c.NewRequest(c.name, "TagService.List", in)
	out := new(TagResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagService) Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*TagResponse, error) {
	req := c.c.NewRequest(c.name, "TagService.Search", in)
	out := new(TagResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TagService service

type TagServiceHandler interface {
	Create(context.Context, *Tag, *TagResponse) error
	Update(context.Context, *Tag, *TagResponse) error
	Delete(context.Context, *Tag, *TagResponse) error
	Get(context.Context, *Tag, *TagResponse) error
	List(context.Context, *common.BaseWhere, *TagResponse) error
	Search(context.Context, *common.BaseWhere, *TagResponse) error
}

func RegisterTagServiceHandler(s server.Server, hdlr TagServiceHandler, opts ...server.HandlerOption) error {
	type tagService interface {
		Create(ctx context.Context, in *Tag, out *TagResponse) error
		Update(ctx context.Context, in *Tag, out *TagResponse) error
		Delete(ctx context.Context, in *Tag, out *TagResponse) error
		Get(ctx context.Context, in *Tag, out *TagResponse) error
		List(ctx context.Context, in *common.BaseWhere, out *TagResponse) error
		Search(ctx context.Context, in *common.BaseWhere, out *TagResponse) error
	}
	type TagService struct {
		tagService
	}
	h := &tagServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TagService{h}, opts...))
}

type tagServiceHandler struct {
	TagServiceHandler
}

func (h *tagServiceHandler) Create(ctx context.Context, in *Tag, out *TagResponse) error {
	return h.TagServiceHandler.Create(ctx, in, out)
}

func (h *tagServiceHandler) Update(ctx context.Context, in *Tag, out *TagResponse) error {
	return h.TagServiceHandler.Update(ctx, in, out)
}

func (h *tagServiceHandler) Delete(ctx context.Context, in *Tag, out *TagResponse) error {
	return h.TagServiceHandler.Delete(ctx, in, out)
}

func (h *tagServiceHandler) Get(ctx context.Context, in *Tag, out *TagResponse) error {
	return h.TagServiceHandler.Get(ctx, in, out)
}

func (h *tagServiceHandler) List(ctx context.Context, in *common.BaseWhere, out *TagResponse) error {
	return h.TagServiceHandler.List(ctx, in, out)
}

func (h *tagServiceHandler) Search(ctx context.Context, in *common.BaseWhere, out *TagResponse) error {
	return h.TagServiceHandler.Search(ctx, in, out)
}
