package main

import (
	"context"
	"log"
	"net"

	pb "github.com/MarceloBaeza/grpc-test-worklist/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedWorklistServiceServer
}

func (s *Server) CreateListWork(ctx context.Context, req *pb.InputRequestCreateListWork) (*pb.WorkList, error) {
	return nil, nil
}
func (s *Server) CreateWorkOfListWorks(ctx context.Context, req *pb.RequestWorkAssociateListWorks) (*emptypb.Empty, error) {
	return nil, nil
}
func (s *Server) CreateWork(ctx context.Context, req *pb.RequestCreateWork) (*emptypb.Empty, error) {
	log.Printf("input new work -> %v", req.NewWork)
	return new(emptypb.Empty), nil
}
func (s *Server) DeleteWork(ctx context.Context, req *pb.RequestDeleteWork) (*emptypb.Empty, error) {
	return nil, nil
}
func (s *Server) ListWorks(ctx context.Context, req *emptypb.Empty) (*pb.ResponseWorkList, error) {
	return nil, nil
}

func main() {
	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterWorklistServiceServer(serv, &Server{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}

}
