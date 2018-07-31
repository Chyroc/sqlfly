package sqlfly

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path"
	"time"

	_ "github.com/go-sql-driver/mysql" // 注册mysql driver
)

var testDB *sql.DB

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	dber, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/sqlfly?charset=utf8mb4,utf8&parseTime=True&loc=UTC")
	if err != nil {
		panic(err)
	}
	testDB = dber
	initMySQLTable("test")
}

func initMySQLTable(sqlName string) {
	f, err := ioutil.ReadFile(path.Join("testdata", sqlName) + ".sql")
	if err != nil {
		panic(err)
	}
	if _, err = testDB.Exec(fmt.Sprintf("DROP TABLE IF EXISTS `%s`;", sqlName)); err != nil {
		panic(err)
	}
	if _, err = testDB.Exec(string(f)); err != nil {
		panic(err)
	}
}

func generateRandomData(table string, count int) {
	switch table {
	case "test":
		for i := 0; i < count; i++ {
			_, err := testDB.Exec("insert into test (name, age) values (?, ?)", randString(10), randInt(3))
			if err != nil {
				panic(err)
			}
		}
	}
}

func randInt(n int) int {
	s := 0
	for i := 0; i < n; i++ {
		s = s*10 + rand.Intn(9) + 1 // [0,n)
	}
	return s
}

func randString(n int) string {
	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type testTable struct {
	ID        int
	Name      string
	Age       int
	CreatedAt time.Time
}
