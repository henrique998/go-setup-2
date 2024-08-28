package createaccountusecase

import (
	"fmt"

	"github.com/henrique998/go-auth-2/internal/app/errors"
	"github.com/henrique998/go-auth-2/internal/app/requests"
)

func (uc *createAccountService) Execute(req requests.CreateAccountRequest) errors.AppErr {
	fmt.Println(req)

	return nil
}
