package testdata

import (
	"testing"

	"github.com/Chyroc/sqlfly"
)

func TestQueryRow(t *testing.T) {
	t.SkipNow()

	t.Run("", func(t *testing.T) {
		var id string

		if err := sqlfly.QueryRow(nil, "select id from t limit 1", nil).Scan(&id); err != nil {
			t.Fail()
		}
	})

	t.Run("", func(t *testing.T) {
		var ids []string
		err := sqlfly.Query(nil, "select id from t limit 1", nil).Each(func() (func(), []interface{}) {
			var id string
			return func() { ids = append(ids, id) }, []interface{}{&id}
		})

		if err != nil {
			t.Fail()
		}
	})
}
