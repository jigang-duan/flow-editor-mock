package common

import (
	"fmt"
	"github.com/kataras/iris"
)

type HttpError struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

func (h HttpError) Error() string {
	return fmt.Sprintf("Status Code: %d\nReason: %s", h.Code, h.Reason)
}

func Fail(ctx iris.Context, statusCode int, format string, a ...interface{}) {
	err := HttpError{
		Code:   statusCode,
		Reason: fmt.Sprintf(format, a...),
	}

	ctx.Application().Logger().Error(err)

	ctx.StatusCode(statusCode)
	_, _ = ctx.JSON(err)

	ctx.StopExecution()
}
