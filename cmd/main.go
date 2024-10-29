// /cmd/main.go
package main

import (
	"hexagonal-bank/internal/app/services"
	"hexagonal-bank/internal/infra/adapters/inbound"
	persistence "hexagonal-bank/internal/infra/adapters/outbound"
	"net/http"
)

func main() {

	accountRepo := &persistence.AccountRepositoryJSONServer{BaseURL: "http://localhost:3001"}
	transactionRepo := &persistence.TransactionRepositoryJSONServer{
		BaseURL: "http://localhost:3001",
	}

	transferMoney := &services.TransferMoney{
		AccountRepo:     accountRepo,
		TransactionRepo: transactionRepo,
	}

	api := &inbound.API{
		TransferMoneyUseCase: transferMoney,
	}

	http.HandleFunc("/transfer", api.Transfer)
	http.ListenAndServe(":8080", nil)
}
