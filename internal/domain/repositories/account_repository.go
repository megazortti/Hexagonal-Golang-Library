package repositories

import "hexagonal-bank/internal/domain/entities"

type AccountRepository interface {
	FindByID(id string) (*entities.Account, error)
	Save(account *entities.Account) error
}
