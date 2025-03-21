// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: messageService.proto

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

// Api Endpoints for MessageService service

func NewMessageServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MessageService service

type MessageService interface {
	//消息增加
	Create(ctx context.Context, in *Message, opts ...client.CallOption) (*MessageResponse, error)
	//消息删除
	Delete(ctx context.Context, in *MessageRequest, opts ...client.CallOption) (*MessageResponse, error)
	//消息已读
	Read(ctx context.Context, in *MessageRequest, opts ...client.CallOption) (*MessageResponse, error)
	//消息获取
	Get(ctx context.Context, in *MessageRequest, opts ...client.CallOption) (*MessageResponse, error)
	//未读消息列表
	Unread(ctx context.Context, in *MessageRequest, opts ...client.CallOption) (*MessageResponse, error)
	//消息查询
	Search(ctx context.Context, in *MessageRequest, opts ...client.CallOption) (*MessageResponse, error)
}

type messageService struct {
	c    client.Client
	name string
}

func NewMessageService(name string, c client.Client) MessageService {
	return &messageService{
		c:    c,
		name: name,
	}
}

func (c *messageService) Create(ctx context.Context, in *Message, opts ...client.CallOption) (*MessageResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.Create", in)
	out := new(MessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) Delete(ctx context.Context, in *MessageRequest, opts ...client.CallOption) (*MessageResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.Delete", in)
	out := new(MessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) Read(ctx context.Context, in *MessageRequest, opts ...client.CallOption) (*MessageResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.Read", in)
	out := new(MessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) Get(ctx context.Context, in *MessageRequest, opts ...client.CallOption) (*MessageResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.Get", in)
	out := new(MessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) Unread(ctx context.Context, in *MessageRequest, opts ...client.CallOption) (*MessageResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.Unread", in)
	out := new(MessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) Search(ctx context.Context, in *MessageRequest, opts ...client.CallOption) (*MessageResponse, error) {
	req := c.c.NewRequest(c.name, "MessageService.Search", in)
	out := new(MessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MessageService service

type MessageServiceHandler interface {
	//消息增加
	Create(context.Context, *Message, *MessageResponse) error
	//消息删除
	Delete(context.Context, *MessageRequest, *MessageResponse) error
	//消息已读
	Read(context.Context, *MessageRequest, *MessageResponse) error
	//消息获取
	Get(context.Context, *MessageRequest, *MessageResponse) error
	//未读消息列表
	Unread(context.Context, *MessageRequest, *MessageResponse) error
	//消息查询
	Search(context.Context, *MessageRequest, *MessageResponse) error
}

func RegisterMessageServiceHandler(s server.Server, hdlr MessageServiceHandler, opts ...server.HandlerOption) error {
	type messageService interface {
		Create(ctx context.Context, in *Message, out *MessageResponse) error
		Delete(ctx context.Context, in *MessageRequest, out *MessageResponse) error
		Read(ctx context.Context, in *MessageRequest, out *MessageResponse) error
		Get(ctx context.Context, in *MessageRequest, out *MessageResponse) error
		Unread(ctx context.Context, in *MessageRequest, out *MessageResponse) error
		Search(ctx context.Context, in *MessageRequest, out *MessageResponse) error
	}
	type MessageService struct {
		messageService
	}
	h := &messageServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MessageService{h}, opts...))
}

type messageServiceHandler struct {
	MessageServiceHandler
}

func (h *messageServiceHandler) Create(ctx context.Context, in *Message, out *MessageResponse) error {
	return h.MessageServiceHandler.Create(ctx, in, out)
}

func (h *messageServiceHandler) Delete(ctx context.Context, in *MessageRequest, out *MessageResponse) error {
	return h.MessageServiceHandler.Delete(ctx, in, out)
}

func (h *messageServiceHandler) Read(ctx context.Context, in *MessageRequest, out *MessageResponse) error {
	return h.MessageServiceHandler.Read(ctx, in, out)
}

func (h *messageServiceHandler) Get(ctx context.Context, in *MessageRequest, out *MessageResponse) error {
	return h.MessageServiceHandler.Get(ctx, in, out)
}

func (h *messageServiceHandler) Unread(ctx context.Context, in *MessageRequest, out *MessageResponse) error {
	return h.MessageServiceHandler.Unread(ctx, in, out)
}

func (h *messageServiceHandler) Search(ctx context.Context, in *MessageRequest, out *MessageResponse) error {
	return h.MessageServiceHandler.Search(ctx, in, out)
}
