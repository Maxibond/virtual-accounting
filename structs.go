package main

import "sync"

type Account struct {
	IdempotencyKey string `json:"idempotency_key"`
	ID             int64  `json:"id"`
	Balance        int64  `json:"balance"`
}

type Move struct {
	IdempotencyKey string `json:"idempotency_key"`
	ID             int64  `json:"id"`
	FromID         int64  `json:"from_id"`
	ToID           int64  `json:"to_id"`
	Amount         int64  `json:"amount"`
}

type AccountRepository struct {
	Accounts   []Account
	IDSequence Sequence
}

type MoveRepository struct {
	Moves      []Move
	IDSequence Sequence
}

type Storage struct {
	AccountRepository AccountRepository
	MoveRepository    MoveRepository
}

type Application struct {
	storage Storage
}

type Sequence struct {
	mu    *sync.Mutex
	value int64
}

func (s *Sequence) Inc() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.value += 1
	return s.value
}

func (s *Sequence) Dec() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.value -= 1
	return s.value
}

func NewSequence() Sequence {
	return Sequence{
		mu:    &sync.Mutex{},
		value: 0,
	}
}

// http structs begin

type ResponseError struct {
	Error string `json:"error"`
}

type RequestCreateAccount struct {
	IdempotencyKey string `json:"idempotency_key"`
	InitialBalance int64  `json:"initial_balance"`
}

type ResponseCreateAccount struct {
	IdempotencyKey string `json:"idempotency_key"`
	AccountID      int64  `json:"account_id"`
}

type RequestCreateMove struct {
	IdempotencyKey string `json:"idempotency_key"`
	FromID         int64  `json:"from_id"`
	ToID           int64  `json:"to_id"`
	Amount         int64  `json:"amount"`
}

type ResponseCreateMove struct {
	IdempotencyKey string `json:"idempotency_key"`
	MoveID         int64  `json:"move_id"`
}

type RequestGetBalance struct {
	AccountID int64 `json:"account_id"`
}

type ResponseGetBalance struct {
	Balance int64 `json:"balance"`
}

// http structs end
