package services

import (
	"blu/repository/account"
	"errors"
	"math/big"
	"strings"
)

func CreateAccount(name string, balance *big.Float) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("[ERROR] account name cannot be empty")
	}
	if balance == nil || balance.Cmp(big.NewFloat(0)) < 0 {
		return errors.New("[ERROR] balance must be a non-negative value")
	}
	return account.Save(name, balance)
}

func RemoveAccount(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("[ERROR] account name cannot be empty")
	}
	return account.DeleteByName(name)
}
