// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 244.

// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	"os"
	"time"
)

//!+
func main() {
	fmt.Println("发射倒计时")
	tick := time.Tick(1 * time.Second)
	terminate := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		terminate <- struct{}{}
	}()
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-terminate:
			fmt.Println("发射终止")
			return
		}
	}
	launch()
}

//!-

func launch() {
	fmt.Println("火箭发射")
}
