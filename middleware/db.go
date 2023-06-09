package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/survey-app/survey/app"
)

func DB() *dbHandler {
	if dbh == nil {
		dbh = &dbHandler{}
	}
	return dbh
}

var dbh *dbHandler

type dbHandler struct{}

func (*dbHandler) New(c *fiber.Ctx) error {
	ctx, ok := c.Locals(app.CtxKey).(*app.Ctx)
	if !ok {
		return app.NewError(http.StatusInternalServerError, "ctx is not found")
	}
	ctx.TxBegin()
	err := c.Next()
	if err != nil || (c.Response().StatusCode() >= http.StatusBadRequest || c.Response().StatusCode() < http.StatusOK) {
		ctx.TxRollback()
	} else {
		ctx.TxCommit()
	}
	return nil
}
