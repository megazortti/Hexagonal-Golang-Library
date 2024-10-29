// /internal/infra/persistence/transaction_repo_impl.go
package persistence

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hexagonal-bank/internal/domain/entities"
	"net/http"
)

type TransactionRepositoryJSONServer struct {
	BaseURL string
}

func (r *TransactionRepositoryJSONServer) Save(transaction *entities.Transaction) error {
	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf("%s/transactions", r.BaseURL), "application/json", bytes.NewBuffer(transactionJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(transaction)
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create transaction, status: %d", resp.StatusCode)
	}

	return nil
}
