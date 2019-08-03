package entity

import (
	"sort"
	"strings"
)

type SignQueue struct {
	arr  []string
	dict map[string]string
}

// TODO
func (queue SignQueue) CheckSign(sign string) bool {
	// todo test
	return true

	// must sort the key first
	sort.Strings(queue.arr)
	memArr := make([]string, 0)
	for _, key := range queue.arr {
		val := queue.dict[key]
		memArr = append(memArr, val)
	}
	str := strings.Join(memArr, "&")
	if str != sign {
		return false
	}
	return true
}

func (queue *SignQueue) AppendSignData(key, val string) {
	if "" != strings.TrimSpace(val) {
		str := key + "=" + val
		queue.arr = append(queue.arr, key)
		queue.dict[key] = str
	}
}
