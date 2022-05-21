// Code generated by Kitex v0.3.1. DO NOT EDIT.

package userservice

import (
	"context"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateUser": kitex.NewMethodInfo(createUserHandler, newUserServiceCreateUserArgs, newUserServiceCreateUserResult, false),
		"CheckUser":  kitex.NewMethodInfo(checkUserHandler, newUserServiceCheckUserArgs, newUserServiceCheckUserResult, false),
		"MGetUser":   kitex.NewMethodInfo(mGetUserHandler, newUserServiceMGetUserArgs, newUserServiceMGetUserResult, false),
		"NewFollow":  kitex.NewMethodInfo(newFollow_Handler, newUserServiceNewFollowArgs, newUserServiceNewFollowResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceCreateUserArgs)
	realResult := result.(*user.UserServiceCreateUserResult)
	success, err := handler.(user.UserService).CreateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCreateUserArgs() interface{} {
	return user.NewUserServiceCreateUserArgs()
}

func newUserServiceCreateUserResult() interface{} {
	return user.NewUserServiceCreateUserResult()
}

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceCheckUserArgs)
	realResult := result.(*user.UserServiceCheckUserResult)
	success, err := handler.(user.UserService).CheckUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCheckUserArgs() interface{} {
	return user.NewUserServiceCheckUserArgs()
}

func newUserServiceCheckUserResult() interface{} {
	return user.NewUserServiceCheckUserResult()
}

func mGetUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceMGetUserArgs)
	realResult := result.(*user.UserServiceMGetUserResult)
	success, err := handler.(user.UserService).MGetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceMGetUserArgs() interface{} {
	return user.NewUserServiceMGetUserArgs()
}

func newUserServiceMGetUserResult() interface{} {
	return user.NewUserServiceMGetUserResult()
}

func newFollow_Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceNewFollowArgs)
	realResult := result.(*user.UserServiceNewFollowResult)
	success, err := handler.(user.UserService).NewFollow_(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceNewFollowArgs() interface{} {
	return user.NewUserServiceNewFollowArgs()
}

func newUserServiceNewFollowResult() interface{} {
	return user.NewUserServiceNewFollowResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateUser(ctx context.Context, req *user.CreateUserRequest) (r *user.CreateUserResponse, err error) {
	var _args user.UserServiceCreateUserArgs
	_args.Req = req
	var _result user.UserServiceCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckUser(ctx context.Context, req *user.CheckUserRequest) (r *user.CheckUserResponse, err error) {
	var _args user.UserServiceCheckUserArgs
	_args.Req = req
	var _result user.UserServiceCheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetUser(ctx context.Context, req *user.MGetUserRequest) (r *user.MGetUserResponse, err error) {
	var _args user.UserServiceMGetUserArgs
	_args.Req = req
	var _result user.UserServiceMGetUserResult
	if err = p.c.Call(ctx, "MGetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) NewFollow_(ctx context.Context, req *user.NewFollowRequest_) (r *user.NewFollowResponse_, err error) {
	var _args user.UserServiceNewFollowArgs
	_args.Req = req
	var _result user.UserServiceNewFollowResult
	if err = p.c.Call(ctx, "NewFollow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
