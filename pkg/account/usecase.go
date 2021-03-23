package account

type Usecase interface {
	Register(*Account) error
	ListAccounts() ([]Account, error)
}
