// logger/logger.go
package logger

import (
    "github.com/labstack/echo/v4"
)

var Log echo.Logger

func Init(e *echo.Echo) {
    Log = e.Logger
}
