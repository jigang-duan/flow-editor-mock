package common

import (
	"github.com/jmespath/go-jmespath"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func SendJSON(ctx iris.Context, resp interface{}) (err error) {
	indent := ctx.URLParamDefault("indent", "  ")
	if query := ctx.URLParam("query"); query != "" && query != "[]" {
		resp, err = jmespath.Search(query, resp)
		if err != nil {
			return
		}
	}

	_, err = ctx.JSON(resp, context.JSON{Indent: indent, UnescapeHTML: true})
	return err
}
