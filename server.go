package main

import (
  "encoding/json"
  "io/ioutil"
  "fmt"
  "github.com/Skovy/metroweb/database"
  "github.com/go-martini/martini"
)

type Configuration struct {
  Db_user        string
  Db_password    string
  Db_host        string
  Db_database    string
}

var conf Configuration

// decode and load the config.json file
func loadConfiguration() {

  content, _ := ioutil.ReadFile("config.json")
  json.Unmarshal(content, &conf)
}

func main() {

  loadConfiguration()
  fmt.Printf("%v", conf)
  database.DatabaseConnect(conf.Db_user, conf.Db_password, conf.Db_host, conf.Db_database)

  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Run()
}
