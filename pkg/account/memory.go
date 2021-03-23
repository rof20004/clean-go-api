package account

import "github.com/google/uuid"

type Memory struct {
	repo map[uuid.UUID]*Account
}

func NewMemory(repo map[uuid.UUID]*Account) *Memory {
	return &Memory{repo}
}

func (m *Memory) Create(account *Account) error {
	account.ID = uuid.New()
	m.repo[account.ID] = account
	return nil
}

func (m *Memory) FindAll() ([]*Account, error) {
	accounts := make([]*Account, 0)

	for _, v := range m.repo {
		accounts = append(accounts, v)
	}

	return accounts, nil
}
