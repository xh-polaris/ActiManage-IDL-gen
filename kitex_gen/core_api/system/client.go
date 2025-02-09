// Code generated by Kitex v0.12.1. DO NOT EDIT.

package system

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	core_api "github.com/xh-polaris/ActiManage-IDL-gen/kitex_gen/core_api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SystemLogin(ctx context.Context, Req *core_api.SystemLoginReq, callOptions ...callopt.Option) (r *core_api.SystemLoginResp, err error)
	SystemListMerchant(ctx context.Context, Req *core_api.SystemListMerchantsReq, callOptions ...callopt.Option) (r *core_api.SystemListMerchantsReq, err error)
	SystemGetMerchant(ctx context.Context, Req *core_api.SystemGetMerchantReq, callOptions ...callopt.Option) (r *core_api.SystemGetMerchantResp, err error)
	SystemCreateMerchant(ctx context.Context, Req *core_api.SystemCreateMerchantReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	SystemUpdateMerchant(ctx context.Context, Req *core_api.SystemUpdateMerchantReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	SystemGetDashboard(ctx context.Context, Req *core_api.SystemGetDashboardReq, callOptions ...callopt.Option) (r *core_api.SystemGetDashboardResp, err error)
	SystemGetOverallDashboard(ctx context.Context, Req *core_api.SystemGetOverallDashboardReq, callOptions ...callopt.Option) (r *core_api.SystemGetOverallDashboardResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kSystemClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kSystemClient struct {
	*kClient
}

func (p *kSystemClient) SystemLogin(ctx context.Context, Req *core_api.SystemLoginReq, callOptions ...callopt.Option) (r *core_api.SystemLoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SystemLogin(ctx, Req)
}

func (p *kSystemClient) SystemListMerchant(ctx context.Context, Req *core_api.SystemListMerchantsReq, callOptions ...callopt.Option) (r *core_api.SystemListMerchantsReq, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SystemListMerchant(ctx, Req)
}

func (p *kSystemClient) SystemGetMerchant(ctx context.Context, Req *core_api.SystemGetMerchantReq, callOptions ...callopt.Option) (r *core_api.SystemGetMerchantResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SystemGetMerchant(ctx, Req)
}

func (p *kSystemClient) SystemCreateMerchant(ctx context.Context, Req *core_api.SystemCreateMerchantReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SystemCreateMerchant(ctx, Req)
}

func (p *kSystemClient) SystemUpdateMerchant(ctx context.Context, Req *core_api.SystemUpdateMerchantReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SystemUpdateMerchant(ctx, Req)
}

func (p *kSystemClient) SystemGetDashboard(ctx context.Context, Req *core_api.SystemGetDashboardReq, callOptions ...callopt.Option) (r *core_api.SystemGetDashboardResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SystemGetDashboard(ctx, Req)
}

func (p *kSystemClient) SystemGetOverallDashboard(ctx context.Context, Req *core_api.SystemGetOverallDashboardReq, callOptions ...callopt.Option) (r *core_api.SystemGetOverallDashboardResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SystemGetOverallDashboard(ctx, Req)
}
