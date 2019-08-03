package comerr

import "fmt"

// customize business error type
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

// customize system error type
type SysErr struct {
	err string
}

func NewSysErr(err string) *SysErr {
	return &SysErr{
		err: err,
	}
}

func SysErrorf(format string, a ...interface{}) *SysErr {
	return NewSysErr(fmt.Sprintf(format, a...))
}

func (b *SysErr) Error() string {
	return b.err
}
