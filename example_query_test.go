package sqlfly_test

import (
	"fmt"
	"strconv"

	"github.com/Chyroc/sqlfly"
)

func Example_QueryRow_NoRow() {
	testDB.Exec("truncate test")

	var id int
	if err := sqlfly.QueryRow(testDB, "select id from test limit 1").Scan(&id); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("id: %v\n", id)

	// output:
	// err: sql: no rows in result set
}

func Example_QueryRow_WithRow() {
	testDB.Exec("truncate test")
	if _, err := testDB.Exec("insert into test (name, age) values (?, ?)", "test_name", 666); err != nil {
		panic(err)
	}

	var b testTable
	if err := sqlfly.QueryRow(testDB, "select id, name, age, created_at from test limit 1").Scan(&b.ID, &b.Name, &b.Age, &b.CreatedAt); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("%s id(%d) age(%d)\n", b.Name, b.ID, b.Age)

	// output:
	// test_name id(1) age(666)
}

func Example_Query() {
	testDB.Exec("truncate test")
	for i := 0; i < 10; i++ {
		if _, err := testDB.Exec("insert into test (name, age) values (?, ?)", strconv.Itoa(i), i); err != nil {
			panic(err)
		}
	}

	var bs []testTable
	err := sqlfly.Query(testDB, "select id, name, age, created_at from test").Each(func() (func(), []interface{}) {
		var b testTable
		return func() { bs = append(bs, b) }, []interface{}{&b.ID, &b.Name, &b.Age, &b.CreatedAt}
	})
	if err != nil {
		panic(err)
	}

	for _, b := range bs {
		fmt.Printf("%s id(%d) age(%d)\n", b.Name, b.ID, b.Age)
	}

	// output:
	// 0 id(1) age(0)
	// 1 id(2) age(1)
	// 2 id(3) age(2)
	// 3 id(4) age(3)
	// 4 id(5) age(4)
	// 5 id(6) age(5)
	// 6 id(7) age(6)
	// 7 id(8) age(7)
	// 8 id(9) age(8)
	// 9 id(10) age(9)
}
