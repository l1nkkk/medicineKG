package main

import (
	"context"
	"fmt"
	"github.com/l1nkkk/medicineKG/demo/rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type GrpcSecWrap struct {
	helloServe proto.HelloWorldClient
	conn       *grpc.ClientConn
}

var (
	Gsw *GrpcSecWrap
)

func InitGRPC() {
	conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	Gsw = &GrpcSecWrap{conn: conn}
	Gsw.helloServe = proto.NewHelloWorldClient(conn)
}

func (g *GrpcSecWrap) Close() {
	if err := g.conn.Close(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	InitGRPC()
	defer Gsw.Close()
	name := "linqing"
	rtn, err := Gsw.helloServe.SayHelloAgain(context.Background(), &proto.HelloRequest{Name: &name})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rtn)
}
