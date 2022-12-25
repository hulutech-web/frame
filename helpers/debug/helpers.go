package debug

import (
	"errors"

	"github.com/davecgh/go-spew/spew"
	"github.com/ztrue/tracerr"

	"gitee.com/hulutech/frame/console"
	"gitee.com/hulutech/frame/helpers/log"
)

func Dump(v ...interface{}) {
	console.Println(console.CODE_ERROR, spew.Sdump(v...))
	debugPrint(errors.New("====== Tmaic Debug ======"))
}

func debugPrint(err error) {
	startFrom := 2
	traceErr := tracerr.Wrap(err)
	frameList := tracerr.StackTrace(traceErr)
	if startFrom > len(frameList) || len(frameList)-2 <= 0 {
		_ = log.Error(err)
	}
	traceErr = tracerr.CustomError(err, frameList[startFrom:len(frameList)-2])
	tracerr.PrintSourceColor(traceErr, 0)
}
