package errors

import (
	"fmt"
	"log"
)

type Error struct {
	Op         Op
	Err        error
	Msg        string
	StatusCode StatusCode
}

type Op string
type StatusCode int

func (e *Error) Error() string {
	return fmt.Sprintf("%#v", e)
}

func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case error:
			e.Err = arg
		case string:
			e.Msg = arg
		case StatusCode:
			e.StatusCode = arg
		default:
			panic("bad call to E")
		}
	}
	return e
}

func (e *Error) Ops() []Op {
	res := []Op{e.Op}
	subErr, ok := e.Err.(*Error)
	if !ok {
		return res
	}

	res = append(res, subErr.Ops()...)
	return res
}

func (e *Error) StatusCodes() []StatusCode {
	res := []StatusCode{}
	if e.StatusCode != 0 {
		res = append(res, e.StatusCode)
	}
	subErr, ok := e.Err.(*Error)
	if !ok {
		return res
	}

	res = append(res, subErr.StatusCodes()...)
	return res
}

func (e *Error) LogError() {
	log.Println(e.Msg)
	log.Println(e.Err.Error())
	log.Println(e.Ops())
}
