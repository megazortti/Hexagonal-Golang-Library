// /internal/infra/persistence/account_repo_impl.go
package persistence

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"hexagonal-bank/internal/domain/entities"
	"io/ioutil"
	"net/http"
)

type AccountRepositoryJSONServer struct {
	BaseURL string
}

func (r *AccountRepositoryJSONServer) FindByID(id string) (*entities.Account, error) {
	resp, err := http.Get(fmt.Sprintf("%s/accounts/%s", r.BaseURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("account not found")
	}

	var account entities.Account
	if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
		return nil, err
	}
	fmt.Println("%v", account)
	return &account, nil
}

func (r *AccountRepositoryJSONServer) Save(account *entities.Account) error {
	// accountJSON, err := json.Marshal(account)
	// if err != nil {
	// return err
	// }
	temp := struct {
		Balance float64 `json:"amount"`
	}{
		Balance: account.Balance,
	}
	dataToSave, err := json.Marshal(temp)

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/accounts/%s", r.BaseURL, account.ID), bytes.NewBuffer(dataToSave))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(body))
	}

	return nil
}
