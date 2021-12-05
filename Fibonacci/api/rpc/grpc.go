// Package rpc contains all methods for working rpc server.
package rpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/rauf95/rauf/core"
	pb "github.com/rauf95/rauf/proto/api/v1"
)

type api struct{}

func (a *api) Fibonacci(_ context.Context, request *pb.FibonacciRequest) (*pb.FibonacciResponse, error) {
	fibonacci := core.Fibonacci(int(request.Arg))
//вызывается в сервер джрпсишном сервере
	var res []*pb.FibonacciValue
	for _, value := range fibonacci {
		res = append(res, &pb.FibonacciValue{
			Number: int64(value.Number),
			Value:  int64(value.Value),
		})
	}
// конвертация типов
	return &pb.FibonacciResponse{Result: res}, nil
}

// New creates and returns gRPC server.
func New() *grpc.Server {
	srv := grpc.NewServer()
	// Регистрация твоего grpc сервера
	pb.RegisterFibonacciServiceServer(srv, &api{})

	return srv
}
