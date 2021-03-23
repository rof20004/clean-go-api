package account

import "github.com/google/uuid"

type Account struct {
	ID      uuid.UUID `json:"id"`
	Owner   string    `json:"owner"`
	Address string    `json:"address"`
}
