package main

import (
	"log"
	"net"

	"github.com/bdyc-org/dousheng/cmd/relation/rpc"
	"github.com/bdyc-org/dousheng/cmd/relation/dal"
	relation "github.com/bdyc-org/dousheng/kitex_gen/relation/relationservice"
	"github.com/bdyc-org/dousheng/pkg/constants"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	rpc.InitRPC()
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constants.RelationResolveTCPAddr)
	if err != nil {
		panic(err)
	}

	Init()

	svr := relation.NewServer(new(RelationServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.RelationServiceName}), // server name
		// server.WithMiddleware(middleware.CommonMiddleware),                                             // middleWare
		// server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		// server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		// server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r),                                             // registry
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
