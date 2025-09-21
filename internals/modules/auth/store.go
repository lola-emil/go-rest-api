package auth

import (
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*User, error) {
	user := User{}

	err := s.db.Get(&user, "SELECT * FROM user where email = ?", email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
