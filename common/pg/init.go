package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"synapsis-challange/common/sqlc"
)

func New(dbURL string) (*pgxpool.Pool, sqlc.Querier) {
	pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}
	return pool, sqlc.New(pool)
}
