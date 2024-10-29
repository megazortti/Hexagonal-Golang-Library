package services

import (
	"errors"
	"hexagonal-bank/internal/domain/entities"
	"hexagonal-bank/internal/domain/repositories"
)

type TransferMoney struct {
	AccountRepo     repositories.AccountRepository
	TransactionRepo repositories.TransactionRepository
}

func (u *TransferMoney) Execute(fromID string, toID string, amount float64) error {
	fromAccount, err := u.AccountRepo.FindByID(fromID)
	if err != nil {
		return err
	}
	toAccount, err := u.AccountRepo.FindByID(toID)
	if err != nil {
		return err
	}

	if fromAccount.Balance < amount {
		return errors.New("saldo insuficiente")
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount

	if err := u.AccountRepo.Save(fromAccount); err != nil {
		return err
	}
	if err := u.AccountRepo.Save(toAccount); err != nil {
		return err
	}

	transaction := &entities.Transaction{
		FromID: fromID,
		ToID:   toID,
		Amount: amount,
	}

	return u.TransactionRepo.Save(transaction)
}
