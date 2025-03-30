package grpcservice

import (
	"context"
	"fmt"

	"go.uber.org/fx"

	"auth/ent"
)

var Module = fx.Provide(
	NewAuthGrpcService,
)

func rollback(tx *ent.Tx, err error) error {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		err = fmt.Errorf("%w: error when rolling back. %v", err, rollbackErr)
	}

	return err
}

func commit(tx *ent.Tx) error {
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func withTransaction(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	var err error

	tx, err := client.Tx(ctx)

	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if v := recover(); v != nil {
			if rollbackErr := rollback(tx, fmt.Errorf("panic: %v", v)); rollbackErr != nil {
				err = rollbackErr
			}
		}
	}()

	if execErr := fn(tx); execErr != nil {
		err = rollback(tx, execErr)
	} else {
		err = commit(tx)
	}

	return err
}
