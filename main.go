package main

import (
	_ "gin_curd/database"
	"gin_curd/router"
)

func main() {
	r := router.InitRouter()
	r.Run()
}
