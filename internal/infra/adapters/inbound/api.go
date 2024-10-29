// /internal/infra/adapters/inbound/api.go
package inbound

import (
	"encoding/json"
	"hexagonal-bank/internal/app/services"
	"net/http"
)

type API struct {
	TransferMoneyUseCase *services.TransferMoney
}

func (api *API) Transfer(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FromID string  `json:"from_id"`
		ToID   string  `json:"to_id"`
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := api.TransferMoneyUseCase.Execute(req.FromID, req.ToID, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
