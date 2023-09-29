package main

import (
    "going-places-api/api/routes"
    "going-places-api/db"
    "going-places-api/logger"
    "github.com/labstack/echo/v4"
)

var e *echo.Echo

func main() {
    e = echo.New()

    // Initialize the logger
    logger.Init(e)

    db.InitRedis() // Initialize Redis
    routes.SetupRoutes(e) // Setup routes

    e.Logger.Info("Starting the server...")
    e.Logger.Fatal(e.Start(":1323"))
}
