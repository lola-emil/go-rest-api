package contact

type Contact struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Email       string `db:"email" json:"email"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	UserId      int64  `db:"user_id" json:"user_id"`
}
