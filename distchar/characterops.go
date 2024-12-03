package distchar

import (
	"fmt"
	"strings"
)

type DistChar struct {
	//inputstr string
	mode       int
	filetrmode bool
}

func NewDistChar(modeop int, filetrmode bool) *DistChar {
	return &DistChar{
		mode:       modeop,
		filetrmode: filetrmode,
	}
}

func (dst *DistChar) SetFilterModeOn() {
	dst.filetrmode = true
}

func (dst *DistChar) SetFilterModeOff() {
	dst.filetrmode = false
}

func (dst *DistChar) CountDistChar(input string) int {
	charset := make(map[byte]bool)

	//fmt.Printf("input=%v\n", input)
	if dst.mode == 1 {
		return dst.countDistCharsCase(input, charset)
	} else {

		input = strings.ToLower(input)
		if dst.filetrmode {
			input = filterSpceialChars(input)
		}
		return dst.countDistCharsCase(input, charset)
	}
}
func (dst *DistChar) countDistCharsCase(input string, set map[byte]bool) int {

	for i := 0; i < len(input); i++ {
		set[input[i]] = true
	}
	return len(set)
}

func filterSpceialChars(input string) string {
	fmt.Printf("before Filter:=%v\n", input)
	var builder strings.Builder
	for i := 0; i < len(input); i++ {
		diff := input[i] - 'a'
		diffWithUpper := input[i] - 'A'
		diffWitNum := input[i] - '0'
		if (diff >= 0 && diff < 26) || (diffWithUpper >= 0 && diffWithUpper < 26) || (diffWitNum >= 0 && diffWitNum < 10) {
			builder.WriteByte(input[i])
		}
	}
	output := builder.String()
	fmt.Printf("After Filter:=%v\n", output)
	return output
}
