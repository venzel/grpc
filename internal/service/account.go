package service

import (
	"context"
	"grpc/internal/database"
	"grpc/internal/pb"
)

type AccountService struct {
	pb.UnimplementedAccountServiceServer
	AccountDB database.Account
}

func NewAccountService(accountDB database.Account) *AccountService {
	return &AccountService{AccountDB: accountDB}
}

func (a *AccountService) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Account, error) {
	account, err := a.AccountDB.Create(in.Name, in.Email)

	if err != nil {
		return nil, err
	}

	return &pb.Account{
		Id:    account.ID,
		Name:  account.Name,
		Email: account.Email,
	}, err
}
