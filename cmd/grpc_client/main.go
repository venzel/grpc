package main

import (
	"context"
	"grpc/internal/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Não foi possível conectar: %v", err)
	}
	defer conn.Close()

	log.Println("Conexão com o servidor gRPC estabelecida")

	client := pb.NewAccountServiceClient(conn)

	/* Cria uma conta */

	createAccount, err := client.CreateAccount(context.Background(), &pb.CreateAccountRequest{
		Name:  "daniel",
		Email: "daniel@gmail.com",
	})
	if err != nil {
		log.Fatalf("Não foi possível criar a conta: %v", err)
	}
	log.Printf("Conta criada: %v", createAccount)

	/* Cria contas com utilizando stream */

	stream, err := client.CreateAccountStream(context.Background())
	if err != nil {
		log.Fatalf("Não foi possível criar a conta: %v", err)
	}
	accounts := []*pb.CreateAccountRequest{
		{
			Name:  "daniel",
			Email: "daniel@gmail.com",
		},
		{
			Name:  "otávio",
			Email: "otavio@gmail.com",
		},
	}
	for _, account := range accounts {
		if err := stream.Send(account); err != nil {
			log.Fatalf("Não foi possível enviar a conta: %v", err)
		}
		log.Printf("Conta criada: %v", account)
	}
	_, err = stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Não foi possível fechar o stream: %v", err)
	}

	/* Lista as contas */

	listAccounts, err := client.ListAccounts(context.Background(), &pb.Blank{})
	if err != nil {
		log.Fatalf("Não foi possível listar as contas: %v", err)
	}
	for _, account := range listAccounts.Accounts {
		log.Printf("Conta: %v", account)
	}
}
