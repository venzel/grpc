package main

import (
	"database/sql"
	"grpc/internal/database"
	"grpc/internal/pb"
	"grpc/internal/service"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	accountDb := database.NewAccount(db)
	accountService := service.NewAccountService(*accountDb)

	grpcServer := grpc.NewServer()

	// Atacha o serviço ao servidor gRPC
	pb.RegisterAccountServiceServer(grpcServer, accountService)

	// Reflection ler e processa sua própria informação, necessário para o Evans
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
