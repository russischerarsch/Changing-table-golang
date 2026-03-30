package databaseconnection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) *pgx.Conn {
	return pgx.Connect(ctx, "postgres://localhost:dkfl26052010@postgres:5432/postgres")
}
