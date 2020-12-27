package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	type a struct {
		Aa int `json:"aa"`
		Bb string `json:"bb"`
	}

	aStruct := a{11, "hello"}
	//marshal, _ := json.Marshal(aStruct)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, aStruct)
	})
	e.Logger.Fatal(e.Start(":1323"))
}