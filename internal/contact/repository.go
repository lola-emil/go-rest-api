package contact

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ContactRepo struct {
	db *sqlx.DB
}

func NewContactRepo(db *sqlx.DB) *ContactRepo {
	return &ContactRepo{
		db: db,
	}
}

func (r *ContactRepo) Save(ctx context.Context, contact Contact) (int64, error) {
	query := `
		INSERT INTO contacts (name, email, phone_number, user_id)
		VALUES (?, ?, ?, ?)
	`
	result, err := r.db.ExecContext(ctx, query,
		contact.Name,
		contact.Email,
		contact.PhoneNumber,
		contact.UserId,
	)

	if err != nil {
		return 0, fmt.Errorf("Error saving contact: %w", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("get inserted contact id: %w", err)
	}

	return id, nil
}

func (r *ContactRepo) FindAll(offset int, limit int) ([]Contact, error) {
	query := "SELECT * FROM contacts WHERE 1 LIMIT ? OFFSET ?"

	contacts := []Contact{}

	if err := r.db.Select(&contacts, query, limit, offset); err != nil {
		return []Contact{}, fmt.Errorf("Error selecting contacts: %w", err)
	}

	return contacts, nil
}

func (r *ContactRepo) FindById(id int64) (Contact, error) {
	query := "SELECT * FROM contacts WHERE id = ?"

	var contact Contact

	if err := r.db.Get(&contact, query, id); err != nil {
		return Contact{}, fmt.Errorf("Error getting contact by ID: %w", err)
	}

	return contact, nil
}

func (r *ContactRepo) FindByUserId(userId int64) ([]Contact, error) {
	query := "SELECT * FROM contacts WHERE user_id = ?"

	contacts := []Contact{}

	if err := r.db.Select(&contacts, query, userId); err != nil {
		return []Contact{}, fmt.Errorf("Error gett contacts by user ID: %w", err)
	}

	return contacts, nil
}

func (r *ContactRepo) DeleteById(ctx context.Context, id int64) error {
	query := "DELETE FROM contacts WHERE id = ?"

	_, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return fmt.Errorf("Error deleting contact: %w", err)
	}

	return nil
}
