package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func echoHandler(c echo.Context) error {
	r := c.Request()
	text := fmt.Sprintf("%s %s %s", r.Method, r.RequestURI, r.Proto)
	text = fmt.Sprintf("%s\nHost: %s", text, r.Host)
	for key, header := range r.Header {
		for _, line := range header {
			text = fmt.Sprintf("%s\n%s: %s", text, key, line)
		}
	}
	return c.String(http.StatusOK, text)
}

func helpHandler(c echo.Context) error {
	cmd := c.Param("cmd")
	return c.String(http.StatusOK, cmd)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/help/:cmd", helpHandler)
	e.GET("/echo", echoHandler)

	e.Logger.Fatal(e.Start(":80"))
}
