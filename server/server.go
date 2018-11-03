package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rootemanuel/rpc-api/pb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Teste(ctx context.Context, req *testepb.TesteRequest) (*testepb.TesteResponse, error) {

	fmt.Printf("call Teste / req => %v", req)

	nome := req.GetTeste().GetNome()
	sobrenome := req.GetTeste().GetSobrenome()

	completo := nome + sobrenome

	res := &testepb.TesteResponse{
		Resultado: completo,
	}

	return res, nil
}

func main() {
	fmt.Println("root server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	testepb.RegisterTesteServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to server: %v", err)
	}

}
