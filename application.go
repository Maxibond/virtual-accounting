package main

import "errors"

func NewApplication() Application {
	return Application{
		storage: Storage{
			AccountRepository: AccountRepository{
				Accounts:   make([]Account, 0),
				IDSequence: NewSequence(),
			},
			MoveRepository: MoveRepository{
				Moves:      make([]Move, 0),
				IDSequence: NewSequence(),
			},
		},
	}
}

func (a *Application) CreateAccount(idempotencyKey string, initial int64) (int64, error) {
	for _, account := range a.storage.AccountRepository.Accounts {
		if account.IdempotencyKey == idempotencyKey {
			return -1, errors.New("idempotency key is used")
		}
	}

	newAccount := Account{
		ID:             a.storage.AccountRepository.IDSequence.Inc(),
		IdempotencyKey: idempotencyKey,
	}

	_, err := a.CreateMove("", 0, newAccount.ID, initial)
	if err != nil {
		a.storage.AccountRepository.IDSequence.Dec()
		return -1, err
	}

	a.storage.AccountRepository.Accounts = append(a.storage.AccountRepository.Accounts, newAccount)

	return newAccount.ID, nil
}

func (a *Application) CreateMove(idempotencyKey string, fromID, toID int64, amount int64) (int64, error) {
	for _, move := range a.storage.MoveRepository.Moves {
		if move.IdempotencyKey == idempotencyKey && idempotencyKey != "" {
			return -1, errors.New("idempotency key is used")
		}
	}

	var fromExists, toExists bool
	for _, account := range a.storage.AccountRepository.Accounts {
		if account.ID == fromID {
			fromExists = true
		}
		if account.ID == toID {
			toExists = true
		}
	}
	if !(fromExists && toExists) && idempotencyKey != "" {
		return -1, errors.New("accounts are not exist")
	}

	balance := a.GetBalance(fromID)
	if balance < amount && idempotencyKey != "" {
		return -1, errors.New("not enough balance")
	}

	newMove := Move{
		ID:             a.storage.MoveRepository.IDSequence.Inc(),
		IdempotencyKey: idempotencyKey,
		FromID:         fromID,
		ToID:           toID,
		Amount:         amount,
	}

	a.storage.MoveRepository.Moves = append(a.storage.MoveRepository.Moves, newMove)

	return newMove.ID, nil
}

func (a *Application) GetBalance(accountID int64) int64 {
	var input, output int64
	for _, move := range a.storage.MoveRepository.Moves {
		if move.FromID == accountID {
			output += move.Amount
		}
		if move.ToID == accountID {
			input += move.Amount
		}
	}

	return input - output
}
