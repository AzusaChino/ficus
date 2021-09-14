package grpc_service

import (
	"context"
	"fmt"
	"github.com/AzusaChino/ficus/pkg/conf"
	pb "github.com/AzusaChino/ficus/service/grpc_service/proto/hello"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"time"
)

func SayHello(msg string) (string, error) {
	// move inside (late initializing)
	var addr = fmt.Sprintf("%s:%s", conf.GrpcConfig.Server, conf.GrpcConfig.Port)
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	client := pb.NewHelloServiceClient(conn)

	r, err := client.SayHello(ctx, &pb.Request{
		Id:   uuid.New().String(),
		Msg:  msg,
		Date: time.Now().Format("20060102"),
	})
	if err != nil {
		return "", err
	}
	log.Printf("Hello From GRPC server, code: %v, msg: %s", r.Code, r.Msg)
	return r.String(), nil
}
