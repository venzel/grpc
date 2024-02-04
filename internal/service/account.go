package service

import (
	"context"
	"grpc/internal/database"
	"grpc/internal/pb"
	"io"
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

func (a *AccountService) CreateAccountStream(stream pb.AccountService_CreateAccountStreamServer) error {
	accounts := &pb.AccountList{}

	for {
		account, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(accounts)
		}

		if err != nil {
			return err
		}

		accountResult, err := a.AccountDB.Create(account.Name, account.Email)

		if err != nil {
			return err
		}

		accounts.Accounts = append(accounts.Accounts, &pb.Account{
			Id:    accountResult.ID,
			Name:  accountResult.Name,
			Email: accountResult.Email,
		})
	}
}
