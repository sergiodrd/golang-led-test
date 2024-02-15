package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo/v4"
)

func main() {
        e := echo.New()
        e.GET("/", func(c echo.Context) error {
                f, err := os.ReadFile("index.html")
                if err != nil {
                        log.Fatal(err)
                }
                return c.HTML(http.StatusOK, string(f))
        })
        e.POST("/set", func(c echo.Context) error {
                txt := c.FormValue("txt")
                color := c.FormValue("color")
                log.Printf("Text: %s\nColor: %s\n", txt, color)
                cmd := exec.Command("python", "pythontest.py", "-t", txt, "--led-gpio-mapping=adafruit-hat", "--led-slowdown", "4", "--rgb", color)
                if errors.Is(cmd.Err, exec.ErrDot) {
                        cmd.Err = nil
                }
                if err := cmd.Run(); err != nil {
                        log.Fatal(err)
                }
                return c.String(http.StatusOK, "")
        })
        e.Logger.Fatal(e.Start(":6969"))
}
