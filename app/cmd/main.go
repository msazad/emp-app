package main

import "log"

func main() {

	dDB, sqlDB, err := config.DBConnection()

	if err != nil {
		log.Fatalf("error due to : %v", err)
	}

	defer sqlDB.Close()

	r := app.ApiRoute(gDB)

	if err:=app.Start(r);err!=nil{
		log.Fatalf("error due to : %v",err)
	}
}
