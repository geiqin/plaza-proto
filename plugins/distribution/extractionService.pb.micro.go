// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: extractionService.proto

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

// Api Endpoints for ExtractionService service

func NewExtractionServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ExtractionService service

type ExtractionService interface {
	//用户取货点首页
	ExtractionIndex(ctx context.Context, in *Extraction, opts ...client.CallOption) (*ExtractionIndexResponse, error)
	//用户取货点申请信息
	ExtractionApplyInfo(ctx context.Context, in *Extraction, opts ...client.CallOption) (*ExtractionResponse, error)
	//用户取货点申请保存
	ExtractionApplySave(ctx context.Context, in *Extraction, opts ...client.CallOption) (*ExtractionResponse, error)
	//获取用户取货点数据
	ExtractionDetail(ctx context.Context, in *Extraction, opts ...client.CallOption) (*ExtractionResponse, error)
	//取货点审核
	ExtractionAudit(ctx context.Context, in *Extraction, opts ...client.CallOption) (*ExtractionResponse, error)
	//取货点解约
	ExtractionRelieve(ctx context.Context, in *ExtractionRequest, opts ...client.CallOption) (*ExtractionResponse, error)
	//取货点删除
	ExtractionDelete(ctx context.Context, in *ExtractionRequest, opts ...client.CallOption) (*ExtractionResponse, error)
	//取货点切换保存
	ExtractionSwitchSave(ctx context.Context, in *ExtractionRequest, opts ...client.CallOption) (*ExtractionResponse, error)
	//获取用户自定义地址
	UserCustomExtractionAddress(ctx context.Context, in *ExtractionRequest, opts ...client.CallOption) (*ExtractionResponse, error)
	//取货点查询
	Search(ctx context.Context, in *ExtractionRequest, opts ...client.CallOption) (*ExtractionResponse, error)
}

type extractionService struct {
	c    client.Client
	name string
}

func NewExtractionService(name string, c client.Client) ExtractionService {
	return &extractionService{
		c:    c,
		name: name,
	}
}

func (c *extractionService) ExtractionIndex(ctx context.Context, in *Extraction, opts ...client.CallOption) (*ExtractionIndexResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionService.ExtractionIndex", in)
	out := new(ExtractionIndexResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extractionService) ExtractionApplyInfo(ctx context.Context, in *Extraction, opts ...client.CallOption) (*ExtractionResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionService.ExtractionApplyInfo", in)
	out := new(ExtractionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extractionService) ExtractionApplySave(ctx context.Context, in *Extraction, opts ...client.CallOption) (*ExtractionResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionService.ExtractionApplySave", in)
	out := new(ExtractionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extractionService) ExtractionDetail(ctx context.Context, in *Extraction, opts ...client.CallOption) (*ExtractionResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionService.ExtractionDetail", in)
	out := new(ExtractionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extractionService) ExtractionAudit(ctx context.Context, in *Extraction, opts ...client.CallOption) (*ExtractionResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionService.ExtractionAudit", in)
	out := new(ExtractionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extractionService) ExtractionRelieve(ctx context.Context, in *ExtractionRequest, opts ...client.CallOption) (*ExtractionResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionService.ExtractionRelieve", in)
	out := new(ExtractionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extractionService) ExtractionDelete(ctx context.Context, in *ExtractionRequest, opts ...client.CallOption) (*ExtractionResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionService.ExtractionDelete", in)
	out := new(ExtractionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extractionService) ExtractionSwitchSave(ctx context.Context, in *ExtractionRequest, opts ...client.CallOption) (*ExtractionResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionService.ExtractionSwitchSave", in)
	out := new(ExtractionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extractionService) UserCustomExtractionAddress(ctx context.Context, in *ExtractionRequest, opts ...client.CallOption) (*ExtractionResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionService.UserCustomExtractionAddress", in)
	out := new(ExtractionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extractionService) Search(ctx context.Context, in *ExtractionRequest, opts ...client.CallOption) (*ExtractionResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionService.Search", in)
	out := new(ExtractionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExtractionService service

type ExtractionServiceHandler interface {
	//用户取货点首页
	ExtractionIndex(context.Context, *Extraction, *ExtractionIndexResponse) error
	//用户取货点申请信息
	ExtractionApplyInfo(context.Context, *Extraction, *ExtractionResponse) error
	//用户取货点申请保存
	ExtractionApplySave(context.Context, *Extraction, *ExtractionResponse) error
	//获取用户取货点数据
	ExtractionDetail(context.Context, *Extraction, *ExtractionResponse) error
	//取货点审核
	ExtractionAudit(context.Context, *Extraction, *ExtractionResponse) error
	//取货点解约
	ExtractionRelieve(context.Context, *ExtractionRequest, *ExtractionResponse) error
	//取货点删除
	ExtractionDelete(context.Context, *ExtractionRequest, *ExtractionResponse) error
	//取货点切换保存
	ExtractionSwitchSave(context.Context, *ExtractionRequest, *ExtractionResponse) error
	//获取用户自定义地址
	UserCustomExtractionAddress(context.Context, *ExtractionRequest, *ExtractionResponse) error
	//取货点查询
	Search(context.Context, *ExtractionRequest, *ExtractionResponse) error
}

func RegisterExtractionServiceHandler(s server.Server, hdlr ExtractionServiceHandler, opts ...server.HandlerOption) error {
	type extractionService interface {
		ExtractionIndex(ctx context.Context, in *Extraction, out *ExtractionIndexResponse) error
		ExtractionApplyInfo(ctx context.Context, in *Extraction, out *ExtractionResponse) error
		ExtractionApplySave(ctx context.Context, in *Extraction, out *ExtractionResponse) error
		ExtractionDetail(ctx context.Context, in *Extraction, out *ExtractionResponse) error
		ExtractionAudit(ctx context.Context, in *Extraction, out *ExtractionResponse) error
		ExtractionRelieve(ctx context.Context, in *ExtractionRequest, out *ExtractionResponse) error
		ExtractionDelete(ctx context.Context, in *ExtractionRequest, out *ExtractionResponse) error
		ExtractionSwitchSave(ctx context.Context, in *ExtractionRequest, out *ExtractionResponse) error
		UserCustomExtractionAddress(ctx context.Context, in *ExtractionRequest, out *ExtractionResponse) error
		Search(ctx context.Context, in *ExtractionRequest, out *ExtractionResponse) error
	}
	type ExtractionService struct {
		extractionService
	}
	h := &extractionServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ExtractionService{h}, opts...))
}

type extractionServiceHandler struct {
	ExtractionServiceHandler
}

func (h *extractionServiceHandler) ExtractionIndex(ctx context.Context, in *Extraction, out *ExtractionIndexResponse) error {
	return h.ExtractionServiceHandler.ExtractionIndex(ctx, in, out)
}

func (h *extractionServiceHandler) ExtractionApplyInfo(ctx context.Context, in *Extraction, out *ExtractionResponse) error {
	return h.ExtractionServiceHandler.ExtractionApplyInfo(ctx, in, out)
}

func (h *extractionServiceHandler) ExtractionApplySave(ctx context.Context, in *Extraction, out *ExtractionResponse) error {
	return h.ExtractionServiceHandler.ExtractionApplySave(ctx, in, out)
}

func (h *extractionServiceHandler) ExtractionDetail(ctx context.Context, in *Extraction, out *ExtractionResponse) error {
	return h.ExtractionServiceHandler.ExtractionDetail(ctx, in, out)
}

func (h *extractionServiceHandler) ExtractionAudit(ctx context.Context, in *Extraction, out *ExtractionResponse) error {
	return h.ExtractionServiceHandler.ExtractionAudit(ctx, in, out)
}

func (h *extractionServiceHandler) ExtractionRelieve(ctx context.Context, in *ExtractionRequest, out *ExtractionResponse) error {
	return h.ExtractionServiceHandler.ExtractionRelieve(ctx, in, out)
}

func (h *extractionServiceHandler) ExtractionDelete(ctx context.Context, in *ExtractionRequest, out *ExtractionResponse) error {
	return h.ExtractionServiceHandler.ExtractionDelete(ctx, in, out)
}

func (h *extractionServiceHandler) ExtractionSwitchSave(ctx context.Context, in *ExtractionRequest, out *ExtractionResponse) error {
	return h.ExtractionServiceHandler.ExtractionSwitchSave(ctx, in, out)
}

func (h *extractionServiceHandler) UserCustomExtractionAddress(ctx context.Context, in *ExtractionRequest, out *ExtractionResponse) error {
	return h.ExtractionServiceHandler.UserCustomExtractionAddress(ctx, in, out)
}

func (h *extractionServiceHandler) Search(ctx context.Context, in *ExtractionRequest, out *ExtractionResponse) error {
	return h.ExtractionServiceHandler.Search(ctx, in, out)
}
