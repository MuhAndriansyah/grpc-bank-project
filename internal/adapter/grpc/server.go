package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/MuhAndriansyah/grpc-bank-project/internal/port"
	bankv1 "github.com/MuhAndriansyah/grpc-bank-project/proto/bank/v1"
	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	grpcPort    int
	bankService port.BankServicePort
	server      *grpc.Server
	bankv1.BankServiceServer
}

func NewGrpcAdapter(bankService port.BankServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		bankService: bankService,
		grpcPort:    grpcPort,
	}
}

func (a *GrpcAdapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen port %d: %v", a.grpcPort, err)
	}

	grpcServer := grpc.NewServer()
	a.server = grpcServer

	bankv1.RegisterBankServiceServer(grpcServer, a)
	log.Println("server starting on port ", a.grpcPort)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to server gRPC on port %d :%v\n", a.grpcPort, err)
	}
}

func (a *GrpcAdapter) Stop() {
	a.server.Stop()
}
