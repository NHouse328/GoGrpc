package main

import (
	"context"
	"fmt"
	"hello/grpc/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewHelloClient(conn)

	startTime := time.Now()

	for i := 1; i <= 5; i++ {
		log.Println(i)

		name := fmt.Sprintf("Gabriel %d", i)

		req := &pb.HelloRequest{
			Name: name,
		}

		res, err := client.SayHello(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		log.Print(res)
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	log.Print(elapsedTime)
}
