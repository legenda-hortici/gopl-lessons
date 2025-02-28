package main

import (
	"fmt"
)

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast
	FlagLoopback                       // is loopback
	FlagPointToPoint                   // point-to-point link
	FlagMulticast                      // supports multicast
)

func IsUP(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUP(v)) // 10001 true
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUP(v)) // 10000 false
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUP(v))   // 10010 true
	fmt.Printf("%b %t\n", v, IsCast(v)) // 10010 true

	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)
}

const (
	KB = 1 << 10 * iota
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
