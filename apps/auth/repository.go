package auth

import (
	"context"
	"database/sql"

	"github.com/aripkur/go-learn-shop/infra/response"
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

func (r repository) CreateAuth(ctx context.Context, model AuthEntity) (err error) {
	query := `
		INSERT INTO auth (
			public_id, email, password, role, created_at, updated_at
		) VALUES (
			:public_id, :email, :password, :role, :created_at, :updated_at
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, model)

	return
}

func (r repository) GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error) {
	query := `
		SELECT id, public_id, email, password, role, created_at, updated_at
		FROM auth
		WHERE email=?
	`

	err = r.db.GetContext(ctx, &model, query, email)

	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
		}
		return
	}

	return
}
