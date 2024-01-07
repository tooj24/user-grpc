package grpc_test

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-clean/proto"
	"testing"
)

func TestCreateUser(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	request := &pb.CreateUserRequest{
		Name:  "test",
		Email: "test@gmail.com",
	}

	res, err := client.Create(context.Background(), request)
	if err != nil {
		t.Fatalf("CREATE FAILED: %v", err)
	}

	if res.Email != request.Email {
		t.Errorf("CREATE returned incorrect email, expected %s got %s", request.Email, res.Email)
	}

	if res.Name != request.Name {
		t.Errorf("CREATE returned incorrect Name, expected %s got %s", request.Name, res.Name)
	}

	if res.GetId() == "" {
		t.Error("CREATE function did not returned id as the response")
	}
}
