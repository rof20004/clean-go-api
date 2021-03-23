package account

type Repository interface {
	FindAll() ([]*Account, error)
	Create(*Account) error
}
