package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-test/proto" // 引入编译生成的包
	"log"
	"net"
	"time"
)

//通过一个结构体，实现proto中定义的所有服务
type ServeRoute struct{}

func (h ServeRoute) Serve1(ctx context.Context, in *pb.Name) (*pb.Msg1, error) {
	log.Println("serve 1 works: get name: ", in.Name)
	resp := &pb.Msg1{Message: "this is serve 1", Time: time.Now().UnixNano() / 1e6}
	return resp, nil
}

func (h ServeRoute) Serve2(ctx context.Context, in *pb.Name) (*pb.Msg2, error) {
	log.Println("serve 2 works, get name: ", in.Name)
	resp := &pb.Msg2{
		Message: &pb.Msg1{Message: "this is serve 2"},
	}
	return resp, nil
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:50052") // Address gRPC服务地址
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 与http的注册路由类似，此处将所有服务注册到grpc服务器上，
	pb.RegisterServeRouteServer(s, ServeRoute{})
	log.Println("grpc serve running")
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
