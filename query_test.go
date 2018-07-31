package sqlfly

import (
	"testing"
)

func TestQueryRow(t *testing.T) {
	t.SkipNow()

	t.Run("", func(t *testing.T) {
		var id string

		if err := QueryRow(nil, "select id from t limit 1", nil).Scan(&id); err != nil {
			t.Fail()
		}
	})

	t.Run("", func(t *testing.T) {
		var ids []string
		err := Query(nil, "select id from t limit 1", nil).Each(func() (func(), []interface{}) {
			var id string
			return func() { ids = append(ids, id) }, []interface{}{&id}
		})

		if err != nil {
			t.Fail()
		}
	})
}
