package handler

import (
	"github.com/admpub/caddyui/application/library/errors"
	"github.com/webx-top/echo"
)

var PageMaxSize = 1000

func Paging(ctx echo.Context) (page int, size int) {
	page = ctx.Formx(`page`).Int()
	size = ctx.Formx(`size`).Int()
	if page < 1 {
		page = 1
	}
	if size < 1 || size > PageMaxSize {
		size = 50
	}
	return
}

func Ok(v string) errors.Successor {
	return errors.NewOk(v)
}
