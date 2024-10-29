package repositories

import "hexagonal-bank/internal/domain/entities"

type TransactionRepository interface {
	Save(transaction *entities.Transaction) error
}
