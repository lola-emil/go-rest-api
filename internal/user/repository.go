package user

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, user UserModel) (int64, error) {
	query := `
		INSERT INTO users (firstname, lastname, email, password)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query,
		user.Firstname,
		user.Lastname,
		user.Email,
		user.Password,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) FindAll(offset int, limit int) ([]UserModel, error) {

	query := "SELECT * FROM users WHERE 1 LIMIT ? OFFSET ?"
	users := []UserModel{}

	if err := r.db.Select(&users, query, limit, offset); err != nil {
		return []UserModel{}, err
	}

	return users, nil
}

func (r *UserRepository) ForEachUser(
	ctx context.Context,
	fn func(UserModel) error,
) error {
	query := "SELECT * FROM users"
	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var u UserModel
		if err := rows.StructScan(&u); err != nil {
			return err
		}
		if err := fn(u); err != nil {
			return err
		}
	}

	return rows.Err()
}

func (r *UserRepository) FindById(id int64) (UserModel, error) {
	user := UserModel{}

	if err := r.db.Get(&user, "SELECT * FROM users WHERE id = ?", id); err != nil {
		return UserModel{}, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (UserModel, error) {
	query := "SELECT * FROM users WHERE email = ?"

	user := UserModel{}
	if err := r.db.Get(&user, query, email); err != nil {
		return UserModel{}, err
	}

	return user, nil
}

func (r *UserRepository) DeleteById(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
