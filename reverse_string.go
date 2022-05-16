package reverse_string

import (
	"bufio"
	"os"
)

func ReverseString(input string) (output string) {

	some_string, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//line:=make([]rune,0,len(some_string))
	for i := range some_string {
		output += string(some_string[len(some_string)-1-i])
	}

	return output
}
