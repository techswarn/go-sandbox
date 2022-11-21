package strings

import "fmt"

// AddFish adds fish to a string.
func AddFish(str string) string {
	return fmt.Sprintf("><>  %s  <><", str)
}
