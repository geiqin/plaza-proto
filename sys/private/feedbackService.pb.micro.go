// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: feedbackService.proto

package services

import (
	fmt "fmt"
	common "github.com/geiqin/microkit/protobuf/common"
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

// Api Endpoints for FeedbackService service

func NewFeedbackServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FeedbackService service

type FeedbackService interface {
	Get(ctx context.Context, in *Feedback, opts ...client.CallOption) (*FeedbackResponse, error)
	Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*FeedbackResponse, error)
	List(ctx context.Context, in *Feedback, opts ...client.CallOption) (*FeedbackResponse, error)
	Tree(ctx context.Context, in *Feedback, opts ...client.CallOption) (*FeedbackResponse, error)
}

type feedbackService struct {
	c    client.Client
	name string
}

func NewFeedbackService(name string, c client.Client) FeedbackService {
	return &feedbackService{
		c:    c,
		name: name,
	}
}

func (c *feedbackService) Get(ctx context.Context, in *Feedback, opts ...client.CallOption) (*FeedbackResponse, error) {
	req := c.c.NewRequest(c.name, "FeedbackService.Get", in)
	out := new(FeedbackResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackService) Search(ctx context.Context, in *common.BaseWhere, opts ...client.CallOption) (*FeedbackResponse, error) {
	req := c.c.NewRequest(c.name, "FeedbackService.Search", in)
	out := new(FeedbackResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackService) List(ctx context.Context, in *Feedback, opts ...client.CallOption) (*FeedbackResponse, error) {
	req := c.c.NewRequest(c.name, "FeedbackService.List", in)
	out := new(FeedbackResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackService) Tree(ctx context.Context, in *Feedback, opts ...client.CallOption) (*FeedbackResponse, error) {
	req := c.c.NewRequest(c.name, "FeedbackService.Tree", in)
	out := new(FeedbackResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FeedbackService service

type FeedbackServiceHandler interface {
	Get(context.Context, *Feedback, *FeedbackResponse) error
	Search(context.Context, *common.BaseWhere, *FeedbackResponse) error
	List(context.Context, *Feedback, *FeedbackResponse) error
	Tree(context.Context, *Feedback, *FeedbackResponse) error
}

func RegisterFeedbackServiceHandler(s server.Server, hdlr FeedbackServiceHandler, opts ...server.HandlerOption) error {
	type feedbackService interface {
		Get(ctx context.Context, in *Feedback, out *FeedbackResponse) error
		Search(ctx context.Context, in *common.BaseWhere, out *FeedbackResponse) error
		List(ctx context.Context, in *Feedback, out *FeedbackResponse) error
		Tree(ctx context.Context, in *Feedback, out *FeedbackResponse) error
	}
	type FeedbackService struct {
		feedbackService
	}
	h := &feedbackServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FeedbackService{h}, opts...))
}

type feedbackServiceHandler struct {
	FeedbackServiceHandler
}

func (h *feedbackServiceHandler) Get(ctx context.Context, in *Feedback, out *FeedbackResponse) error {
	return h.FeedbackServiceHandler.Get(ctx, in, out)
}

func (h *feedbackServiceHandler) Search(ctx context.Context, in *common.BaseWhere, out *FeedbackResponse) error {
	return h.FeedbackServiceHandler.Search(ctx, in, out)
}

func (h *feedbackServiceHandler) List(ctx context.Context, in *Feedback, out *FeedbackResponse) error {
	return h.FeedbackServiceHandler.List(ctx, in, out)
}

func (h *feedbackServiceHandler) Tree(ctx context.Context, in *Feedback, out *FeedbackResponse) error {
	return h.FeedbackServiceHandler.Tree(ctx, in, out)
}