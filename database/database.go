package database

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

// create a connection with the databae
func DatabaseConnect(user, password, host, database string) {
  db, err := sql.Open("mysql", user + ":" + password + "@tcp(" + host + ")/" + database)
  if err != nil {
    fmt.Println(err)
  }

  _, err = db.Exec("CREATE TABLE users(id int)")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
}
