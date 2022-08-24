package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func Recovery() HandleFunc {
	return func(ctx *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				// 捕捉堆栈
				log.Printf("%s\n\n", trace(message))
				ctx.Fail(http.StatusInternalServerError, "internal server error")
			}
		}()
		ctx.Next()
	}
}

func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])
	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, l := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, l))
	}
	return str.String()
}
