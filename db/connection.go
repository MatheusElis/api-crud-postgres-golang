package db

import (
	"database/sql"
	"fmt"

	"github.com/MatheusElis/api-crud-postgres-golang/configs"
  _ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error){
  conf := configs.GetDB()

  sc := fmt.Sprintf("host=%s port=%s user=%s pass=%s dbname=%s sslmode=disable",
    conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

  conn, err := sql.Open("postgres",sc)
  if err != nil {
    panic(err)
  }

  err = conn.Ping()

  return conn, err
}
