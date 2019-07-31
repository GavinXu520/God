package comerr

import "fmt"

type BizErr struct {
	err string
}

func NewBizErr(err string) *BizErr {
	return &BizErr{
		err: err,
	}
}

func BizErrorf(format string, a ...interface{}) *BizErr {
	return NewBizErr(fmt.Sprintf(format, a...))
}

func (b *BizErr) Error() string {
	return b.err
}
