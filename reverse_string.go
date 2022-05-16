package reverse_string

import (
	"bufio"
	"os"
)

func ReverseString(input string) (output string) {

	some_string, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	line := []rune(some_string)
	line2 := make([]rune, len(some_string))

	for i := range line {
		line2[i] = line[len(some_string)-1-i]
	}
	output = string(line2)
	return output
}
