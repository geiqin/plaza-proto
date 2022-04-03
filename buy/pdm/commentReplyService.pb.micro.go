// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: commentReplyService.proto

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

// Api Endpoints for CommentReplyService service

func NewCommentReplyServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CommentReplyService service

type CommentReplyService interface {
	Create(ctx context.Context, in *CommentReply, opts ...client.CallOption) (*CommentReplyResponse, error)
	Delete(ctx context.Context, in *CommentReply, opts ...client.CallOption) (*CommentReplyResponse, error)
	Get(ctx context.Context, in *CommentReply, opts ...client.CallOption) (*CommentReplyResponse, error)
	List(ctx context.Context, in *CommentReplyRequest, opts ...client.CallOption) (*CommentReplyResponse, error)
	Search(ctx context.Context, in *CommentReplyRequest, opts ...client.CallOption) (*CommentReplyResponse, error)
}

type commentReplyService struct {
	c    client.Client
	name string
}

func NewCommentReplyService(name string, c client.Client) CommentReplyService {
	return &commentReplyService{
		c:    c,
		name: name,
	}
}

func (c *commentReplyService) Create(ctx context.Context, in *CommentReply, opts ...client.CallOption) (*CommentReplyResponse, error) {
	req := c.c.NewRequest(c.name, "CommentReplyService.Create", in)
	out := new(CommentReplyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentReplyService) Delete(ctx context.Context, in *CommentReply, opts ...client.CallOption) (*CommentReplyResponse, error) {
	req := c.c.NewRequest(c.name, "CommentReplyService.Delete", in)
	out := new(CommentReplyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentReplyService) Get(ctx context.Context, in *CommentReply, opts ...client.CallOption) (*CommentReplyResponse, error) {
	req := c.c.NewRequest(c.name, "CommentReplyService.Get", in)
	out := new(CommentReplyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentReplyService) List(ctx context.Context, in *CommentReplyRequest, opts ...client.CallOption) (*CommentReplyResponse, error) {
	req := c.c.NewRequest(c.name, "CommentReplyService.List", in)
	out := new(CommentReplyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentReplyService) Search(ctx context.Context, in *CommentReplyRequest, opts ...client.CallOption) (*CommentReplyResponse, error) {
	req := c.c.NewRequest(c.name, "CommentReplyService.Search", in)
	out := new(CommentReplyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CommentReplyService service

type CommentReplyServiceHandler interface {
	Create(context.Context, *CommentReply, *CommentReplyResponse) error
	Delete(context.Context, *CommentReply, *CommentReplyResponse) error
	Get(context.Context, *CommentReply, *CommentReplyResponse) error
	List(context.Context, *CommentReplyRequest, *CommentReplyResponse) error
	Search(context.Context, *CommentReplyRequest, *CommentReplyResponse) error
}

func RegisterCommentReplyServiceHandler(s server.Server, hdlr CommentReplyServiceHandler, opts ...server.HandlerOption) error {
	type commentReplyService interface {
		Create(ctx context.Context, in *CommentReply, out *CommentReplyResponse) error
		Delete(ctx context.Context, in *CommentReply, out *CommentReplyResponse) error
		Get(ctx context.Context, in *CommentReply, out *CommentReplyResponse) error
		List(ctx context.Context, in *CommentReplyRequest, out *CommentReplyResponse) error
		Search(ctx context.Context, in *CommentReplyRequest, out *CommentReplyResponse) error
	}
	type CommentReplyService struct {
		commentReplyService
	}
	h := &commentReplyServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CommentReplyService{h}, opts...))
}

type commentReplyServiceHandler struct {
	CommentReplyServiceHandler
}

func (h *commentReplyServiceHandler) Create(ctx context.Context, in *CommentReply, out *CommentReplyResponse) error {
	return h.CommentReplyServiceHandler.Create(ctx, in, out)
}

func (h *commentReplyServiceHandler) Delete(ctx context.Context, in *CommentReply, out *CommentReplyResponse) error {
	return h.CommentReplyServiceHandler.Delete(ctx, in, out)
}

func (h *commentReplyServiceHandler) Get(ctx context.Context, in *CommentReply, out *CommentReplyResponse) error {
	return h.CommentReplyServiceHandler.Get(ctx, in, out)
}

func (h *commentReplyServiceHandler) List(ctx context.Context, in *CommentReplyRequest, out *CommentReplyResponse) error {
	return h.CommentReplyServiceHandler.List(ctx, in, out)
}

func (h *commentReplyServiceHandler) Search(ctx context.Context, in *CommentReplyRequest, out *CommentReplyResponse) error {
	return h.CommentReplyServiceHandler.Search(ctx, in, out)
}
