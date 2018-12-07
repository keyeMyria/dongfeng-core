// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/ilovelili/dongfeng-core/services/proto/api.proto

/*
Package dongfeng_svc_core_server is a generated protocol buffer package.

It is generated from these files:
	github.com/ilovelili/dongfeng-core/services/proto/api.proto

It has these top-level messages:
*/
package dongfeng_svc_core_server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import dongfeng_protobuf "github.com/ilovelili/dongfeng-protobuf"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = dongfeng_protobuf.UpdateAttendanceResponse{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Api service

type ApiService interface {
	Login(ctx context.Context, in *dongfeng_protobuf.LoginRequest, opts ...client.CallOption) (*dongfeng_protobuf.LoginResponse, error)
	Dashboard(ctx context.Context, in *dongfeng_protobuf.DashboardRequest, opts ...client.CallOption) (*dongfeng_protobuf.DashboardResponse, error)
	UpdateUser(ctx context.Context, in *dongfeng_protobuf.UpdateUserRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdateUserResponse, error)
	GetAttendance(ctx context.Context, in *dongfeng_protobuf.GetAttendanceRequest, opts ...client.CallOption) (*dongfeng_protobuf.GetAttendanceResponse, error)
	UpdateAttendance(ctx context.Context, in *dongfeng_protobuf.UpdateAttendanceRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdateAttendanceResponse, error)
}

type apiService struct {
	c    client.Client
	name string
}

func NewApiService(name string, c client.Client) ApiService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "dongfeng.svc.core.server"
	}
	return &apiService{
		c:    c,
		name: name,
	}
}

func (c *apiService) Login(ctx context.Context, in *dongfeng_protobuf.LoginRequest, opts ...client.CallOption) (*dongfeng_protobuf.LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Api.Login", in)
	out := new(dongfeng_protobuf.LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) Dashboard(ctx context.Context, in *dongfeng_protobuf.DashboardRequest, opts ...client.CallOption) (*dongfeng_protobuf.DashboardResponse, error) {
	req := c.c.NewRequest(c.name, "Api.Dashboard", in)
	out := new(dongfeng_protobuf.DashboardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) UpdateUser(ctx context.Context, in *dongfeng_protobuf.UpdateUserRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdateUserResponse, error) {
	req := c.c.NewRequest(c.name, "Api.UpdateUser", in)
	out := new(dongfeng_protobuf.UpdateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) GetAttendance(ctx context.Context, in *dongfeng_protobuf.GetAttendanceRequest, opts ...client.CallOption) (*dongfeng_protobuf.GetAttendanceResponse, error) {
	req := c.c.NewRequest(c.name, "Api.GetAttendance", in)
	out := new(dongfeng_protobuf.GetAttendanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) UpdateAttendance(ctx context.Context, in *dongfeng_protobuf.UpdateAttendanceRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdateAttendanceResponse, error) {
	req := c.c.NewRequest(c.name, "Api.UpdateAttendance", in)
	out := new(dongfeng_protobuf.UpdateAttendanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiHandler interface {
	Login(context.Context, *dongfeng_protobuf.LoginRequest, *dongfeng_protobuf.LoginResponse) error
	Dashboard(context.Context, *dongfeng_protobuf.DashboardRequest, *dongfeng_protobuf.DashboardResponse) error
	UpdateUser(context.Context, *dongfeng_protobuf.UpdateUserRequest, *dongfeng_protobuf.UpdateUserResponse) error
	GetAttendance(context.Context, *dongfeng_protobuf.GetAttendanceRequest, *dongfeng_protobuf.GetAttendanceResponse) error
	UpdateAttendance(context.Context, *dongfeng_protobuf.UpdateAttendanceRequest, *dongfeng_protobuf.UpdateAttendanceResponse) error
}

func RegisterApiHandler(s server.Server, hdlr ApiHandler, opts ...server.HandlerOption) {
	type api interface {
		Login(ctx context.Context, in *dongfeng_protobuf.LoginRequest, out *dongfeng_protobuf.LoginResponse) error
		Dashboard(ctx context.Context, in *dongfeng_protobuf.DashboardRequest, out *dongfeng_protobuf.DashboardResponse) error
		UpdateUser(ctx context.Context, in *dongfeng_protobuf.UpdateUserRequest, out *dongfeng_protobuf.UpdateUserResponse) error
		GetAttendance(ctx context.Context, in *dongfeng_protobuf.GetAttendanceRequest, out *dongfeng_protobuf.GetAttendanceResponse) error
		UpdateAttendance(ctx context.Context, in *dongfeng_protobuf.UpdateAttendanceRequest, out *dongfeng_protobuf.UpdateAttendanceResponse) error
	}
	type Api struct {
		api
	}
	h := &apiHandler{hdlr}
	s.Handle(s.NewHandler(&Api{h}, opts...))
}

type apiHandler struct {
	ApiHandler
}

func (h *apiHandler) Login(ctx context.Context, in *dongfeng_protobuf.LoginRequest, out *dongfeng_protobuf.LoginResponse) error {
	return h.ApiHandler.Login(ctx, in, out)
}

func (h *apiHandler) Dashboard(ctx context.Context, in *dongfeng_protobuf.DashboardRequest, out *dongfeng_protobuf.DashboardResponse) error {
	return h.ApiHandler.Dashboard(ctx, in, out)
}

func (h *apiHandler) UpdateUser(ctx context.Context, in *dongfeng_protobuf.UpdateUserRequest, out *dongfeng_protobuf.UpdateUserResponse) error {
	return h.ApiHandler.UpdateUser(ctx, in, out)
}

func (h *apiHandler) GetAttendance(ctx context.Context, in *dongfeng_protobuf.GetAttendanceRequest, out *dongfeng_protobuf.GetAttendanceResponse) error {
	return h.ApiHandler.GetAttendance(ctx, in, out)
}

func (h *apiHandler) UpdateAttendance(ctx context.Context, in *dongfeng_protobuf.UpdateAttendanceRequest, out *dongfeng_protobuf.UpdateAttendanceResponse) error {
	return h.ApiHandler.UpdateAttendance(ctx, in, out)
}
