package main

import (
	"github.com/labstack/echo"
	"gopkg.in/rjz/githubhook.v0"
	"flag"
)

func githubWebHook(c echo.Context) error {
	key := flag.String("key", "not secured", "github token")
	flag.Parse()
	secret := []byte(*key);
	_, err := githubhook.Parse(secret, c.Request());
	return err;
}

func main() {


	// Initialize echo
e := echo.New()

	   // Route definition
	   e.POST("/github", githubWebHook);

   // Start server
   e.Start("127.0.0.1:3000");
}
