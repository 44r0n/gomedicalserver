package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
	if db == nil {
		dbt, err := sql.Open("mysql", "root:root@/gotest")
		if err != nil {
			panic(err.Error())
		}

		err = dbt.Ping()

		if err != nil {
			panic (err.Error())
		}

		db = dbt	
	}	
}

func ExecuteNonQuery(query string, args ...interface{}) {
	_, err := db.Exec(query, args...)

	if err != nil {
		panic(err.Error())
	}
}

func ExecuteQuery(query string, args ...interface{}) *sql.Rows {
	rows, err := db.Query(query, args...)
	
	if err != nil {
		panic(err.Error())
	}
	return rows
}

func Close() {

	if db != nil {
		db.Close()
		db = nil
	}
}