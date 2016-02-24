package tmpconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius    = -273.15
	FreazingC     Celsius    = 0
	BoilingC      Celsius    = 100
	AbsoluteZeroF Fahrenheit = -459.66999999999996
	FreazingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f)}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5/ 9) }
