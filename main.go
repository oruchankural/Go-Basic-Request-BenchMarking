package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Main handler is working ...")

}
func userHandler(c echo.Context) error {
	return c.String(http.StatusOK, "User handler is working ...")
}
func main() {

	fmt.Printf("Test")

	e := echo.New()
	e.GET("/main", mainHandler)
	e.GET("/main", userHandler)
	e.Start(":8080")

}
