package main

import "github.com/labstack/echo/v4"

func main() {

	server := echo.New()

	server.Start("localhost:8080")

}
