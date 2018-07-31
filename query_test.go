package sqlfly_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Chyroc/sqlfly"
)

func TestQuery(t *testing.T) {
	as := assert.New(t)

	t.Run("query rows", func(t *testing.T) {
		var ids []string
		as.Nil(sqlfly.Query(testDB, "select id from test").Each(func() (func(), []interface{}) {
			var id string
			return func() { ids = append(ids, id) }, []interface{}{&id}
		}))
		as.Len(ids, 0)
	})

	t.Run("more", func(t *testing.T) {
		generateRandomData("test", 20)

		var tests []testTable
		as.Nil(sqlfly.Query(testDB, "select id, name, age, created_at from test").Each(func() (func(), []interface{}) {
			var ta testTable
			return func() { tests = append(tests, ta) }, []interface{}{&ta.ID, &ta.Name, &ta.Age, &ta.CreatedAt}
		}))
		as.Len(tests, 20)
		first := tests[0]
		as.Equal(1, first.ID)
		t.Logf("frst %#v\n", first)
	})
}
