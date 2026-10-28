package services

import (
	"blu/repository/deposit"
	"fmt"
	"math/big"
	"time"
)

var interestRate = big.NewFloat(0.01)

func AddDeposit(name string, amount *big.Float) error {
	if amount == nil || amount.Cmp(big.NewFloat(0)) < 0 {
		return fmt.Errorf("invalid amount format")
	}

	return deposit.Save(name, amount)
}

func ShowDeposits() error {
	deposits, err := deposit.FindAll()
	if err != nil {
		return err
	}

	fmt.Println("Current deposits:")
	for i, d := range deposits {
		name := d["name"].(string)
		amount := d["amount"].(*big.Float)
		created := d["created_at"].(time.Time)
		fmt.Printf("%d. %s: %s (created %s)\n", i+1, name, amount.Text('f', 2), created.Format(time.RFC822))
	}

	return nil
}

func AccrueInterest() error {
	deposits, err := deposit.FindAll()
	if err != nil {
		return err
	}

	now := time.Now()

	for _, d := range deposits {
		amount := d["amount"].(*big.Float)

		createdAt, ok := d["created_at"].(time.Time)
		if !ok {
			str := d["created_at"].(string)
			parsed, err := time.Parse(time.RFC822, str)
			if err != nil {
				continue
			}
			createdAt = parsed
			d["created_at"] = createdAt
		}

		minutesPassed := int(now.Sub(createdAt).Minutes())
		if minutesPassed <= 0 {
			continue
		}

		rate := new(big.Float).Mul(interestRate, big.NewFloat(float64(minutesPassed)))
		interest := new(big.Float).Mul(amount, rate)
		amount.Add(amount, interest)
		d["amount"] = amount
		d["created_at"] = now
	}

	return deposit.SaveAll(deposits)
}
