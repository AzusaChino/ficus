package grpc_service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/azusachino/ficus/pkg/conf"
	"github.com/azusachino/ficus/pkg/etcd"
	"github.com/azusachino/ficus/pkg/rpc"
	pb "github.com/azusachino/ficus/service/grpc_service/proto/hello"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/valyala/fasthttp"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
)

var serviceName = "myrica-grpc-server"

func DoHello(msg string, ctx *fasthttp.RequestCtx) string {
	var addr = fmt.Sprintf("%s:%s", conf.GrpcConfig.Server, conf.GrpcConfig.Port)
	var opts []grpc.DialOption
	// use etcd as service discovery
	var target = fmt.Sprintf("/etcdv3://ficus/grpc/%s", serviceName)
	// TODO go through etcd
	if addr == "" {
		addr = target
		builder, err := resolver.NewBuilder(etcd.Client)
		if err != nil {
			log.Panic(err)
		}
		opts = append(opts, grpc.WithResolvers(builder))
	}
	// fetch custom values from ctx.Locals()
	c := ctx.UserValue("ctx").(context.Context)
	tracer := ctx.UserValue("tracer").(opentracing.Tracer)
	opts = append(opts, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithUnaryInterceptor(rpc.ClientInterceptor(tracer)))
	// wrap grpc client with interceptor
	conn, err := grpc.DialContext(c, addr, opts...)
	if err != nil {
		log.Panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Panic(err)
		}
	}(conn)

	client := pb.NewHelloServiceClient(conn)

	r, err := client.SayHello(c, &pb.Request{
		Id:   uuid.New().String(),
		Msg:  msg,
		Date: time.Now().Format("20060102"),
	})
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Hello From GRPC server, code: %v, msg: %s", r.Code, r.Msg)
	return r.String()
}
