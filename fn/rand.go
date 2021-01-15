package fn

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandInt(low, high int) int {
	return low + rand.Intn(high-low)
}

func RandInt64(low, high int64) int64 {
	return low + rand.Int63n(high-low)
}
