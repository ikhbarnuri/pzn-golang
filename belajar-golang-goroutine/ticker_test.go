package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for curTime := range ticker.C {
		fmt.Println(curTime)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for curTime := range channel {
		fmt.Println(curTime)
	}
}
