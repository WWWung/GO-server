package main

import (
  _"fmt"
  _"net/http"

  "./appconfig"

  _"github.com/go-sql-driver/mysql"
  _"github.com/ilibs/gosql"
)

func main() {
  appconfig.Init()
}
