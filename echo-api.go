package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()
    e.GET("/", hello)
    e.GET("/:id", helloId)
    e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, Echo.")
}

func helloId(c echo.Context) error {
    id := c.Param("id")
    return c.String(http.StatusOK, "Hello "+id)
}
