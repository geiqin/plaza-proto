// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: attachmentService.proto

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

// Api Endpoints for AttachmentService service

func NewAttachmentServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AttachmentService service

type AttachmentService interface {
	// 附件新增
	Create(ctx context.Context, in *Attachment, opts ...client.CallOption) (*AttachmentResponse, error)
	// 附件修改
	Update(ctx context.Context, in *Attachment, opts ...client.CallOption) (*AttachmentResponse, error)
	// 附件删除
	Delete(ctx context.Context, in *Attachment, opts ...client.CallOption) (*AttachmentResponse, error)
	// 附件获取
	Get(ctx context.Context, in *Attachment, opts ...client.CallOption) (*AttachmentResponse, error)
	// 附件查询
	Search(ctx context.Context, in *AttachmentRequest, opts ...client.CallOption) (*AttachmentResponse, error)
	// 附件列表
	List(ctx context.Context, in *AttachmentRequest, opts ...client.CallOption) (*AttachmentResponse, error)
	//上传凭证生成
	UploadToken(ctx context.Context, in *AttachmentRequest, opts ...client.CallOption) (*AttachmentResponse, error)
	//上传回调处理
	UploadCallback(ctx context.Context, in *AttachmentRequest, opts ...client.CallOption) (*AttachmentResponse, error)
}

type attachmentService struct {
	c    client.Client
	name string
}

func NewAttachmentService(name string, c client.Client) AttachmentService {
	return &attachmentService{
		c:    c,
		name: name,
	}
}

func (c *attachmentService) Create(ctx context.Context, in *Attachment, opts ...client.CallOption) (*AttachmentResponse, error) {
	req := c.c.NewRequest(c.name, "AttachmentService.Create", in)
	out := new(AttachmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentService) Update(ctx context.Context, in *Attachment, opts ...client.CallOption) (*AttachmentResponse, error) {
	req := c.c.NewRequest(c.name, "AttachmentService.Update", in)
	out := new(AttachmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentService) Delete(ctx context.Context, in *Attachment, opts ...client.CallOption) (*AttachmentResponse, error) {
	req := c.c.NewRequest(c.name, "AttachmentService.Delete", in)
	out := new(AttachmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentService) Get(ctx context.Context, in *Attachment, opts ...client.CallOption) (*AttachmentResponse, error) {
	req := c.c.NewRequest(c.name, "AttachmentService.Get", in)
	out := new(AttachmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentService) Search(ctx context.Context, in *AttachmentRequest, opts ...client.CallOption) (*AttachmentResponse, error) {
	req := c.c.NewRequest(c.name, "AttachmentService.Search", in)
	out := new(AttachmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentService) List(ctx context.Context, in *AttachmentRequest, opts ...client.CallOption) (*AttachmentResponse, error) {
	req := c.c.NewRequest(c.name, "AttachmentService.List", in)
	out := new(AttachmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentService) UploadToken(ctx context.Context, in *AttachmentRequest, opts ...client.CallOption) (*AttachmentResponse, error) {
	req := c.c.NewRequest(c.name, "AttachmentService.UploadToken", in)
	out := new(AttachmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentService) UploadCallback(ctx context.Context, in *AttachmentRequest, opts ...client.CallOption) (*AttachmentResponse, error) {
	req := c.c.NewRequest(c.name, "AttachmentService.UploadCallback", in)
	out := new(AttachmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AttachmentService service

type AttachmentServiceHandler interface {
	// 附件新增
	Create(context.Context, *Attachment, *AttachmentResponse) error
	// 附件修改
	Update(context.Context, *Attachment, *AttachmentResponse) error
	// 附件删除
	Delete(context.Context, *Attachment, *AttachmentResponse) error
	// 附件获取
	Get(context.Context, *Attachment, *AttachmentResponse) error
	// 附件查询
	Search(context.Context, *AttachmentRequest, *AttachmentResponse) error
	// 附件列表
	List(context.Context, *AttachmentRequest, *AttachmentResponse) error
	//上传凭证生成
	UploadToken(context.Context, *AttachmentRequest, *AttachmentResponse) error
	//上传回调处理
	UploadCallback(context.Context, *AttachmentRequest, *AttachmentResponse) error
}

func RegisterAttachmentServiceHandler(s server.Server, hdlr AttachmentServiceHandler, opts ...server.HandlerOption) error {
	type attachmentService interface {
		Create(ctx context.Context, in *Attachment, out *AttachmentResponse) error
		Update(ctx context.Context, in *Attachment, out *AttachmentResponse) error
		Delete(ctx context.Context, in *Attachment, out *AttachmentResponse) error
		Get(ctx context.Context, in *Attachment, out *AttachmentResponse) error
		Search(ctx context.Context, in *AttachmentRequest, out *AttachmentResponse) error
		List(ctx context.Context, in *AttachmentRequest, out *AttachmentResponse) error
		UploadToken(ctx context.Context, in *AttachmentRequest, out *AttachmentResponse) error
		UploadCallback(ctx context.Context, in *AttachmentRequest, out *AttachmentResponse) error
	}
	type AttachmentService struct {
		attachmentService
	}
	h := &attachmentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AttachmentService{h}, opts...))
}

type attachmentServiceHandler struct {
	AttachmentServiceHandler
}

func (h *attachmentServiceHandler) Create(ctx context.Context, in *Attachment, out *AttachmentResponse) error {
	return h.AttachmentServiceHandler.Create(ctx, in, out)
}

func (h *attachmentServiceHandler) Update(ctx context.Context, in *Attachment, out *AttachmentResponse) error {
	return h.AttachmentServiceHandler.Update(ctx, in, out)
}

func (h *attachmentServiceHandler) Delete(ctx context.Context, in *Attachment, out *AttachmentResponse) error {
	return h.AttachmentServiceHandler.Delete(ctx, in, out)
}

func (h *attachmentServiceHandler) Get(ctx context.Context, in *Attachment, out *AttachmentResponse) error {
	return h.AttachmentServiceHandler.Get(ctx, in, out)
}

func (h *attachmentServiceHandler) Search(ctx context.Context, in *AttachmentRequest, out *AttachmentResponse) error {
	return h.AttachmentServiceHandler.Search(ctx, in, out)
}

func (h *attachmentServiceHandler) List(ctx context.Context, in *AttachmentRequest, out *AttachmentResponse) error {
	return h.AttachmentServiceHandler.List(ctx, in, out)
}

func (h *attachmentServiceHandler) UploadToken(ctx context.Context, in *AttachmentRequest, out *AttachmentResponse) error {
	return h.AttachmentServiceHandler.UploadToken(ctx, in, out)
}

func (h *attachmentServiceHandler) UploadCallback(ctx context.Context, in *AttachmentRequest, out *AttachmentResponse) error {
	return h.AttachmentServiceHandler.UploadCallback(ctx, in, out)
}