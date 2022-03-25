package util

import (
	"fmt"
	"time"
)

func ExecuteWithTime(do func()) {
	start := time.Now()
	do()
	fmt.Println("Time Cost: ", time.Since(start))
}
