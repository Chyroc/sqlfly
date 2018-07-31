package sqlfly

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // 注册mysql driver
)

type queryRow struct {
	row *sql.Row
}

func (r queryRow) Scan(dest ...interface{}) error {
	return r.row.Scan(dest...)
}

func QueryRowContext(ctx context.Context, dber *sql.DB, query string, args ...interface{}) queryRow {
	return queryRow{row: dber.QueryRowContext(ctx, query, args...)}
}

func QueryRow(dber *sql.DB, query string, args ...interface{}) queryRow {
	return QueryRowContext(context.Background(), dber, query, args...)
}
