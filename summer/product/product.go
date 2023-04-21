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
	service   = "rpc.xshop.product"
)

func Create(c *fiber.Ctx) error {
	d := etcd.NewDiscovery(Endpoints, service)
	resolver.Register(d)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 通过etcd注册中心和grpc服务建立连接
	conn, err := grpc.DialContext(ctx,
		fmt.Sprintf(d.Scheme()+":///"+service),
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
		return err
	}
	cancel()

	return c.JSON(r)
}

func Detail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	d := etcd.NewDiscovery(Endpoints, service)
	resolver.Register(d)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 通过etcd注册中心和grpc服务建立连接
	conn, err := grpc.DialContext(ctx,
		fmt.Sprintf(d.Scheme()+":///"+service),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalencingConfig": [{"round_robin": {}}]}`),
		grpc.WithBlock(),
	)
	if err != nil {
		fmt.Println("和rpc建立连接失败：", err)
	}

	client := pb.NewProductClient(conn)
	r, err := client.Detail(ctx, &pb.DetailRequest{Id: int32(id)})
	if err != nil {
		return err
	}
	cancel()

	return c.JSON(r)
}
