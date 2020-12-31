//Conversion between meter and centimeter
package lengthConversion

type Metre float64
type Centimeter float64

const multiple = 100

func MToC(m Metre) Centimeter  {
	return Centimeter(m*multiple)
}

func CToM(c Centimeter) Metre  {
	return Metre(c/multiple)
}

