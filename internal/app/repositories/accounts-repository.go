package repositories

import "github.com/henrique998/go-auth-2/internal/app/entities"

type AccountsRepository interface {
	FindById(accountId string) *entities.Account
	FindByEmail(email string) *entities.Account
	Create(a entities.Account) error
	Update(a entities.Account) error
}
