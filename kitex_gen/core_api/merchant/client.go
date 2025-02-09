// Code generated by Kitex v0.12.1. DO NOT EDIT.

package merchant

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	core_api "github.com/xh-polaris/ActiManage-IDL-gen/kitex_gen/core_api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MerchantListActivities(ctx context.Context, Req *core_api.MerchantListActivitiesReq, callOptions ...callopt.Option) (r *core_api.MerchantListActivitiesResp, err error)
	MerchantCreateActivity(ctx context.Context, Req *core_api.MerchantCreateActivityReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	MerchantDeleteActivity(ctx context.Context, Req *core_api.MerchantDeleteActivityReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	MerchantTopActivity(ctx context.Context, Req *core_api.MerchantTopActivityReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	MerchantLogin(ctx context.Context, Req *core_api.MerchantLoginReq, callOptions ...callopt.Option) (r *core_api.MerchantLoginResp, err error)
	MerchantGetSetting(ctx context.Context, Req *core_api.MerchantGetSettingReq, callOptions ...callopt.Option) (r *core_api.MerchantGetSettingResp, err error)
	MerchantUpdateSetting(ctx context.Context, Req *core_api.MerchantUpdateSettingReq, callOptions ...callopt.Option) (r *core_api.MerchantGetSettingResp, err error)
	MerchantGetBookRecords(ctx context.Context, Req *core_api.MerchantListBookRecordsReq, callOptions ...callopt.Option) (r *core_api.MerchantListBookRecordsResp, err error)
	MerchantUpdateInfo(ctx context.Context, Req *core_api.MerchantUpdateInfoReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	MerchantGetInfo(ctx context.Context, Req *core_api.MerchantGetInfoReq, callOptions ...callopt.Option) (r *core_api.MerchantGetInfoResp, err error)
	MerchantSetPassword(ctx context.Context, Req *core_api.MerchantSetPasswordReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
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
	return &kMerchantClient{
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

type kMerchantClient struct {
	*kClient
}

func (p *kMerchantClient) MerchantListActivities(ctx context.Context, Req *core_api.MerchantListActivitiesReq, callOptions ...callopt.Option) (r *core_api.MerchantListActivitiesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantListActivities(ctx, Req)
}

func (p *kMerchantClient) MerchantCreateActivity(ctx context.Context, Req *core_api.MerchantCreateActivityReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantCreateActivity(ctx, Req)
}

func (p *kMerchantClient) MerchantDeleteActivity(ctx context.Context, Req *core_api.MerchantDeleteActivityReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantDeleteActivity(ctx, Req)
}

func (p *kMerchantClient) MerchantTopActivity(ctx context.Context, Req *core_api.MerchantTopActivityReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantTopActivity(ctx, Req)
}

func (p *kMerchantClient) MerchantLogin(ctx context.Context, Req *core_api.MerchantLoginReq, callOptions ...callopt.Option) (r *core_api.MerchantLoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantLogin(ctx, Req)
}

func (p *kMerchantClient) MerchantGetSetting(ctx context.Context, Req *core_api.MerchantGetSettingReq, callOptions ...callopt.Option) (r *core_api.MerchantGetSettingResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantGetSetting(ctx, Req)
}

func (p *kMerchantClient) MerchantUpdateSetting(ctx context.Context, Req *core_api.MerchantUpdateSettingReq, callOptions ...callopt.Option) (r *core_api.MerchantGetSettingResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantUpdateSetting(ctx, Req)
}

func (p *kMerchantClient) MerchantGetBookRecords(ctx context.Context, Req *core_api.MerchantListBookRecordsReq, callOptions ...callopt.Option) (r *core_api.MerchantListBookRecordsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantGetBookRecords(ctx, Req)
}

func (p *kMerchantClient) MerchantUpdateInfo(ctx context.Context, Req *core_api.MerchantUpdateInfoReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantUpdateInfo(ctx, Req)
}

func (p *kMerchantClient) MerchantGetInfo(ctx context.Context, Req *core_api.MerchantGetInfoReq, callOptions ...callopt.Option) (r *core_api.MerchantGetInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantGetInfo(ctx, Req)
}

func (p *kMerchantClient) MerchantSetPassword(ctx context.Context, Req *core_api.MerchantSetPasswordReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantSetPassword(ctx, Req)
}
