package service

import (
	"errors"
	"fmt"
)

// go install go.uber.org/mock/mockgen@latest

//go:generate mockgen -source $GOFILE -destination ./mock.go -package ${GOPACKAGE}

// go generate ./...

type PaymentGateway interface {
	SendMoneyAndGetCurrentBalance(username string, amount int) (int, error)
}

type Logger interface {
	Error(message string)
	Info(message string)
}

var ErrConnection = errors.New("connection error")

func TransferMoney(e PaymentGateway, l Logger, username string, amount int) (int, error) {
	balance, err := e.SendMoneyAndGetCurrentBalance(username, amount)
	if err != nil {
		l.Error("send money error")
		return 0, fmt.Errorf("send money: %w", err)
	}

	l.Info("send money ok")
	return balance, nil
}
