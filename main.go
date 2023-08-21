package main

import (
	"echo-graphql-tutorial/database"
	"echo-graphql-tutorial/routes"
)

func main() {
	database.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
