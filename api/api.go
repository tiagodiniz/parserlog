package api

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/parserlog/parser"
)

func API(ctx echo.Context) error {

	// parser log files

	result, err := parser.Process()

	if err == nil {

		// render response
		err = ctx.JSON(http.StatusOK, result)

	}

	return err
}