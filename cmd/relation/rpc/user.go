package rpc

// import (
// 	"time"

// 	"github.com/bdyc-org/dousheng/kitex_gen/user/userservice"
// 	"github.com/bdyc-org/dousheng/pkg/constants"
// 	"github.com/cloudwego/kitex/client"
// 	"github.com/cloudwego/kitex/pkg/retry"
// 	etcd "github.com/kitex-contrib/registry-etcd"
// )

// var userClient userservice.Client

// user服务发现
// func initUserRpc() {
// 	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
// 	if err != nil {
// 		panic(err)
// 	}

// 	c, err := userservice.NewClient(
// 		constants.UserServiceName,
// 		// client.WithMiddleware(middleware.CommonMiddleware),
// 		// client.WithInstanceMW(middleware.ClientMiddleware),
// 		client.WithMuxConnection(1),                       // mux
// 		client.WithRPCTimeout(3*time.Second),              // rpc timeout
// 		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
// 		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
// 		// client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
// 		client.WithResolver(r),                            // resolver
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	userClient = c
// }

// 发现token鉴权,成功返回user_id,否则返回0
// func Authentication(ctx context.Context, req *user.AuthenticationRequest) (int64, error) {
// 	resp, err := userClient.Authentication(ctx, req)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return resp.UserId, nil
// }

// 发现user的MGetUser
// func MGetUser(ctx context.Context, req *relation.QueryUserListRequest) ([]*user.User, error) {
// 	// 打包为user.MGetUserRequest
// 	r := pack.MGetUserReq(req.UserId, req.UserIds)
	
// 	resp, err := userClient.MGetUser(ctx, r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if resp.BaseResp.StatusCode != errno.SuccessCode {
// 		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
// 	}
// 	return resp.UserList, nil
// }