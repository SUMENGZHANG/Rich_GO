package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "rich_go/api/proto/examples" // 根据实际生成的代码路径调整
)

func main() {
	// 连接到 gRPC 服务器
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用 GetUser
	userResp, err := client.GetUser(ctx, &pb.GetUserRequest{UserId: 1})
	if err != nil {
		log.Fatalf("GetUser 调用失败: %v", err)
	}
	log.Printf("用户信息: %v", userResp.User)

	// 调用 CreateUser
	createResp, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  "新用户",
		Email: "newuser@example.com",
	})
	if err != nil {
		log.Fatalf("CreateUser 调用失败: %v", err)
	}
	log.Printf("创建的用户: %v", createResp.User)

	// 调用流式接口 ListUsers
	stream, err := client.ListUsers(ctx, &pb.ListUsersRequest{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		log.Fatalf("ListUsers 调用失败: %v", err)
	}

	for {
		user, err := stream.Recv()
		if err != nil {
			break
		}
		log.Printf("收到用户: %v", user)
	}
}

