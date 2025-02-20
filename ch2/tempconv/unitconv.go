package tempconv

import "fmt"

type Millimetr float64
type Centimetr float64
type Metr float64

func (mill Millimetr) String() string { return fmt.Sprintf("%g mill", mill) }
func (cm Centimetr) String() string   { return fmt.Sprintf("%g cm", cm) }
func (m Metr) String() string         { return fmt.Sprintf("%g m", m) }
