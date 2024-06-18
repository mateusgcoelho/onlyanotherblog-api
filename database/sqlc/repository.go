package database

import (
	"context"
	"fmt"
	"onlyanotherblog/config"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type DatabaseRepository struct {
	*Queries
	Db *pgxpool.Pool
}

func NewDatabaseRepository(serverConfig config.ServerConfig) (*DatabaseRepository, error) {
	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		serverConfig.DatabaseUsername, serverConfig.DatabasePassword, serverConfig.DatabaseHost, serverConfig.DatabasePort, serverConfig.DatabaseName,
	)
	conn, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}

	return &DatabaseRepository{
		Queries: New(conn),
		Db:      conn,
	}, nil
}

func (r *DatabaseRepository) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := r.Db.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit(ctx)
}
