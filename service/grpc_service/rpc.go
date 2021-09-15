package grpc_service

import (
	"context"
	"fmt"
	"github.com/AzusaChino/ficus/pkg/conf"
	"github.com/AzusaChino/ficus/pkg/rpc"
	pb "github.com/AzusaChino/ficus/service/grpc_service/proto/hello"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"log"
	"time"
)

func DoHello(msg string, ctx *fasthttp.RequestCtx) string {
	var addr = fmt.Sprintf("%s:%s", conf.GrpcConfig.Server, conf.GrpcConfig.Port)
	// fetch custom values from ctx.Locals()
	c := ctx.UserValue("ctx").(context.Context)
	tracer := ctx.UserValue("tracer").(opentracing.Tracer)
	// wrap grpc client with interceptor
	conn, err := grpc.DialContext(c, addr, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithUnaryInterceptor(rpc.ClientInterceptor(tracer)))
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
