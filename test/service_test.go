package test

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-gRPC/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
)

func TestUser(t *testing.T) {
	user := &service.User1{
		Username: "BoyChai",
		Age:      18,
	}
	// 序列化过程
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}

	// 反序列化
	newUser := &service.User1{}
	proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}

	// 输出
	fmt.Println(newUser)
}

func TestProdService(t *testing.T) {
	// 新建一个gprc服务
	server := grpc.NewServer()
	// 注册
	service.RegisterProdServiceServer(server, service.ProductService)
	// 启动服务
	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		t.Fatal("启动服务出错", err)
	}
	err = server.Serve(listen)
	if err != nil {
		t.Fatal("启动服务出错", err)
	}
	fmt.Println("启动grpc服务成功")
}

func TestClientProdService(t *testing.T) {
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("链接出错:", err)
	}
	// 退出时关闭
	defer conn.Close()
	client := service.NewProdServiceClient(conn)
	stock, err := client.GetProductStock(context.Background(), &service.ProductRequest{ProdId: 233})
	if err != nil {
		t.Fatal("链接出错:", err)
	}
	fmt.Println(stock.ProdStock)
}
