package sqlfly

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path"

	_ "github.com/go-sql-driver/mysql" // 注册mysql driver
)

var testDB *sql.DB

func init() {
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
