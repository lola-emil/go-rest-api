package product

import "github.com/jmoiron/sqlx"

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]Product, error) {
	products := []Product{}

	err := s.db.Select(&products, "SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	return products, nil
}
