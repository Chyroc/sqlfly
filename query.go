package sqlfly

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // 注册mysql driver
)

type Generator func() (func(), []interface{})

type queryRows struct {
	Rows *sql.Rows
	Rrr  error
}

func (r queryRows) Each(g Generator) error {
	if r.Rrr != nil {
		return r.Rrr
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
	return queryRows{Rows: rows, Rrr: err}
}

func Query(dber *sql.DB, query string, args ...interface{}) queryRows {
	return QueryContext(context.Background(), dber, query, args...)
}
