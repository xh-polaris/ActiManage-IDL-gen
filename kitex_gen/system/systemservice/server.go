// Code generated by Kitex v0.12.3. DO NOT EDIT.
package systemservice

import (
	server "github.com/cloudwego/kitex/server"
	system "github.com/xh-polaris/ActiManage-IDL-gen/kitex_gen/system"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler system.SystemService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler system.SystemService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
