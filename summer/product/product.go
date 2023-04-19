package product

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	pb "xshop/summer/protos"

	"github.com/gofiber/fiber/v2"
	"github.com/happyin3/x/autumn/registry/etcd"
)

var (
	Endpoints = []string{"127.0.0.1:12379", "127.0.0.1:22379", "127.0.0.1:32379"}
)

func Hello(c *fiber.Ctx) error {
	d := etcd.NewDiscovery(Endpoints, "rpc.xshop.product")
	resolver.Register(d)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 通过etcd注册中心和grpc服务建立连接
	conn, err := grpc.DialContext(ctx,
		fmt.Sprintf(d.Scheme()+":///"+"rpc.xshop.product"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalencingConfig": [{"round_robin": {}}]}`),
		grpc.WithBlock(),
	)
	if err != nil {
		fmt.Println("和rpc建立连接失败：", err)
	}

	client := pb.NewProductClient(conn)
	r, err := client.Create(ctx, &pb.CreateRequest{Name: c.Params("name"), Desc: c.Params("name"), Stock: 1, Amount: 2, Status: 3})
	if err != nil {
		fmt.Println("could not greet: ", err)
		return err
	}
	fmt.Println(r.GetId())
	cancel()

	return c.SendString("1")
}
