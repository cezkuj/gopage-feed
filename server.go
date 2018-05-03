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

func (s *gopageServer) get1(ctx context.Context) (*pb.Number, error) {
	return &pb.Feature{2}, nil
}

func (s *gopageServer) get2(ctx context.Context) (*pb.Number, error) {
	return &pb.Number{2}, nil
}

func main() {
	http.HandleFunc("/get1", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "1")
	})
	http.HandleFunc("/get2", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "2")
	})
	go log.Fatal(http.ListenAndServe(":8001", nil))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8002))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, &routeGuideServer{})
	grpcServer.Serve(lis)

}
