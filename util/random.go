// this program helps to run random inputs to the database for testing purpose
package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	followedByAt = "@"
)

// function to display different time out for each input
func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName() string {
	return RandomString(4)
}

func RandomEmail() string {
	return RandomString(5) + followedByAt + "yahoo.com"
}

func RandomUserName() string {
	return followedByAt + RandomString(5)
}

func RandomPassword() int {
	return int(RandomInt(4, 12))
}
