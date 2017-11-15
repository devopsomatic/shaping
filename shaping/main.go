package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func helloWorld(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!");

}

func main() {

	// Initialize echo
	e := echo.New()

	// Route definition
	e.GET("/", helloWorld);

	// Start server
	e.Start("127.0.0.1:3000");
}
