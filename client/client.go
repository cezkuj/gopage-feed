package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"

	pb "github.com/cezkuj/gopage-feed/grpc"
)
func getRest(path string){
for {
                        res, err := http.Get("http://localhost:8001/" + path)
                        if err != nil {
                                log.Println(err)
                        }
                        res.Body.Close()
                }


}

func main() {
        go getRest("get1")
        go getRest("get1")
        go getRest("get2")
	conn, err := grpc.Dial("localhost:8002", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	client := pb.NewGopageClient(conn)
	go func() {
		for {
			_, err = client.Get1(context.Background(), &pb.GetParams{})
			if err != nil {
				log.Println(err)
				return
			}
		}
	}()

	for {
		_, err := client.Get2(context.Background(), &pb.GetParams{})
		if err != nil {
			log.Println(err)
			return
		}
	}
}
