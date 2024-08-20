package util

import (
	"math/rand"
	"time"
)

func MockSleep(dur time.Duration) {
	time.Sleep(time.Duration(rand.Int63n(int64(dur))))
}

func LeaveItToRNG() bool {
	return rand.Intn(3) == 1
}
