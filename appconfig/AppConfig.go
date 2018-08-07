package appconfig

import "github.com/ilibs/gosql"

type AppConfig struct {
  Host string
  Port string
  DataBase gosql.Config
}
