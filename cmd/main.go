package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"tk_app_v1/src/api"
	"tk_app_v1/src/database"

	"github.com/joho/godotenv"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	envFilePath := basepath + "/../.env"
	envFile, err := godotenv.Read(envFilePath)
	if err != nil {
		panic(err)
	}

	database.Connect(database.DatabaseLogin{
		Host:     envFile["PG_HOST"],
		Port:     envFile["PG_PORT"],
		Username: envFile["PG_USER"],
		Password: envFile["PG_PASSWORD"],
		Name:     envFile["PG_DATABASE_NAME"],
	})
	defer database.Connection.Db.Close()

	fmt.Println("success")

	database.CreateTables(database.Connection)

	// Router
	api.StartRouter(envFile["SERVER_PORT"])
	
}
