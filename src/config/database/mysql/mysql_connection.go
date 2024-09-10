package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MYSQL_URL    = "MYSQL_URL"
	MYSQL_DRIVER = "MYSQL_DRIVER"
)

func NewMysqlConnection() (*sql.DB, error) {
	mysql_url := os.Getenv(MYSQL_URL)
	mysql_driver := os.Getenv(MYSQL_DRIVER)

	db, err := sql.Open(mysql_driver, mysql_url)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Conexão bem sucedida")

	return db, nil
}
