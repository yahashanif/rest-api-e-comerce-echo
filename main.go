package main

import (
	"rest-api-e-comerce/db"
	"rest-api-e-comerce/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Start(":9000")
}
