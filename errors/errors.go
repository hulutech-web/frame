package errors

import (
	"fmt"

	"github.com/ztrue/tracerr"

	"gitee.com/hulutech/frame/helpers/toto"
	"gitee.com/hulutech/frame/logs"
)

func ErrPrintln(err error, fields toto.V) {
	startFrom := 2
	if err == nil {
		return
	}
	traceErr := tracerr.Wrap(err)
	frameList := tracerr.StackTrace(traceErr)
	if startFrom > len(frameList) || len(frameList)-2 <= 0 {
		logs.Println(logs.ERROR, err.Error(), fields)
	}

	traceErr = tracerr.CustomError(err, frameList[startFrom:len(frameList)-2])
	traceErr = tracerr.CustomError(err, frameList)

	if fields == nil {
		fields = toto.V{}
	}
	fields["tmaic_trace"] = tracerr.SprintSource(traceErr, 0)
	logs.Println(logs.ERROR, err.Error(), fields)
}

func ErrPrint(err error, startFrom int, fields toto.V) string {
	if err == nil {
		return ""
	}
	traceErr := tracerr.Wrap(err)
	frameList := tracerr.StackTrace(traceErr)

	traceErr = tracerr.CustomError(err, frameList)

	if fields == nil {
		fields = toto.V{}
	}
	fields["tmaic_trace"] = tracerr.SprintSource(traceErr)
	return fmt.Sprint(err.Error(), fields)
}
