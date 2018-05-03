package main

import (
	"io"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/cezkuj/gopage-feed/grpc"
)

type gopageServer struct {
}

func (s *gopageServer) Get1(ctx context.Context, n *pb.GetParams) (*pb.Number, error) {
	return &pb.Number{Value: 1}, nil
}

func (s *gopageServer) Get2(ctx context.Context, n *pb.GetParams) (*pb.Number, error) {
	return &pb.Number{Value: 2}, nil
}

func main() {
	http.HandleFunc("/get1", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "1")
	})
	http.HandleFunc("/get2", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "2")
	})
	go func(){ log.Fatal(http.ListenAndServe(":8001", nil))}()
	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGopageServer(grpcServer, &gopageServer{})
	grpcServer.Serve(lis)

}
