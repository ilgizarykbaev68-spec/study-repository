package simpleconnection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("CONN_STRING")
	return pgx.Connect(ctx, connString)
	// return pgx.Connect(ctx, "postgres://postgres:salamandra@localhost:5432/postgres")
}
