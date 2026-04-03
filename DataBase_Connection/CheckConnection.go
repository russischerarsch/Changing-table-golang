package databaseconnection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

var Ctx = context.Background()
var Conn *pgx.Conn

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {
	connStr := os.Getenv("CONN_STRING")
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
