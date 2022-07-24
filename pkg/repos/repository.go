package repos

import "database/sql"

type Repository struct {
	Db *sql.DB
}

func NewRepository() (*Repository, error) {
	db, err := Connect()
	if err != nil {
		return &Repository{}, err
	}

	return &Repository{Db: db}, nil
}
