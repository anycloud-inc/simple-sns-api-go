package main

import "simple_sns_api/src/db"

func main() {
	db.CreateConnection()
	defer router().Run()
}
