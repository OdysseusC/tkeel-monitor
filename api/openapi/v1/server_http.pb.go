// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http 0.1.0

package v1

import (
	context "context"
	go_restful "github.com/emicklei/go-restful"
	errors "github.com/tkeel-io/kit/errors"
	result "github.com/tkeel-io/kit/result"
	v1 "github.com/tkeel-io/tkeel-interface/openapi/v1"
	protojson "google.golang.org/protobuf/encoding/protojson"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
)

import transportHTTP "github.com/tkeel-io/kit/transport/http"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the tkeel package it is being compiled against.
// import package.context.http.anypb.result.protojson.go_restful.errors.emptypb.

var (
	_ = protojson.MarshalOptions{}
	_ = anypb.Any{}
	_ = emptypb.Empty{}
)

type OpenapiHTTPServer interface {
	AddonsIdentify(context.Context, *v1.AddonsIdentifyRequest) (*v1.AddonsIdentifyResponse, error)
	Identify(context.Context, *emptypb.Empty) (*v1.IdentifyResponse, error)
	Status(context.Context, *emptypb.Empty) (*v1.StatusResponse, error)
	TenantDisable(context.Context, *v1.TenantDisableRequest) (*v1.TenantDisableResponse, error)
	TenantEnable(context.Context, *v1.TenantEnableRequest) (*v1.TenantEnableResponse, error)
}

type OpenapiHTTPHandler struct {
	srv OpenapiHTTPServer
}

func newOpenapiHTTPHandler(s OpenapiHTTPServer) *OpenapiHTTPHandler {
	return &OpenapiHTTPHandler{srv: s}
}

func (h *OpenapiHTTPHandler) AddonsIdentify(req *go_restful.Request, resp *go_restful.Response) {
	in := v1.AddonsIdentifyRequest{}
	if err := transportHTTP.GetBody(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.AddonsIdentify(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		if httpCode == http.StatusMovedPermanently {
			resp.Header().Set("Location", tErr.Message)
		}
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	resp.WriteHeaderAndJson(http.StatusOK, out, "application/json")
}

func (h *OpenapiHTTPHandler) Identify(req *go_restful.Request, resp *go_restful.Response) {
	in := emptypb.Empty{}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.Identify(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		if httpCode == http.StatusMovedPermanently {
			resp.Header().Set("Location", tErr.Message)
		}
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	resp.WriteHeaderAndJson(http.StatusOK, out, "application/json")
}

func (h *OpenapiHTTPHandler) Status(req *go_restful.Request, resp *go_restful.Response) {
	in := emptypb.Empty{}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.Status(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		if httpCode == http.StatusMovedPermanently {
			resp.Header().Set("Location", tErr.Message)
		}
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	resp.WriteHeaderAndJson(http.StatusOK, out, "application/json")
}

func (h *OpenapiHTTPHandler) TenantDisable(req *go_restful.Request, resp *go_restful.Response) {
	in := v1.TenantDisableRequest{}
	if err := transportHTTP.GetBody(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.TenantDisable(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		if httpCode == http.StatusMovedPermanently {
			resp.Header().Set("Location", tErr.Message)
		}
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	resp.WriteHeaderAndJson(http.StatusOK, out, "application/json")
}

func (h *OpenapiHTTPHandler) TenantEnable(req *go_restful.Request, resp *go_restful.Response) {
	in := v1.TenantEnableRequest{}
	if err := transportHTTP.GetBody(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(errors.InternalError.Reason, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.TenantEnable(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		if httpCode == http.StatusMovedPermanently {
			resp.Header().Set("Location", tErr.Message)
		}
		resp.WriteHeaderAndJson(httpCode,
			result.Set(tErr.Reason, tErr.Message, out), "application/json")
		return
	}
	resp.WriteHeaderAndJson(http.StatusOK, out, "application/json")
}

func RegisterOpenapiHTTPServer(container *go_restful.Container, srv OpenapiHTTPServer) {
	var ws *go_restful.WebService
	for _, v := range container.RegisteredWebServices() {
		if v.RootPath() == "/v1" {
			ws = v
			break
		}
	}
	if ws == nil {
		ws = new(go_restful.WebService)
		ws.ApiVersion("/v1")
		ws.Path("/v1").Produces(go_restful.MIME_JSON)
		container.Add(ws)
	}

	handler := newOpenapiHTTPHandler(srv)
	ws.Route(ws.GET("/identify").
		To(handler.Identify))
	ws.Route(ws.POST("/addons/identify").
		To(handler.AddonsIdentify))
	ws.Route(ws.POST("/tenant/enable").
		To(handler.TenantEnable))
	ws.Route(ws.POST("/tenant/disable").
		To(handler.TenantDisable))
	ws.Route(ws.GET("/status").
		To(handler.Status))
}
