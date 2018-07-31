package sqlfly

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	as := assert.New(t)

	t.Run("query row", func(t *testing.T) {
		var id string
		as.Equal(sql.ErrNoRows, QueryRow(testDB, "select id from test limit 1").Scan(&id))
	})

	t.Run("query rows", func(t *testing.T) {
		var ids []string
		as.Nil(Query(testDB, "select id from test").Each(func() (func(), []interface{}) {
			var id string
			return func() { ids = append(ids, id) }, []interface{}{&id}
		}))
		as.Len(ids, 0)
	})
}
