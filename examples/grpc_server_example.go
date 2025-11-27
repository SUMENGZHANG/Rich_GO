package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "rich_go/api/proto/examples" // 根据实际生成的代码路径调整
)

// UserServiceServer 实现 UserService 接口
type server struct {
	pb.UnimplementedUserServiceServer
}

// GetUser 获取用户信息
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// 模拟从数据库获取用户
	user := &pb.User{
		Id:        req.UserId,
		Name:      "用户" + string(rune(req.UserId)),
		Email:     "user" + string(rune(req.UserId)) + "@example.com",
		CreatedAt: 1234567890,
	}

	return &pb.GetUserResponse{User: user}, nil
}

// CreateUser 创建用户
func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// 模拟创建用户
	user := &pb.User{
		Id:        1,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: 1234567890,
	}

	return &pb.CreateUserResponse{User: user}, nil
}

// ListUsers 流式返回用户列表
func (s *server) ListUsers(req *pb.ListUsersRequest, stream pb.UserService_ListUsersServer) error {
	// 模拟用户列表
	users := []*pb.User{
		{Id: 1, Name: "用户1", Email: "user1@example.com", CreatedAt: 1234567890},
		{Id: 2, Name: "用户2", Email: "user2@example.com", CreatedAt: 1234567891},
		{Id: 3, Name: "用户3", Email: "user3@example.com", CreatedAt: 1234567892},
	}

	for _, user := range users {
		if err := stream.Send(user); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()

	// 注册服务
	pb.RegisterUserServiceServer(s, &server{})

	log.Println("gRPC 服务器启动在 :50051")

	// 启动服务器
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

