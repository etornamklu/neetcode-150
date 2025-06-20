package enodeanddecode

import (
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder

	for _, str := range strs {
		builder.WriteString(strconv.Itoa(len(str)))
		builder.WriteByte('#')
		builder.WriteString(str)
	}

	return builder.String()

}

func (s *Solution) Decode(str string) []string {
	var result []string

	i := 0
	//iterate over string
	for i < len(str) {
		j := i

		// count till you encounter delimiter
		if str[j] != '#' {
			j++ // current value of j is set to the index of delimiter
		}

		length, _ := strconv.Atoi(str[i:j])        // this get's the number before the delimiter, the length of the word
		result = append(result, str[j+1:length+1]) // slice, start at char after delimiter and end at the length of the string
		i = j + length + 1                         // set i to the next char after the slice
	}
	return result
}
