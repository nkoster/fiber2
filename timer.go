package main

import (
	"fmt"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println(name+":", "finished in", elapsed)
}
