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

type Generator func() (func(), []interface{})

type queryRows struct {
	*sql.Rows
	err error
}

func (r queryRows) Each(g Generator) error {
	if r.err != nil {
		return r.err
	}
	defer r.Rows.Close()

	for r.Rows.Next() {
		f, dest := g()
		if err := r.Rows.Scan(dest...); err != nil {
			return err
		}

		f() // append
	}

	return r.Rows.Err()
}

func QueryContext(ctx context.Context, dber *sql.DB, query string, args ...interface{}) queryRows {
	rows, err := dber.QueryContext(ctx, query, args...)
	return queryRows{Rows: rows, err: err}
}

func Query(dber *sql.DB, query string, args ...interface{}) queryRows {
	return QueryContext(context.Background(), dber, query, args)
}
