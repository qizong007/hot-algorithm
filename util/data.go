package util

import (
	"math/rand"
	"time"
)

func GenerateIntList(n int) []int {
	list := make([]int, n)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range list {
		list[i] = r.Int()
	}
	return list
}

func CheckIntSorted(list []int) bool {
	n := len(list)
	for i := 0; i < n-1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}
