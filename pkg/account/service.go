package account

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) Register(account *Account) error {
	if err := s.repo.Create(account); err != nil {
		return err
	}

	return nil
}

func (s *Service) ListAccounts() ([]*Account, error) {
	accounts, err := s.repo.FindAll()
	if err != nil {
		return accounts, err
	}

	if len(accounts) == 0 {
		return accounts, ErrNoAccountsFound
	}

	return accounts, nil
}
