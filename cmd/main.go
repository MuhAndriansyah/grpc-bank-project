package main

import (
	"log"

	"github.com/MuhAndriansyah/grpc-bank-project/internal/adapter/database"
	"github.com/MuhAndriansyah/grpc-bank-project/internal/adapter/grpc"
	"github.com/MuhAndriansyah/grpc-bank-project/internal/application"
)

func main() {

	dsn := "postgres://andre-dev:bankgrpcrahasia@localhost:5432/bank_grpc_db?sslmode=disable"

	dbAdapter, err := database.NewDatabaseAdapter(dsn)

	if err != nil {
		log.Fatalf("failed to connect database :%v", err)
	}

	bs := application.NewBankService(dbAdapter)

	grpcAdapter := grpc.NewGrpcAdapter(bs, 9090)

	grpcAdapter.Run()
}
