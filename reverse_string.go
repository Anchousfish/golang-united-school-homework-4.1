package reverse_string

func ReverseString(input string) (output string) {

	line := []rune(input)
	line2 := make([]rune, len(line))

	for i := range line {
		line2[i] = line[len(line)-1-i]
	}
	output = string(line2)
	return output
}
