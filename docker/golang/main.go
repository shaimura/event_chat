package main

import (
	"example.com/go-mod/app/db"
	"example.com/go-mod/app/router"
)

func main() {

	db.Connect()
	defer db.Close()

	router.Router()

}
