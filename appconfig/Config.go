package appconfig

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
)

var Config AppConfig

func Init()  {
    b, err := ioutil.ReadFile("appconfig.json")
    if err != nil {
      panic(err)
      fmt.Println("err")
    }
    json.Unmarshal(b, &Config)
    fmt.Println(Config.DataBase)
}
