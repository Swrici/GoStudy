//Overriding the output of a string
package lengthConversion

import "fmt"

func (metre Metre) String() string {
	return fmt.Sprintf("%gm",metre)
}

func (centimeter Centimeter) String() string {
	return fmt.Sprintf("%gcm",centimeter)
}