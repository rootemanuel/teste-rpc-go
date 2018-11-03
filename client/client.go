package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rootemanuel/rpc-api/pb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("root client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client error connect : %v", err)
	}

	defer cc.Close()

	c := testepb.NewTesteServiceClient(cc)
	for i := 0; i < 100; i++ {
		rootUnary(c)
	}
}

func rootUnary(c testepb.TesteServiceClient) {
	req := &testepb.TesteRequest{
		Teste: &testepb.Teste{
			Nome:      "root",
			Sobrenome: "grpc",
		},
	}

	res, err := c.Teste(context.Background(), req)
	if err != nil {
		log.Fatalf("error call Teste RPC: %v", err)
	}

	log.Printf("response Teste, => %v", res)
}
