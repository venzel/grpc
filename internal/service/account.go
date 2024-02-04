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

func (a *AccountService) ListAccounts(ctx context.Context, in *pb.Blank) (*pb.AccountList, error) {
	accounts, err := a.AccountDB.FindAll()

	if err != nil {
		return nil, err
	}

	var accountsResponse []*pb.Account

	for _, account := range accounts {
		accountsResponse = append(accountsResponse, &pb.Account{
			Id:    account.ID,
			Name:  account.Name,
			Email: account.Email,
		})
	}

	return &pb.AccountList{Accounts: accountsResponse}, nil
}

func (a *AccountService) GetAccount(ctx context.Context, in *pb.AccountGetRequest) (*pb.Account, error) {
	account, err := a.AccountDB.FindOne(in.Id)

	if err != nil {
		return nil, err
	}

	return &pb.Account{
		Id:    account.ID,
		Name:  account.Name,
		Email: account.Email,
	}, nil
}
