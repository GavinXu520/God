package util

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// Generate str character by rand
func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(rs, "")
}

// Sub string by [start, end]
func SubString(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}
	return string(rs[(start - int(1)):end])
}

//
func CheckPhoneNo(phoneNo string) bool {
	//reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	reg := `^((13[0-9])|(14[0-9])|(15[0-9])|(16[0-9])|(17[0-9])|(18[0-9]))\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phoneNo)
}

func CheckIP(ip string) bool {
	reg := `(2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})(\.(2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})){3}`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(ip)
}

func CheckPwdLen(pwd string, limit int) bool {
	return utf8.RuneCountInString(pwd) == limit
}
