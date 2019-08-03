package entity

import (
	"sort"
	"strings"
)

type SignQueue struct {
	Arr  []string
	Dict map[string]string
}

// TODO
func (queue SignQueue) CheckSign(sign string) bool {
	// todo test
	return true

	// must sort the key first
	sort.Strings(queue.Arr)
	memArr := make([]string, 0)
	for _, key := range queue.Arr {
		val := queue.Dict[key]
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
		queue.Arr = append(queue.Arr, key)
		queue.Dict[key] = str
	}
}
