package user

import (
	"context"
	"database/sql"
	"handarudwiki/mini-online-shop-go/infra/response"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateUser(ctx context.Context, user UserEntity) (err error) {
	query := `
		INSERT INTO users (public_id,email, password, role, created_at, updated_at) 
		VALUES (:public_id,:email, :password, :role, :created_at, :updated_at)
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user)
	return
}

func (r repository) GetUserByEmail(ctx context.Context, email string) (user UserEntity, err error) {
	query := `SELECT id, email, public_id, password, role, created_at, updated_at
				 FROM users WHERE email = $1`
	err = r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}
	return
}
