package main

import (
	"fmt"
	"password-generator-api/config"
	"password-generator-api/routes"
)

func main() {
	enviroment := config.GoDotEnvVariable("GO_ENV")
	host := config.GoDotEnvVariable("GO_HOST")
	port := config.GoDotEnvVariable("GO_PORT")

	fmt.Println("initializing go server in", enviroment)
	routes.Router(host, port)
}
