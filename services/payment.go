package services

import (
	"blu/repository/account"
	"blu/repository/payment"
	"errors"
	"math/big"
	"strings"
)

func Transfer(from string, to string, balance *big.Float) error {
	if strings.TrimSpace(from) == "" || strings.TrimSpace(to) == "" {
		return errors.New("[ERROR] account name cannot be empty")
	}
	if balance == nil || balance.Cmp(big.NewFloat(0)) < 0 {
		return errors.New("[ERROR] balance must be a non-negative value")
	}

	fromName, fromBalance, err := account.FindByName(from)
	if err != nil {
		return err
	}

	toName, toBalance, err := account.FindByName(to)
	if err != nil {
		return err
	}

	if err := account.Save(fromName, fromBalance.Sub(fromBalance, balance)); err != nil {
		return err
	}

	if err := account.Save(toName, toBalance.Add(toBalance, balance)); err != nil {
		return err
	}

	return payment.Save(from, to, balance)
}
