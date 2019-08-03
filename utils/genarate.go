package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//生成一个订单号
func GetAOrderId() string {
	orderId := time.Now().Format("20060102150405")
	orderId += strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(90000) + 10000)
	return orderId
}

func RandSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//随机数生成，size随机数的位数
func RandNumLetter(size int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	arr := make([]string, size)
	//随机种子，加随机数防止并发
	source := rand.NewSource(time.Now().UnixNano() + rand.Int63n(9999))
	r := rand.New(source)
	for i := 0; i < size; i++ {
		index := r.Intn(62)
		arr[i] = str[index : index+1]
	}
	return strings.Join(arr, "")
}

func RandStr() string {
	t := time.Now()
	h := md5.New()
	io.WriteString(h, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	io.WriteString(h, t.String())
	passwd := fmt.Sprintf("%x", h.Sum(nil))
	return passwd
}
