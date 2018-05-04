package main

import (
	"io"
	"log"
	"math"
	"net"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/cezkuj/gopage-feed/grpc"
)

type gopageServer struct {
	i1 int
	i2 int
}

func (s *gopageServer) Get1(ctx context.Context, n *pb.GetParams) (*pb.Number, error) {
	s.i1 += 1
	if math.Mod(float64(s.i1), 1000.0) == 0.0 {
	log.Printf("grpc i1: %v", s.i1)
	}
	return &pb.Number{Value: 1}, nil
}

func (s *gopageServer) Get2(ctx context.Context, n *pb.GetParams) (*pb.Number, error) {
	s.i2 += 2
	if math.Mod(float64(s.i2), 1000.0) == 0.0 {
	log.Printf("grpc i2: %v", s.i2)
	}
	return &pb.Number{Value: 2}, nil
}

func main() {
	i1 := 0
	i2 := 0
	http.HandleFunc("/get1", func(w http.ResponseWriter, r *http.Request) {
		i1 += 1
		if math.Mod(float64(i1), 1000.0) == 0.0 {
			log.Printf("rest i1: %v", i1)
		}

		io.WriteString(w, "1")
	})
	http.HandleFunc("/get2", func(w http.ResponseWriter, r *http.Request) {
		i2 += 2
		if math.Mod(float64(i2), 1000.0) == 0.0 {
			log.Printf("rest i2: %v", i2)
		}
		io.WriteString(w, "2")
	})
	go func() { log.Fatal(http.ListenAndServe(":8001", nil)) }()
	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGopageServer(grpcServer, &gopageServer{})
	grpcServer.Serve(lis)

}
