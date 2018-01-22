package main

import (
	db "gin-restful/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":3001")
}
