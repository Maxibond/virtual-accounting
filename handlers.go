package main

import (
	"encoding/json"
	"net/http"
)

func HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var b RequestCreateAccount
	err := decoder.Decode(&b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var res []byte
	var jsonErr error
	accountID, err := app.CreateAccount(b.IdempotencyKey, b.InitialBalance)
	if err != nil {
		res, jsonErr = json.Marshal(ResponseError{err.Error()})
	} else {
		res, jsonErr = json.Marshal(ResponseCreateAccount{
			IdempotencyKey: b.IdempotencyKey,
			AccountID:      accountID,
		})
	}

	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}

func HandleGetAccountBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var b RequestGetBalance
	err := decoder.Decode(&b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	balance := app.GetBalance(b.AccountID)

	res, err := json.Marshal(ResponseGetBalance{
		Balance: balance,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}

func HandleCreateMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var b RequestCreateMove
	err := decoder.Decode(&b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var res []byte
	var jsonErr error
	moveID, err := app.CreateMove(b.IdempotencyKey, b.FromID, b.ToID, b.Amount)
	if err != nil {
		res, jsonErr = json.Marshal(ResponseError{err.Error()})
	} else {
		res, jsonErr = json.Marshal(ResponseCreateMove{
			IdempotencyKey: b.IdempotencyKey,
			MoveID:         moveID,
		})
	}

	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}
