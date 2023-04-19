package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"xshop/product/internal/data"
	"xshop/product/internal/server"

	"github.com/happyin3/x/autumn/registry/etcd"
)

var (
	port      = flag.Int("port", 50051, "The server port")
	Endpoints = []string{"127.0.0.1:12379", "127.0.0.1:22379", "127.0.0.1:32379"}
)

func main() {
	flag.Parse()

	data.Connect()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	newService, err := etcd.NewService(etcd.ServiceInfo{
		Name: "rpc.xshop.product",
		Ip:   "127.0.0.1:" + fmt.Sprintf("%d", *port),
	}, Endpoints)
	if err != nil {
		fmt.Println("添加到etcd失败：", err)
	}

	fmt.Println("xshop product")

	go func() {
		err = newService.Start(context.Background())
		if err != nil {
			fmt.Println("开启服务注册失败：", err)
		}
	}()

	s := server.NewGRPCServer()
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
