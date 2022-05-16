package reverse_string

import (
	//"fmt"
	"bufio"
	"os"
)

func ReverseString(input string) (output string) {
	//var some_string string
	some_string, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//fmt.Scan(&some_string)
	line := []rune(some_string)
	line2 := make([]rune, len(line))

	for i := range line {
		line2[i] = line[len(line)-1-i]
	}
	output = string(line2)
	return output
}
