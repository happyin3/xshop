package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	pb "helloworld/helloworld"

	"helloworld/server"
)

const defaultName = "world"

var (
	addr      = flag.String("addr", "localhost:50051", "the address to connect to")
	name      = flag.String("name", defaultName, "Name to greet")
	Endpoints = []string{"127.0.0.1:12379", "127.0.0.1:22379", "127.0.0.1:32379"}
)

func main() {
	flag.Parse()

	d := server.NewDiscovery(Endpoints, "rpc.xshop.hello")
	resolver.Register(d)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		// 通过etcd注册中心和grpc服务建立连接
		conn, err := grpc.DialContext(ctx,
			fmt.Sprintf(d.Scheme()+":///"+"rpc.xshop.hello"),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(`{"loadBalencingConfig": [{"round_robin": {}}]}`),
			grpc.WithBlock(),
		)
		if err != nil {
			fmt.Println("和rpc建立连接失败：", err)
			return
		}

		client := pb.NewGreeterClient(conn)
		r, err := client.SayHello(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			fmt.Println("could not greet: ", err)
			return
		}
		fmt.Println("Greeting: ", r.GetMessage())

		time.Sleep(3 * time.Second)
		cancel()
	}

	/*
		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewGreeterClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	*/
}
