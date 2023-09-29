package routes

import (
    "net/http"
    "going-places-api/api/handlers"
    "github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Welcome to the travel psychic!")
    })
    e.GET("/add-word", handlers.AddWord)
    e.GET("/remove-word", handlers.RemoveWord)
}
