package service

import (
	"context"
	"grpc/internal/db"
	"grpc/internal/pb"
)

type AccountService struct {
	pb.UnimplementedAccountServiceServer
	AccountDB db.Account
}

func NewAccountService(accountDB db.Account) *AccountService {
	return &AccountService{AccountDB: accountDB}
}

func (a *AccountService) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.AccountResponse, error) {
	account, err := a.AccountDB.Create(in.Name, in.Email)

	if err != nil {
		return nil, err
	}

	newAccount := &pb.Account{
		Id:    account.ID,
		Name:  account.Name,
		Email: account.Email,
	}

	return &pb.AccountResponse{
		Account: newAccount,
	}, nil
}
