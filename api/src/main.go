package main

import "simple_sns_api/db"

func main() {
	db.CreateConnection()
	defer db.Client.Close()
	router().Run()
}
