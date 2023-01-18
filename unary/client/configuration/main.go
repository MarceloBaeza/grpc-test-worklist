package configuration

import (
	"sync"

	pb "github.com/MarceloBaeza/grpc-test-worklist/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ConnectionGrpc struct {
	Connection *grpc.ClientConn
	Client     pb.WorklistServiceClient
}

var instace *ConnectionGrpc
var once sync.Once

func NewConnection() *ConnectionGrpc {
	once.Do(func() {
		conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic("cannot connect with server " + err.Error())
		}

		serviceClient := pb.NewWorklistServiceClient(conn)
		instace = &ConnectionGrpc{
			Connection: conn,
			Client:     serviceClient,
		}
	})
	return instace
}
