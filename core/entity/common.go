package entity

import "strings"

type SignQueue []string

// TODO
func (queue SignQueue) CheckSign(sign string) bool {

	return true
}

func (queue SignQueue) AppendSignData(key, val string) SignQueue {
	if "" != strings.TrimSpace(val) {
		str := key + "=" + val
		queue = append(queue, str)
	}
	return queue
}
