package tests

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/google/uuid"

	kitlog "github.com/go-kit/kit/log"
	transportgrpc "github.com/mahdi-asadzadeh/go-kit-accounts/pkg/transport/grpc"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/transport/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()
	var logger kitlog.Logger
	pb.RegisterUSerServiceServer(server, transportgrpc.NewGrpcUserServer(testCommon.UserEnd, logger))

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestUserTransportgRPC(t *testing.T) {
	// Implement user data
	email := uuid.New().String()
	fullName := uuid.New().String()
	password := uuid.New().String()
	var newFullName string
	var userID int64

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewUSerServiceClient(conn)

	t.Run("create-user", func(t *testing.T) {
		req := pb.CreateUserRequest{Email: email, FullName: fullName, Password: password}
		res, err := client.CreateUser(ctx, &req)

		if err != nil {
			t.Errorf("expected %v received %v", nil, err)
		}
		
		// Check response
		if res.GetEmail() != email {
			t.Errorf("want %s have %s", email, res.GetEmail())
		}
		if res.GetFullName() != fullName {
			t.Errorf("want %s, have %s", fullName, res.GetFullName())
		}
		userID = res.GetId()
	})

	t.Run("update-user", func(t *testing.T) {
		newFullName = "My test"
		req := pb.UpdateUserRequest{Email: email, FullName: newFullName}
		res, err := client.UpdateUser(ctx, &req)

		if err != nil {
			t.Errorf("expected %v received %v", nil, err)
		}

		// Check response
		if res.GetFullName() != newFullName {
			t.Errorf("want %s, have %s", fullName, res.GetFullName())
		}
	})

	t.Run("get-user", func(t *testing.T) {
		req := pb.GetUserRequest{Email: email}
		res, err := client.GetUser(ctx, &req)

		if err != nil {
			t.Errorf("expected %v received %v", nil, err)
		}

		// Check response
		if res.GetId() != userID {
			t.Errorf("want %d, have %d", userID, res.GetId())
		}
		if res.GetEmail() != email {
			t.Errorf("want %s, have %s", email, res.GetEmail())
		}
		if res.GetFullName() != newFullName {
			t.Errorf("want %s, have %s", newFullName, res.GetFullName())
		}
	})

	t.Run("login", func(t *testing.T) {
		req := pb.LoginUserRequest{Email: email, Password: password}
		res, err := client.Login(ctx, &req)

		if err != nil {
			t.Errorf("expected %v received %v", nil, err)
		}

		// Check response
		if res.GetToken() == "" {
			t.Errorf("want %s, have %s", "Not nill", res.GetToken())
		}
	})

	t.Run("delete-user", func(t *testing.T) {
		req := pb.DeleteUserRequest{Email: email}
		res, err := client.DeleteUser(ctx, &req)

		if err != nil {
			t.Errorf("expected %v received %v", nil, err)
		}

		// Check response
		if res.GetOk() != true {
			t.Errorf("want %t, have %t", true, res.GetOk())
		}
	})
}
