package util

import (
	"math/rand"
	"time"
)

var randSrc = rand.NewSource(time.Now().UnixNano())

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	sliceLength = 6
	mask = 1 << sliceLength - 1
	indexMax = 63 / sliceLength
	returnLength = 40
)

func RandString() string{
	b := make([]byte, returnLength)
	cache, remain :=  randSrc.Int63(), indexMax
	for i:= returnLength-1; i>=0;{
		if remain == 0{
			cache, remain =  randSrc.Int63(), indexMax
		}
		index := int(cache & mask)
		if index < len(letters){
			b[i] = letters[index]
			i--
		}
		cache >>= sliceLength
		remain--
	}
	return string(b)
}