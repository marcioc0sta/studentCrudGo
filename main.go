package main

import (
	"github.com/marcioc0sta/gin-api-rest/db"
	"github.com/marcioc0sta/gin-api-rest/routes"
)

func main() {
	db.ConnectWithDatabase()
	routes.HandleRequests()
}