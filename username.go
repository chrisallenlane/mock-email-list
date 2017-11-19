package main

import (
	"math/rand"
	"time"
)

var username = func() func(n int) string {

	// Flagrantly stolen, with gratitude:
	// https://stackoverflow.com/a/31832326/461108
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)

	src := rand.NewSource(time.Now().UnixNano())

	return func(length int) string {
		bytes := make([]byte, length)
		for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
			if remain == 0 {
				cache, remain = src.Int63(), letterIdxMax
			}
			if idx := int(cache & letterIdxMask); idx < len(letters) {
				bytes[i] = letters[idx]
				i--
			}
			cache >>= letterIdxBits
			remain--
		}

		return string(bytes)
	}
}()
