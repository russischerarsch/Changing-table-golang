package databaseconnection

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	SQLquery := `
	CREATE TABLE IF NOT EXISTS employees (
	id SERIAL PRIMARY KEY,
	fullname VARCHAR(100) NOT NULL,
	position VARCHAR(30) NOT NULL,
	created_at TIMESTAMP NOT NULL
	)	
	`
	_, err := conn.Exec(ctx, SQLquery)
	if err != nil {
		return err
	}
	return nil
}

func AddEmployeeDB(
	conn *pgx.Conn,
	ctx context.Context,
	ID int,
	FullName string,
	Position string,
	createdAt time.Time,
) error {
	SQLquery := ` 
	INSERT INTO employees (id, fullname, position, created_at)
	VALUES($1, $2, $3, $4)
	`
	_, err := conn.Exec(ctx, SQLquery, ID, FullName, Position, createdAt)
	return err
}

func DeleteEmployeeDB(
	conn *pgx.Conn,
	ctx context.Context,
	ID int,
) error {
	SQLquery := `
	DELETE FROM employees
	WHERE id = $1
	`
	_, err := conn.Exec(ctx, SQLquery, ID)
	return err
}
