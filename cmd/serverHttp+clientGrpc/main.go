package main

import (
	"context"
	"fmt"
	"hello/grpc/pb"
	"log"
	"net/http"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Obtenha o caminho da URL sem a barra inicial
		path := r.URL.Path[1:]

		// Verifique se o caminho é um número (o ID)
		if id := path; id != "" {
			startTime := time.Now()

			req := &pb.HelloRequest{
				Name: id,
			}

			res, err := client.SayHello(context.Background(), req)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(w, "%s", res.GetMessage())
			log.Print(res)

			endTime := time.Now()
			elapsedTime := endTime.Sub(startTime)
			log.Print(elapsedTime)
		} else {
			fmt.Fprint(w, "Rota inválida.")
		}
	})

	http.ListenAndServe(":8080", nil)
}
