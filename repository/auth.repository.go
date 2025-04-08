package repository

import (
	"auth/ent"
	"auth/ent/auth"
	"context"
)
import db "database/sql"

type authRepository struct {
	client *ent.Client
	db     *db.DB
}

func (a authRepository) UpdateAccountState(ctx context.Context, tx *ent.Tx, username string, state int32) error {
	return tx.Auth.Update().Where(auth.UsernameEQ(username)).SetState(state).Exec(ctx)
}

func (a authRepository) FindByUsername(ctx context.Context, username string) (*ent.Auth, error) {
	return a.client.Auth.Query().Where(auth.UsernameEQ(username)).WithEmployee(
		func(query *ent.EmployeeQuery) {
			query.WithDoctor().WithNurse().WithDepartment()
		},
	).First(ctx)
}

func (a authRepository) UpdatePassword(ctx context.Context, tx *ent.Tx, username string, password string) error {
	return tx.Auth.Update().Where(auth.UsernameEQ(username)).SetPassword(password).Exec(ctx)
}

func (a authRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return a.client.Auth.Query().Where(auth.EmailEQ(email)).Exist(ctx)
}

func NewAuthRepository(client *ent.Client, db *db.DB) AuthRepository {
	return &authRepository{client: client, db: db}
}
