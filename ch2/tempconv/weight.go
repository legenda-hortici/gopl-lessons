package tempconv

import "fmt"

type Gram float64
type Kilogram float64
type Ton float64

func (g Gram) String() string     { return fmt.Sprintf("%g gr", g) }
func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }
func (t Ton) String() string      { return fmt.Sprintf("%g T", t) }
