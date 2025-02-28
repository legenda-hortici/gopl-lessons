package main

import (
	"fmt"
	"time"
)

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint
const (
	FlagUp Flags = 1 << iota // is up
	FlagBroadcast            // supports broadcast
	FlagLoopback             // is loopback
	FlagPointToPoint         // point-to-point link
	FlagMulticast            // supports multicast
)

const (
	noDelay time.Duration = 0
	timeOut               = 5 * time.Minute
)

const (
	a = 1
	b
	c = 2
	d
)

func main() {
	fmt.Printf("%T %[1]v\n", noDelay)     // "0"
	fmt.Printf("%T %[1]v\n", timeOut)     // "5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "1m0s"

	fmt.Println(a, b, c, d) // 1 1 2 2
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
}
