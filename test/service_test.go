package test

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-gRPC/service"
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

func TestUser2(t *testing.T) {

}
