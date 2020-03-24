package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func getInfo(c echo.Context) string {
	r := c.Request()
	text := fmt.Sprintf("%s %s %s", r.Method, r.RequestURI, r.Proto)
	text = fmt.Sprintf("%s\nHost: %s", text, r.Host)
	for key, header := range r.Header {
		for _, line := range header {
			text = fmt.Sprintf("%s\n%s: %s", text, key, line)
		}
	}

	params, err := c.FormParams()
	if err != nil {
		text = fmt.Sprintf("%s\n--- error: %v", text, err)
	} else {
		text += "\n---"
		for param := range params {
			text = fmt.Sprintf("%s\n%s: %s", text, param, c.FormValue(param))
		}
	}
	return text
}

func echoHandler(c echo.Context) error {
	text := getInfo(c)
	fmt.Printf("GET request ------\n%v\n", text)
	return c.String(http.StatusOK, text)
}

func loginHandler(c echo.Context) error {
	text := getInfo(c)

	cookie := new(http.Cookie)
	cookie.Name = "meari_session"
	cookie.Value = "RANDOMIZED_SESSION_STRING"
	cookie.Expires = time.Now().Add(1 * time.Hour)
	c.SetCookie(cookie)

	fmt.Printf("POST request ------\n%v\n", text)
	return c.Redirect(http.StatusTemporaryRedirect, "/")
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
	e.POST("/echo", echoHandler)
	e.GET("/login", echoHandler)
	e.POST("/login", loginHandler)

	e.Logger.Fatal(e.Start(":80"))
}
