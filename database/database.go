package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	return sql.Open("mysql", "root:123321@/learn")
}
